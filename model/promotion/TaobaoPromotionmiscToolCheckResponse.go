package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
UMP工具检测 APIResponse
taobao.promotionmisc.tool.check

UMP工具检测。ISV通过该接口检测（通过taobao.ump.tool.add）创建的UMP工具（tool）是否符合规范，如果不符合，则返回错误信息和对应的解决方案的；工具检测通过后才可以提交工具审核邮件，提交工具审核时，需提供该接口的返回值。
*/
type TaobaoPromotionmiscToolCheckAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"promotionmisc_tool_check_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 工具审核结果。
    
    CheckToolModule   *CheckToolModule `json:"check_tool_module,omitempty" xml:"