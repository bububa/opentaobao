package travel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【API3.0】套餐级别日历价格库存增删操作 APIResponse
taobao.alitrip.travel.item.sku.package.modify

【API3.0】套餐级别日历价格库存增删操作
*/
type TaobaoAlitripTravelItemSkuPackageModifyAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripTravelItemSkuPackageModifyResponse
}

type TaobaoAlitripTravelItemSkuPackageModifyResponse struct {
    XMLName xml.Name `xml:"alitrip_travel_item_sku_package_modify_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 商品sku修改结果
    
    TravelItem   *TopTravelItem `json:"travel_item,omitempty" xml:"travel_item,omitempty"`

    
}
