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
type TaobaoTradeDrugConfirmorderAPIRequest struct {
    model.Params
    // 代送宝 代送商ID
    _deliveryId   int64
    // public static int NORMAL_TYPE=0; 普通发货 默认 public static int DD_DAI_SONG=2; 代送宝	public static int DD_SONG_TYPE_V2=3; 点点送发货
    _confirmType   int64
    // 订单ID
    _orderId   int64
    // 子账号nick
    _subUserNick   string
}

// 初始化TaobaoTradeDrugConfirmorderAPIRequest对象
func NewTaobaoTradeDrugConfirmorderRequest() *TaobaoTradeDrugConfirmorderAPIRequest{
    return &TaobaoTradeDrugConfirmorderAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTradeDrugConfirmorderAPIRequest) GetApiMethodName() string {
    return "taobao.trade.drug.confirmorder"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTradeDrugConfirmorderAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeliveryId Setter
// 代送宝 代送商ID
func (r *TaobaoTradeDrugConfirmorderAPIRequest) SetDeliveryId(_deliveryId int64) error {
    r._deliveryId = _deliveryId
    r.Set("delivery_id", _deliveryId)
    return nil
}

// DeliveryId Getter
func (r TaobaoTradeDrugConfirmorderAPIRequest) GetDeliveryId() int64 {
    return r._deliveryId
}
// ConfirmType Setter
// public static int NORMAL_TYPE=0; 普通发货 默认 public static int DD_DAI_SONG=2; 代送宝	public static int DD_SONG_TYPE_V2=3; 点点送发货
func (r *TaobaoTradeDrugConfirmorderAPIRequest) SetConfirmType(_confirmType int64) error {
    r._confirmType = _confirmType
    r.Set("confirm_type", _confirmType)
    return nil
}

// ConfirmType Getter
func (r TaobaoTradeDrugConfirmorderAPIRequest) GetConfirmType() int64 {
    return r._confirmType
}
// OrderId Setter
// 订单ID
func (r *TaobaoTradeDrugConfirmorderAPIRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r TaobaoTradeDrugConfirmorderAPIRequest) GetOrderId() int64 {
    return r._orderId
}
// SubUserNick Setter
// 子账号nick
func (r *TaobaoTradeDrugConfirmorderAPIRequest) SetSubUserNick(_subUserNick string) error {
    r._subUserNick = _subUserNick
    r.Set("sub_user_nick", _subUserNick)
    return nil
}

// SubUserNick Getter
func (r TaobaoTradeDrugConfirmorderAPIRequest) GetSubUserNick() string {
    return r._subUserNick
}
