package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
ISV报名官方活动(中心化流量) APIResponse
alibaba.interact.activity.apply

支持商家将使用isv创建的活动所对应的权益信息同步到手淘，供过滤是否在中心化流量入口透出
*/
type AlibabaInteractActivityApplyAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaInteractActivityApplyResponse `json:"alibaba_interact_activity_apply_response,omitempty"` 
    AlibabaInteractActivityApplyResponse
}

/* model for simplify = false
type AlibabaInteractActivityApplyResponse struct {

    // 服务结果对象
    
    Data  *struct {
        ActivityWriteResult  *ActivityWriteResult `json:"activity_write_result,omitempty"`
    } `json:"data,omitempty"`
    

    // top接口执行成功与否
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // 出错提示信息
    
    ErrMsg   string `json:"err_msg,omitempty"`
    

}
*/

type AlibabaInteractActivityApplyResponse struct {

    // 服务结果对象
    Data   *ActivityWriteResult `json:"data,omitempty"`

    // top接口执行成功与否
    IsSuccess   bool `json:"is_success,omitempty"`

    // 出错提示信息
    ErrMsg   string `json:"err_msg,omitempty"`

}
