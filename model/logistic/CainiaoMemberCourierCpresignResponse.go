package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
cp清理离职用户信息 APIResponse
cainiao.member.courier.cpresign

CP清理内部离职的用户信息
*/
type CainiaoMemberCourierCpresignAPIResponse struct {
    model.CommonResponse
    // Response *CainiaoMemberCourierCpresignResponse `json:"cainiao_member_courier_cpresign_response,omitempty"` 
    CainiaoMemberCourierCpresignResponse
}

/* model for simplify = false
type CainiaoMemberCourierCpresignResponse struct {

    // 具体错误信息
    
    StatusMessage   string `json:"status_message,omitempty"`
    

    // 错误编码
    
    StatusCode   string `json:"status_code,omitempty"`
    

    // 业务处理是否成功
    
    Data   bool `json:"data,omitempty"`
    

    // 调用是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type CainiaoMemberCourierCpresignResponse struct {

    // 具体错误信息
    StatusMessage   string `json:"status_message,omitempty"`

    // 错误编码
    StatusCode   string `json:"status_code,omitempty"`

    // 业务处理是否成功
    Data   bool `json:"data,omitempty"`

    // 调用是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
