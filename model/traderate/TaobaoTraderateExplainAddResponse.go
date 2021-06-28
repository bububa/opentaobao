package traderate

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商城评价解释接口 APIResponse
taobao.traderate.explain.add

商城卖家给评价做出解释（买家追加评论后、评价超过30天的，都不能再做评价解释）
*/
type TaobaoTraderateExplainAddAPIResponse struct {
    model.CommonResponse
    TaobaoTraderateExplainAddResponse
}

type TaobaoTraderateExplainAddResponse struct {
    XMLName xml.Name `xml:"traderate_explain_add_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 商城卖家给评价解释是否成功。
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
