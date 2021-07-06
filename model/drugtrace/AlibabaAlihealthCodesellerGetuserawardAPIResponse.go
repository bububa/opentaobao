package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthCodesellerGetuserawardAPIResponse 贩卖机扫码查询领奖状态 API返回值
// alibaba.alihealth.codeseller.getuseraward
//
// 贩卖机扫码查询领奖状态
type AlibabaAlihealthCodesellerGetuserawardAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthCodesellerGetuserawardAPIResponseModel
}

// AlibabaAlihealthCodesellerGetuserawardAPIResponseModel is 贩卖机扫码查询领奖状态 成功返回结果
type AlibabaAlihealthCodesellerGetuserawardAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_codeseller_getuseraward_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 状态码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 状态值
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 是否领取奖品 true:已领取 false:未领取 null:未知
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 响应标识
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}
