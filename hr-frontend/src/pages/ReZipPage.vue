<template>
  <div class="q-pa-md q-gutter-md">
    <div class="row">
      <div class="col-4"></div>
      <div class="col-4">
        <q-uploader url="http://112.124.0.66:9900/rezip/upload" color="purple" label="选择压缩包" field-name="f" flat
          bordered style="width: 100%" @uploaded="onUploaded" />
      </div>
      <div class="col-4"></div>
    </div>
  </div>
</template>

<script setup>
const onUploaded = function (file) {
  const res = JSON.parse(file.xhr.response)

  alert(res.msg)
  if (res.code === 200) location.href = `http://${res.data.to}`
}
</script>
