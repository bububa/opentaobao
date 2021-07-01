package aliqin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinFcIotUseroscontrolAPIResponse
物联卡用户管理停开机功能 API返回值
alibaba.aliqin.fc.iot.useroscontrol

物联网针对用户级管理停开支持 */
type AlibabaAliqinFcIotUseroscontrolAPIResponse struct {
	model.CommonResponse
	AlibabaAliqinFcIotUseroscontrolAPIResponseModel
}

// AlibabaAliqinFcIotUseroscontrolAPIResponseModel is 物联卡用户管理停开机功能 成功返回结果
type AlibabaAliqinFcIotUseroscontrolAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_fc_iot_useroscontrol_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaAliqinFcIotUseroscontrolResult `json:"result,omitempty" xml:"result,omitempty"`
}
