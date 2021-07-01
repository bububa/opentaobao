package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
手淘拉起旺旺接口 API返回值 
alibaba.interact.sensor.wangwang

手淘开放专用接口，没有数据返回，仅用于手淘容器中jssdk接口鉴权。手淘开放旺旺拉起功能给ISV。
*/
type AlibabaInteractSensorWangwangAPIResponse struct {
    model.CommonResponse
    AlibabaInteractSensorWangwangAPIResponseModel
}

// 手淘拉起旺旺接口 成功返回结果
type AlibabaInteractSensorWangwangAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_interact_sensor_wangwang_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result=0
    Result   string `json:"result,omitempty" xml:"result,omitempty"`
}
