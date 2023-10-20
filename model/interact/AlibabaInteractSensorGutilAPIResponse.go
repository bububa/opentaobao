package interact

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractsensorgutilAPIResponse canvas工具包 API返回值
// alibaba.interact.sensor.gutil
//
// canvas工具包
type AlibabainteractsensorgutilAPIResponse struct {
	model.CommonResponse
	AlibabainteractsensorgutilAPIResponseModel
}

// AlibabainteractsensorgutilAPIResponseModel is canvas工具包 成功返回结果
type AlibabainteractsensorgutilAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_sensor_gutil_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result=0 表示成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}
