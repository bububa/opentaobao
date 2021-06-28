package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
UMP工具检测 APIResponse
taobao.promotionmisc.tool.check

UMP工具检测。ISV通过该接口检测（通过taobao.ump.tool.add）创建的UMP工具（tool）是否符合规范，如果不符合，则返回错误信息和对应的解决方案的；工具检测通过后才可以提交工具审核邮件，提交工具审核时，需提供该接口的返回值。
*/
type TaobaoPromotionmiscToolCheckAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoPromotionmiscToolCheckResponse `json:"promotionmisc_tool_check_response,omitempty"` 
    TaobaoPromotionmiscToolCheckResponse
}

/* model for simplify = false
type TaobaoPromotionmiscToolCheckResponse struct {

    // 工具审核结果。
    
    CheckToolModule  *struct {
        CheckToolModule  *CheckToolModule `json:"check_tool_module,omitempty"`
    } `json:"check_tool_module,omitempty"`
    

    // 工具检测动作是否成功。检测是否通过请查看返回值default_model里的is_pass。
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // 接口调用错误信息描述。
    
    ErrorMessage   string `json:"error_message,omitempty"`
    

}
*/

type TaobaoPromotionmiscToolCheckResponse struct {

    // 工具审核结果。
    CheckToolModule   *CheckToolModule `json:"check_tool_module,omitempty"`

    // 工具检测动作是否成功。检测是否通过请查看返回值default_model里的is_pass。
    IsSuccess   bool `json:"is_success,omitempty"`

    // 接口调用错误信息描述。
    ErrorMessage   string `json:"error_message,omitempty"`

}
