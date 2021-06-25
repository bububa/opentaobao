package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
交易库存调整单 APIRequest
taobao.inventory.adjust.trade

商家交易调整库存，淘宝交易、B2B经销等
*/
type TaobaoInventoryAdjustTradeRequest struct {
    model.Params

    // 订单类型：B2C、B2B
    tbOrderType   string 

    // 业务操作时间
    operateTime   string 

    // 商家外部定单号
    bizUniqueCode   string 

    // 商品初始库存信息： [{ "TBOrderCode”:”淘宝交易号”,"TBSubOrderCode ":"淘宝子交易单号,赠品可以不填","”isGift”:”TRUE或者FALSE,是否赠品”,storeCode":"商家仓库编码"," scItemId ":"商品后端ID","scItemCode":"商品商家编码"," originScItemId ":"原商品ID","inventoryType":"","quantity":"111","isComplete":"TRUE或者FALSE，是否全部确认出库"}]
    items   string 

}

func NewTaobaoInventoryAdjustTradeRequest() *TaobaoInventoryAdjustTradeRequest{
    return &TaobaoInventoryAdjustTradeRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoInventoryAdjustTradeRequest) GetApiMethodName() string {
    return "taobao.inventory.adjust.trade"
}

func (r TaobaoInventoryAdjustTradeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoInventoryAdjustTradeRequest) SetTbOrderType(tbOrderType string) error {
    r.tbOrderType = tbOrderType
    r.Set("tb_order_type", tbOrderType)
    return nil
}

func (r TaobaoInventoryAdjustTradeRequest) GetTbOrderType() string {
    return r.tbOrderType
}

func (r *TaobaoInventoryAdjustTradeRequest) SetOperateTime(operateTime string) error {
    r.operateTime = operateTime
    r.Set("operate_time", operateTime)
    return nil
}

func (r TaobaoInventoryAdjustTradeRequest) GetOperateTime() string {
    return r.operateTime
}

func (r *TaobaoInventoryAdjustTradeRequest) SetBizUniqueCode(bizUniqueCode string) error {
    r.bizUniqueCode = bizUniqueCode
    r.Set("biz_unique_code", bizUniqueCode)
    return nil
}

func (r TaobaoInventoryAdjustTradeRequest) GetBizUniqueCode() string {
    return r.bizUniqueCode
}

func (r *TaobaoInventoryAdjustTradeRequest) SetItems(items string) error {
    r.items = items
    r.Set("items", items)
    return nil
}

func (r TaobaoInventoryAdjustTradeRequest) GetItems() string {
    return r.items
}

