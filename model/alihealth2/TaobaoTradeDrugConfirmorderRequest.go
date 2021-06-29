package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里健康020接单 API请求
taobao.trade.drug.confirmorder

阿里健康020接单
*/
type TaobaoTradeDrugConfirmorderRequest struct {
    model.Params
    // 代送宝 代送商ID
    deliveryId   int64
    // public static int NORMAL_TYPE=0; 普通发货 默认 public static int DD_DAI_SONG=2; 代送宝	public static int DD_SONG_TYPE_V2=3; 点点送发货
    confirmType   int64
    // 订单ID
    orderId   int64
    // 子账号nick
    subUserNick   string
}

// 初始化TaobaoTradeDrugConfirmorderRequest对象
func NewTaobaoTradeDrugConfirmorderRequest() *TaobaoTradeDrugConfirmorderRequest{
    return &TaobaoTradeDrugConfirmorderRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTradeDrugConfirmorderRequest) GetApiMethodName() string {
    return "taobao.trade.drug.confirmorder"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTradeDrugConfirmorderRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeliveryId Setter
// 代送宝 代送商ID
func (r *TaobaoTradeDrugConfirmorderRequest) SetDeliveryId(deliveryId int64) error {
    r.deliveryId = deliveryId
    r.Set("delivery_id", deliveryId)
    return nil
}

// DeliveryId Getter
func (r TaobaoTradeDrugConfirmorderRequest) GetDeliveryId() int64 {
    return r.deliveryId
}
// ConfirmType Setter
// public static int NORMAL_TYPE=0; 普通发货 默认 public static int DD_DAI_SONG=2; 代送宝	public static int DD_SONG_TYPE_V2=3; 点点送发货
func (r *TaobaoTradeDrugConfirmorderRequest) SetConfirmType(confirmType int64) error {
    r.confirmType = confirmType
    r.Set("confirm_type", confirmType)
    return nil
}

// ConfirmType Getter
func (r TaobaoTradeDrugConfirmorderRequest) GetConfirmType() int64 {
    return r.confirmType
}
// OrderId Setter
// 订单ID
func (r *TaobaoTradeDrugConfirmorderRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r TaobaoTradeDrugConfirmorderRequest) GetOrderId() int64 {
    return r.orderId
}
// SubUserNick Setter
// 子账号nick
func (r *TaobaoTradeDrugConfirmorderRequest) SetSubUserNick(subUserNick string) error {
    r.subUserNick = subUserNick
    r.Set("sub_user_nick", subUserNick)
    return nil
}

// SubUserNick Getter
func (r TaobaoTradeDrugConfirmorderRequest) GetSubUserNick() string {
    return r.subUserNick
}
