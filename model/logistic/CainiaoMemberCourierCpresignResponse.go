package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
cp清理离职用户信息 API返回值 
cainiao.member.courier.cpresign

CP清理内部离职的用户信息
*/
type CainiaoMemberCourierCpresignAPIResponse struct {
    model.CommonResponse
    CainiaoMemberCourierCpresignResponse
}

// cp清理离职用户信息 成功返回结果
type CainiaoMemberCourierCpresignResponse struct {
    XMLName xml.Name `xml:"cainiao_member_courier_cpresign_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 具体错误信息
    StatusMessage   string `json:"status_message,omitempty" xml:"status_message,omitempty"`
    // 错误编码
    StatusCode   string `json:"status_code,omitempty" xml:"status_code,omitempty"`
    // 业务处理是否成功
    Data   bool `json:"data,omitempty" xml:"data,omitempty"`
    // 调用是否成功
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
