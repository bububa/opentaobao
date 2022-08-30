package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthSecondardNodeUserSyncAPIResponse 二级节点用户数据同步 API返回值
// alibaba.alihealth.secondard.node.user.sync
//
// 二级节点用户数据同步
type AlibabaAlihealthSecondardNodeUserSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthSecondardNodeUserSyncAPIResponseModel
}

// AlibabaAlihealthSecondardNodeUserSyncAPIResponseModel is 二级节点用户数据同步 成功返回结果
type AlibabaAlihealthSecondardNodeUserSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_secondard_node_user_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// model
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 是否成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}
