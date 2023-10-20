package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotradedrugconfirmorderAPIRequest 阿里健康020接单 API请求
// taobao.trade.drug.confirmorder
//
// 阿里健康020接单
type TaobaotradedrugconfirmorderAPIRequest struct {
	model.Params
	// 子账号nick
	_subUserNick string
	// 代送宝 代送商ID
	_deliveryId int64
	// public static int NORMAL_TYPE=0; 普通发货 默认 public static int DD_DAI_SONG=2; 代送宝	public static int DD_SONG_TYPE_V2=3; 点点送发货
	_confirmType int64
	// 订单ID
	_orderId int64
}

// NewTaobaotradedrugconfirmorderRequest 初始化TaobaotradedrugconfirmorderAPIRequest对象
func NewTaobaotradedrugconfirmorderRequest() *TaobaotradedrugconfirmorderAPIRequest {
	return &TaobaotradedrugconfirmorderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotradedrugconfirmorderAPIRequest) GetApiMethodName() string {
	return "taobao.trade.drug.confirmorder"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotradedrugconfirmorderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotradedrugconfirmorderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSubUserNick is SubUserNick Setter
// 子账号nick
func (r *TaobaotradedrugconfirmorderAPIRequest) SetSubUserNick(_subUserNick string) error {
	r._subUserNick = _subUserNick
	r.Set("sub_user_nick", _subUserNick)
	return nil
}

// GetSubUserNick SubUserNick Getter
func (r TaobaotradedrugconfirmorderAPIRequest) GetSubUserNick() string {
	return r._subUserNick
}

// SetDeliveryId is DeliveryId Setter
// 代送宝 代送商ID
func (r *TaobaotradedrugconfirmorderAPIRequest) SetDeliveryId(_deliveryId int64) error {
	r._deliveryId = _deliveryId
	r.Set("delivery_id", _deliveryId)
	return nil
}

// GetDeliveryId DeliveryId Getter
func (r TaobaotradedrugconfirmorderAPIRequest) GetDeliveryId() int64 {
	return r._deliveryId
}

// SetConfirmType is ConfirmType Setter
// public static int NORMAL_TYPE=0; 普通发货 默认 public static int DD_DAI_SONG=2; 代送宝	public static int DD_SONG_TYPE_V2=3; 点点送发货
func (r *TaobaotradedrugconfirmorderAPIRequest) SetConfirmType(_confirmType int64) error {
	r._confirmType = _confirmType
	r.Set("confirm_type", _confirmType)
	return nil
}

// GetConfirmType ConfirmType Getter
func (r TaobaotradedrugconfirmorderAPIRequest) GetConfirmType() int64 {
	return r._confirmType
}

// SetOrderId is OrderId Setter
// 订单ID
func (r *TaobaotradedrugconfirmorderAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaotradedrugconfirmorderAPIRequest) GetOrderId() int64 {
	return r._orderId
}
