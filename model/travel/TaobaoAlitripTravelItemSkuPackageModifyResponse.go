package travel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【API3.0】套餐级别日历价格库存增删操作 API返回值 
taobao.alitrip.travel.item.sku.package.modify

【API3.0】套餐级别日历价格库存增删操作
*/
type TaobaoAlitripTravelItemSkuPackageModifyAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripTravelItemSkuPackageModifyResponse
}

// 【API3.0】套餐级别日历价格库存增删操作 成功返回结果
type TaobaoAlitripTravelItemSkuPackageModifyResponse struct {
    XMLName xml.Name `xml:"alitrip_travel_item_sku_package_modify_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 商品sku修改结果
    TravelItem   *TopTravelItem `json:"travel_item,omitempty" xml:"travel_item,omitempty"`
}
