package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫商品/SKU价格更新接口 APIResponse
tmall.item.price.update

天猫商品/SKU价格更新接口，支持商品、SKU价格同时更新，支持同一商品下的SKU批量更新。
*/
type TmallItemPriceUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_item_price_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 价格更新结果
    
    PriceUpdateResult   string `json:"price_update_result,omitempty" xml:"