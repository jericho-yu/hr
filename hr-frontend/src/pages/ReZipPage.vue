<template>
  <div class="q-pa-md q-gutter-md">
    <div class="row">
      <div class="col-4"></div>
      <div class="col-4">
        <q-uploader url="http://127.0.0.1:9900/rezip/upload" color="purple" label="选择压缩包" field-name="f" flat bordered
          style="width: 100%" @uploaded="onUploaded" />
      </div>
      <div class="col-4"></div>
    </div>
  </div>
</template>

<script setup>
const onUploaded = function (file) {
  // location.href = "http://127.0.0.1:9900/upload/rezip/repacked.zip"
  const res = JSON.parse(file.xhr.response)

  alert(res.msg)
  if (res.code === 200) location.href = `http://${res.data.to}`
}
</script>
