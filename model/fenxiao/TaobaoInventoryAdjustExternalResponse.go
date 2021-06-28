package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
非交易库存调整单 APIResponse
taobao.inventory.adjust.external

建议使用新接口：taobao.inventory.merchant.adjust ，该接口会逐步停用。
商家非交易调整库存，调拨出库、盘点等时调用
*/
type TaobaoInventoryAdjustExternalAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoInventoryAdjustExternalResponse `json:"inventory_adjust_external_response,omitempty"` 
    TaobaoInventoryAdjustExternalResponse
}

/* model for simplify = false
type TaobaoInventoryAdjustExternalResponse struct {

    // 操作返回码
    
    OperateCode   string `json:"operate_code,omitempty"`
    

    // 提示信息
    
    TipInfos  struct {
        TipInfo  []TipInfo `json:"tip_info,omitempty"`
    } `json:"tip_infos,omitempty"`
    

}
*/

type TaobaoInventoryAdjustExternalResponse struct {

    // 操作返回码
    OperateCode   string `json:"operate_code,omitempty"`

    // 提示信息
    TipInfos   []TipInfo `json:"tip_infos,omitempty"`

}
