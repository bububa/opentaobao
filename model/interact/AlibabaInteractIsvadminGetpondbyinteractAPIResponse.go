package interact

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractIsvadminGetpondbyinteractAPIResponse 根据互动实例ID查询奖池信息 API返回值
// alibaba.interact.isvadmin.getpondbyinteract
//
// 根据互动实例ID查询奖池信息
type AlibabaInteractIsvadminGetpondbyinteractAPIResponse struct {
	model.CommonResponse
	AlibabaInteractIsvadminGetpondbyinteractAPIResponseModel
}

// AlibabaInteractIsvadminGetpondbyinteractAPIResponseModel is 根据互动实例ID查询奖池信息 成功返回结果
type AlibabaInteractIsvadminGetpondbyinteractAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_isvadmin_getpondbyinteract_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 奖池信息
	Data *PrizePondDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否调用成功
	Succ bool `json:"succ,omitempty" xml:"succ,omitempty"`
	// 调用错误原因
	Error string `json:"error,omitempty" xml:"error,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}
