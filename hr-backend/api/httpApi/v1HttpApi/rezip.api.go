package v1HttpApi

import (
	"archive/zip"
	"bytes"
	"io"
	"mime/multipart"
	"os"

	"gitcode.com/JerichoYu/marmot/global"
	"gitcode.com/JerichoYu/marmot/module/httpModule/v1HttpModule"
	"github.com/gin-gonic/gin"
	filesystem "github.com/jericho-yu/aid/filesystem/v2"
	"go.uber.org/zap"
)

type (
	RezipApi struct{}
)

var (
	Rezip RezipApi
)

func (*RezipApi) New() *RezipApi { return &Rezip }

func (*RezipApi) Upload(c *gin.Context) {
	var (
		title     = "重新压缩"
		file      *multipart.FileHeader
		res       = v1HttpModule.Response.NewInstance(c, v1HttpModule.ResponseConfig{})
		err       error
		srcFile   multipart.File
		buf       []byte
		zipReader *zip.Reader
		rc        io.ReadCloser
	)

	if file, err = c.FormFile("f"); err != nil {
		global.Logger.Error(title, zap.Errors("上传文件", []error{err}))
		res.Fail().Forbidden("上传文件失败")
		return
	}

	// 打开上传文件
	srcFile, err = file.Open()
	if err != nil {
		global.Logger.Error(title, zap.Errors("打开文件", []error{err}))
		res.Fail().Forbidden("打开文件错误")
		return
	}
	defer func() { _ = srcFile.Close() }()

	// 读到内存
	buf = make([]byte, file.Size)
	if _, err = io.ReadFull(srcFile, buf); err != nil {
		global.Logger.Error(title, zap.Errors("读取文件", []error{err}))
		res.Fail().Forbidden("读取文件错误")
		return
	}

	// 创建 zip reader
	zipReader, err = zip.NewReader(bytes.NewReader(buf), file.Size)
	if err != nil {
		global.Logger.Error(title, zap.Errors("创建zip reader", []error{err}))
		res.Fail().Forbidden("创建zip容器失败")
		return
	}

	// 准备写入新 zip 到内存
	var outBuf bytes.Buffer
	zipWriter := zip.NewWriter(&outBuf)

	for _, f := range zipReader.File {
		if f.FileInfo().IsDir() {
			continue
		}

		// 打开原文件
		rc, err = f.Open()
		if err != nil {
			global.Logger.Error(title, zap.Errors("读取压缩包文件", []error{err}))
			res.Fail().Format("读取压缩包文件失败：%s", err.Error())
			return
		}

		// 创建新 zip 内的文件条目
		header := &zip.FileHeader{Name: f.Name, Method: zip.Deflate}
		header.SetModTime(f.ModTime())
		header.SetMode(f.Mode())
		w, err := zipWriter.CreateHeader(header)
		if err != nil {
			_ = rc.Close()
			global.Logger.Error(title, zap.Errors("重建压缩文件", []error{err}))
			res.Fail().Format("重建压缩文件失败：%s", err.Error())
			return
		}

		// 拷贝内容
		if _, err := io.Copy(w, rc); err != nil {
			_ = rc.Close()
			global.Logger.Error(title, zap.Errors("拷贝压缩文件内容", []error{err}))
			res.Fail().Format("拷贝压缩文件内容失败：%s", err.Error())
			return
		}
		_ = rc.Close()
	}

	// 关闭 zip writer
	if err = zipWriter.Close(); err != nil {
		global.Logger.Error(title, zap.Errors("保存新压缩文件", []error{err}))
		res.Fail().Format("保存新压缩文件失败：%s", err.Error())
		return
	}

	fs := filesystem.FileApp.NewByRel("./runtime/upload/rezip/repacked.zip")
	global.Logger.Info(title, zap.String("保存路径", fs.GetFullPath()))
	if err = os.MkdirAll(fs.GetBasePath(), os.ModePerm); err != nil {
		global.Logger.Error(title, zap.Errors("创建目录", []error{err}))
		res.Fail().Format("创建目录失败：%s", err.Error())
		return
	}

	// 保存到本地文件
	if err = os.WriteFile(fs.GetFullPath(), outBuf.Bytes(), 0644); err != nil {
		global.Logger.Error(title, zap.Errors("保存压缩文件", []error{err}))
		res.Fail().Format("保存压缩文件失败：%s", err.Error())
		return
	}

	res.Success().Data("上传成功", gin.H{
		"to": c.Request.Host + "/upload/rezip/repacked.zip",
	})
}
