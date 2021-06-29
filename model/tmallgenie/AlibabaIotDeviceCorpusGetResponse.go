package tmallgenie

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
IoT设备支持语料获取 API返回值 
alibaba.iot.device.corpus.get

ISV通过该接口获取天猫精灵IoT设备支持控制或查询的语料
*/
type AlibabaIotDeviceCorpusGetAPIResponse struct {
    model.CommonResponse
    AlibabaIotDeviceCorpusGetResponse
}

// IoT设备支持语料获取 成功返回结果
type AlibabaIotDeviceCorpusGetResponse struct {
    XMLName xml.Name `xml:"alibaba_iot_device_corpus_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结构体
    RetValues   []DeviceCorpusTopDto `json:"ret_values,omitempty" xml:"ret_values>device_corpus_top_dto,omitempty"`
}
