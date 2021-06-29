package inventory

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
货品库存商家端调整 API返回值 
taobao.inventory.merchant.adjust

货品库存商家端调整 ，入库，出库，盘点
*/
type TaobaoInventoryMerchantAdjustAPIResponse struct {
    model.CommonResponse
    TaobaoInventoryMerchantAdjustResponse
}

// 货品库存商家端调整 成功返回结果
type TaobaoInventoryMerchantAdjustResponse struct {
    XMLName xml.Name `xml:"inventory_merchant_adjust_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *SingleResult `json:"result,omitempty" xml:"result,omitempty"`
}
