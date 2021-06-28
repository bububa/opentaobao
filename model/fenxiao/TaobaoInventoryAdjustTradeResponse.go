package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
交易库存调整单 APIResponse
taobao.inventory.adjust.trade

商家交易调整库存，淘宝交易、B2B经销等
*/
type TaobaoInventoryAdjustTradeAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoInventoryAdjustTradeResponse `json:"inventory_adjust_trade_response,omitempty"` 
    TaobaoInventoryAdjustTradeResponse
}

/* model for simplify = false
type TaobaoInventoryAdjustTradeResponse struct {

    // 操作返回码
    
    OperateCode   string `json:"operate_code,omitempty"`
    

    // 提示信息
    
    TipInfos  struct {
        TipInfo  []TipInfo `json:"tip_info,omitempty"`
    } `json:"tip_infos,omitempty"`
    

}
*/

type TaobaoInventoryAdjustTradeResponse struct {

    // 操作返回码
    OperateCode   string `json:"operate_code,omitempty"`

    // 提示信息
    TipInfos   []TipInfo `json:"tip_infos,omitempty"`

}
