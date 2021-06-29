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
    TaobaoInventoryAdjustTradeResponse
}

type TaobaoInventoryAdjustTradeResponse struct {
    XMLName xml.Name `xml:"inventory_adjust_trade_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 操作返回码
    
    OperateCode   string `json:"operate_code,omitempty" xml:"operate_code,omitempty"`

    
    // 提示信息
    
    TipInfos   []TipInfo `json:"tip_infos,omitempty" xml:"tip_infos>tip_info,omitempty"`
    
    
}
