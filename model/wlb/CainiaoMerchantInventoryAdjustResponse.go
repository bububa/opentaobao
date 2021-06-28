package wlb

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家库存调整 APIResponse
cainiao.merchant.inventory.adjust

商家仓库存调整接口，目前仅支持全量更新
*/
type CainiaoMerchantInventoryAdjustAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"cainiao_merchant_inventory_adjust_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *SingleResultDto `json:"result,omitempty" xml:"