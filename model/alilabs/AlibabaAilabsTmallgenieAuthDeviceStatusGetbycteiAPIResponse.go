package alilabs

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabstmallgenieauthdevicestatusgetbycteiAPIResponse 根据ctei查询状态 API返回值
// alibaba.ailabs.tmallgenie.auth.device.status.getbyctei
//
// 提供给电信查询设备在线状态值
type AlibabaailabstmallgenieauthdevicestatusgetbycteiAPIResponse struct {
	model.CommonResponse
	AlibabaailabstmallgenieauthdevicestatusgetbycteiAPIResponseModel
}

// AlibabaailabstmallgenieauthdevicestatusgetbycteiAPIResponseModel is 根据ctei查询状态 成功返回结果
type AlibabaailabstmallgenieauthdevicestatusgetbycteiAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_tmallgenie_auth_device_status_getbyctei_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口对象封装
	Result *AlibabaailabstmallgenieauthdevicestatusgetbycteiResult `json:"result,omitempty" xml:"result,omitempty"`
}
