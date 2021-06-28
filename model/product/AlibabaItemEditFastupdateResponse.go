package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商品编辑增量更新 APIResponse
alibaba.item.edit.fastupdate

商品编辑增量更新;
<br/>该接口编辑sku，只能更新价格、库存等信息，不能新增sku;
<br/>新增sku用全量接口alibaba.item.edit.submit，先设置销售属性;
*/
type AlibabaItemEditFastupdateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaItemEditFastupdateResponse `json:"alibaba_item_edit_fastupdate_response,omitempty"` 
    AlibabaItemEditFastupdateResponse
}

/* model for simplify = false
type AlibabaItemEditFastupdateResponse struct {

    // 商品更新时间
    
    UpdateTime   string `json:"update_time,omitempty"`
    

    // 商品ID
    
    ItemId   int64 `json:"item_id,omitempty"`
    

    // 商品所属市场
    
    Market   string `json:"market,omitempty"`
    

}
*/

type AlibabaItemEditFastupdateResponse struct {

    // 商品更新时间
    UpdateTime   string `json:"update_time,omitempty"`

    // 商品ID
    ItemId   int64 `json:"item_id,omitempty"`

    // 商品所属市场
    Market   string `json:"market,omitempty"`

}
