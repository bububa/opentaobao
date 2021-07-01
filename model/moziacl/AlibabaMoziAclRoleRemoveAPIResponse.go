package moziacl

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除角色 API返回值 
alibaba.mozi.acl.role.remove

根据传入的角色code、租户id，删除租户内对应的角色
*/
type AlibabaMoziAclRoleRemoveAPIResponse struct {
    model.CommonResponse
    AlibabaMoziAclRoleRemoveAPIResponseModel
}

// 删除角色 成功返回结果
type AlibabaMoziAclRoleRemoveAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_mozi_acl_role_remove_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 无值
    Data   string `json:"data,omitempty" xml:"data,omitempty"`
    // 是否操作成功,true代表操作成功
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
    // 请求id
    MoziRequestId   string `json:"mozi_request_id,omitempty" xml:"mozi_request_id,omitempty"`
    // 如果success不为true，则自此段返回详细的错误信息
    ResponseMessage   string `json:"response_message,omitempty" xml:"response_message,omitempty"`
    // 如果success为true，则返回0，否则此段返回详细的错误code
    ResponseCode   string `json:"response_code,omitempty" xml:"response_code,omitempty"`
}
