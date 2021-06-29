package wlb

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家库存调整 API返回值 
cainiao.merchant.inventory.adjust

商家仓库存调整接口，目前仅支持全量更新
*/
type CainiaoMerchantInventoryAdjustAPIResponse struct {
    model.CommonResponse
    CainiaoMerchantInventoryAdjustResponse
}

// 商家库存调整 成功返回结果
type CainiaoMerchantInventoryAdjustResponse struct {
    XMLName xml.Name `xml:"cainiao_merchant_inventory_adjust_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *SingleResultDTO `json:"result,omitempty" xml:"result,omitempty"`
}
