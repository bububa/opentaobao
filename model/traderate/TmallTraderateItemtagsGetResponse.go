package traderate

import (
    "github.com/bububa/opentaobao/model"
)

/* 
通过商品ID获取标签列表 APIResponse
tmall.traderate.itemtags.get

通过商品ID获取标签详细信息
*/
type TmallTraderateItemtagsGetAPIResponse struct {
    model.CommonResponse
    Response *TmallTraderateItemtagsGetResponse `json:"tmall_traderate_itemtags_get_response,omitempty"`
}

type TmallTraderateItemtagsGetResponse struct {

    // 标签列表
    Tags   []TmallRateTagDetail `json:"tags,omitempty"`

}
