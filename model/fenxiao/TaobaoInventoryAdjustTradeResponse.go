package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
交易库存调整单 APIResponse
taobao.inventory.adjust.trade

商家交易调整库存，淘宝交易、B2B经销等
*/
type TaobaoInventoryAdjustTradeAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"inventory_adjust_trade_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 操作返回码
    
    OperateCode   string `json:"operate_code,omitempty" xml:"