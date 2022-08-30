package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthSecondardNodeUserUpdatestatusAPIResponse 二级节点用户注销 API返回值
// alibaba.alihealth.secondard.node.user.updatestatus
//
// 注销二级节点用户
type AlibabaAlihealthSecondardNodeUserUpdatestatusAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthSecondardNodeUserUpdatestatusAPIResponseModel
}

// AlibabaAlihealthSecondardNodeUserUpdatestatusAPIResponseModel is 二级节点用户注销 成功返回结果
type AlibabaAlihealthSecondardNodeUserUpdatestatusAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_secondard_node_user_updatestatus_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口调用返回描述
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 接口调用返回标识
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 接口调用是否成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
	// 业务逻辑处理是否成功
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}
