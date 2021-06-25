package traderate

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商城评价解释接口 APIResponse
taobao.traderate.explain.add

商城卖家给评价做出解释（买家追加评论后、评价超过30天的，都不能再做评价解释）
*/
type TaobaoTraderateExplainAddAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTraderateExplainAddResponse `json:"taobao_traderate_explain_add_response,omitempty"`
}

type TaobaoTraderateExplainAddResponse struct {

    // 商城卖家给评价解释是否成功。
    IsSuccess   bool `json:"is_success,omitempty"`

}
