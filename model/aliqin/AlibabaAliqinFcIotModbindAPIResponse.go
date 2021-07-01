package aliqin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinFcIotModbindAPIResponse
物联网绑定/换绑API API返回值
alibaba.aliqin.fc.iot.modbind

支持用户的设备的换绑和解绑操作 */
type AlibabaAliqinFcIotModbindAPIResponse struct {
	model.CommonResponse
	AlibabaAliqinFcIotModbindAPIResponseModel
}

// AlibabaAliqinFcIotModbindAPIResponseModel is 物联网绑定/换绑API 成功返回结果
type AlibabaAliqinFcIotModbindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_fc_iot_modbind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaAliqinFcIotModbindResult `json:"result,omitempty" xml:"result,omitempty"`
}
