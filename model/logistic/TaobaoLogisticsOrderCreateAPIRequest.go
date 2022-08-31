package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsOrderCreateAPIRequest 创建物流订单 API请求
// taobao.logistics.order.create
//
// 用户调用该接口可以创建物流订单。目前仅支持手工订单的创建，创建完毕后默认自动使用“自己联系”的方式发货并且初始状态为”已发货“。也可以通过可选参数选择是否发货以及何种方式进行发货。
type TaobaoLogisticsOrderCreateAPIRequest struct {
	model.Params
	// 运送货物的单价列表(注意：单位为分），用|号隔开
	_itemValues string
	// 运送的货物名称列表，用|号隔开
	_goodsNames string
	// 运送货物的数量列表，用|号隔开
	_goodsQuantities string
	// 发货的物流公司运单号。在logis_type=OFFLINE且is_consign=true时，此项必填。
	_mailNo string
	// 卖家旺旺号
	_sellerWangwangId string
	// 发货的物流公司代码，如申通=STO，圆通=YTO。is_consign=true时，此项必填。
	_logisCompanyCode string
	// 发货方式,默认为自己联系发货。可选值:ONLINE(在线下单)、OFFLINE(自己联系)。
	_logisType string
	// 物流发货类型：1-普通订单<br/>不填为默认类型 1-普通订单
	_orderType int64
	// 订单的交易号码
	_tradeId int64
	// 运费承担方式。1为买家承担运费，2为卖家承担运费，其他值为错误参数。
	_shipping int64
	// 创建订单同时是否进行发货，默认发货。
	_isConsign bool
}

// NewTaobaoLogisticsOrderCreateRequest 初始化TaobaoLogisticsOrderCreateAPIRequest对象
func NewTaobaoLogisticsOrderCreateRequest() *TaobaoLogisticsOrderCreateAPIRequest {
	return &TaobaoLogisticsOrderCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsOrderCreateAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsOrderCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetItemValues is ItemValues Setter
// 运送货物的单价列表(注意：单位为分），用|号隔开
func (r *TaobaoLogisticsOrderCreateAPIRequest) SetItemValues(_itemValues string) error {
	r._itemValues = _itemValues
	r.Set("item_values", _itemValues)
	return nil
}

// GetItemValues ItemValues Getter
func (r TaobaoLogisticsOrderCreateAPIRequest) GetItemValues() string {
	return r._itemValues
}

// SetGoodsNames is GoodsNames Setter
// 运送的货物名称列表，用|号隔开
func (r *TaobaoLogisticsOrderCreateAPIRequest) SetGoodsNames(_goodsNames string) error {
	r._goodsNames = _goodsNames
	r.Set("goods_names", _goodsNames)
	return nil
}

// GetGoodsNames GoodsNames Getter
func (r TaobaoLogisticsOrderCreateAPIRequest) GetGoodsNames() string {
	return r._goodsNames
}

// SetGoodsQuantities is GoodsQuantities Setter
// 运送货物的数量列表，用|号隔开
func (r *TaobaoLogisticsOrderCreateAPIRequest) SetGoodsQuantities(_goodsQuantities string) error {
	r._goodsQuantities = _goodsQuantities
	r.Set("goods_quantities", _goodsQuantities)
	return nil
}

// GetGoodsQuantities GoodsQuantities Getter
func (r TaobaoLogisticsOrderCreateAPIRequest) GetGoodsQuantities() string {
	return r._goodsQuantities
}

// SetMailNo is MailNo Setter
// 发货的物流公司运单号。在logis_type=OFFLINE且is_consign=true时，此项必填。
func (r *TaobaoLogisticsOrderCreateAPIRequest) SetMailNo(_mailNo string) error {
	r._mailNo = _mailNo
	r.Set("mail_no", _mailNo)
	return nil
}

// GetMailNo MailNo Getter
func (r TaobaoLogisticsOrderCreateAPIRequest) GetMailNo() string {
	return r._mailNo
}

// SetSellerWangwangId is SellerWangwangId Setter
// 卖家旺旺号
func (r *TaobaoLogisticsOrderCreateAPIRequest) SetSellerWangwangId(_sellerWangwangId string) error {
	r._sellerWangwangId = _sellerWangwangId
	r.Set("seller_wangwang_id", _sellerWangwangId)
	return nil
}

// GetSellerWangwangId SellerWangwangId Getter
func (r TaobaoLogisticsOrderCreateAPIRequest) GetSellerWangwangId() string {
	return r._sellerWangwangId
}

// SetLogisCompanyCode is LogisCompanyCode Setter
// 发货的物流公司代码，如申通=STO，圆通=YTO。is_consign=true时，此项必填。
func (r *TaobaoLogisticsOrderCreateAPIRequest) SetLogisCompanyCode(_logisCompanyCode string) error {
	r._logisCompanyCode = _logisCompanyCode
	r.Set("logis_company_code", _logisCompanyCode)
	return nil
}

// GetLogisCompanyCode LogisCompanyCode Getter
func (r TaobaoLogisticsOrderCreateAPIRequest) GetLogisCompanyCode() string {
	return r._logisCompanyCode
}

// SetLogisType is LogisType Setter
// 发货方式,默认为自己联系发货。可选值:ONLINE(在线下单)、OFFLINE(自己联系)。
func (r *TaobaoLogisticsOrderCreateAPIRequest) SetLogisType(_logisType string) error {
	r._logisType = _logisType
	r.Set("logis_type", _logisType)
	return nil
}

// GetLogisType LogisType Getter
func (r TaobaoLogisticsOrderCreateAPIRequest) GetLogisType() string {
	return r._logisType
}

// SetOrderType is OrderType Setter
// 物流发货类型：1-普通订单&lt;br/&gt;不填为默认类型 1-普通订单
func (r *TaobaoLogisticsOrderCreateAPIRequest) SetOrderType(_orderType int64) error {
	r._orderType = _orderType
	r.Set("order_type", _orderType)
	return nil
}

// GetOrderType OrderType Getter
func (r TaobaoLogisticsOrderCreateAPIRequest) GetOrderType() int64 {
	return r._orderType
}

// SetTradeId is TradeId Setter
// 订单的交易号码
func (r *TaobaoLogisticsOrderCreateAPIRequest) SetTradeId(_tradeId int64) error {
	r._tradeId = _tradeId
	r.Set("trade_id", _tradeId)
	return nil
}

// GetTradeId TradeId Getter
func (r TaobaoLogisticsOrderCreateAPIRequest) GetTradeId() int64 {
	return r._tradeId
}

// SetShipping is Shipping Setter
// 运费承担方式。1为买家承担运费，2为卖家承担运费，其他值为错误参数。
func (r *TaobaoLogisticsOrderCreateAPIRequest) SetShipping(_shipping int64) error {
	r._shipping = _shipping
	r.Set("shipping", _shipping)
	return nil
}

// GetShipping Shipping Getter
func (r TaobaoLogisticsOrderCreateAPIRequest) GetShipping() int64 {
	return r._shipping
}

// SetIsConsign is IsConsign Setter
// 创建订单同时是否进行发货，默认发货。
func (r *TaobaoLogisticsOrderCreateAPIRequest) SetIsConsign(_isConsign bool) error {
	r._isConsign = _isConsign
	r.Set("is_consign", _isConsign)
	return nil
}

// GetIsConsign IsConsign Getter
func (r TaobaoLogisticsOrderCreateAPIRequest) GetIsConsign() bool {
	return r._isConsign
}
