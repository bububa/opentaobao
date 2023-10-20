package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotrainpurchaseorderpayurlAPIRequest 火车票采购商接口-获取支付链接 API请求
// taobao.train.purchase.order.payurl
//
// 火车票采购商接口-获取支付链接
type TaobaotrainpurchaseorderpayurlAPIRequest struct {
	model.Params
	// 支付渠道(alipay-支付宝、union-银联、cmb-中国招商银行)
	_payChannel string
	// 支付类型(1-app,0-pc)，暂时只支持app
	_payType int64
	// 主订单id
	_mainBizOrderId int64
	// 订单业务类型(1-正向,2-改签)
	_businessType int64
	// 改签id,取自mtop.trip.train.createChangeOrder的响应{"headers":{},"httpStatusCode":200,"model":{"applyId":532585123123,"occupyChannelType":1,"occupyType":0},"success":true}
	_changeApplyId int64
}

// NewTaobaotrainpurchaseorderpayurlRequest 初始化TaobaotrainpurchaseorderpayurlAPIRequest对象
func NewTaobaotrainpurchaseorderpayurlRequest() *TaobaotrainpurchaseorderpayurlAPIRequest {
	return &TaobaotrainpurchaseorderpayurlAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotrainpurchaseorderpayurlAPIRequest) GetApiMethodName() string {
	return "taobao.train.purchase.order.payurl"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotrainpurchaseorderpayurlAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotrainpurchaseorderpayurlAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPayChannel is PayChannel Setter
// 支付渠道(alipay-支付宝、union-银联、cmb-中国招商银行)
func (r *TaobaotrainpurchaseorderpayurlAPIRequest) SetPayChannel(_payChannel string) error {
	r._payChannel = _payChannel
	r.Set("pay_channel", _payChannel)
	return nil
}

// GetPayChannel PayChannel Getter
func (r TaobaotrainpurchaseorderpayurlAPIRequest) GetPayChannel() string {
	return r._payChannel
}

// SetPayType is PayType Setter
// 支付类型(1-app,0-pc)，暂时只支持app
func (r *TaobaotrainpurchaseorderpayurlAPIRequest) SetPayType(_payType int64) error {
	r._payType = _payType
	r.Set("pay_type", _payType)
	return nil
}

// GetPayType PayType Getter
func (r TaobaotrainpurchaseorderpayurlAPIRequest) GetPayType() int64 {
	return r._payType
}

// SetMainBizOrderId is MainBizOrderId Setter
// 主订单id
func (r *TaobaotrainpurchaseorderpayurlAPIRequest) SetMainBizOrderId(_mainBizOrderId int64) error {
	r._mainBizOrderId = _mainBizOrderId
	r.Set("main_biz_order_id", _mainBizOrderId)
	return nil
}

// GetMainBizOrderId MainBizOrderId Getter
func (r TaobaotrainpurchaseorderpayurlAPIRequest) GetMainBizOrderId() int64 {
	return r._mainBizOrderId
}

// SetBusinessType is BusinessType Setter
// 订单业务类型(1-正向,2-改签)
func (r *TaobaotrainpurchaseorderpayurlAPIRequest) SetBusinessType(_businessType int64) error {
	r._businessType = _businessType
	r.Set("business_type", _businessType)
	return nil
}

// GetBusinessType BusinessType Getter
func (r TaobaotrainpurchaseorderpayurlAPIRequest) GetBusinessType() int64 {
	return r._businessType
}

// SetChangeApplyId is ChangeApplyId Setter
// 改签id,取自mtop.trip.train.createChangeOrder的响应{&#34;headers&#34;:{},&#34;httpStatusCode&#34;:200,&#34;model&#34;:{&#34;applyId&#34;:532585123123,&#34;occupyChannelType&#34;:1,&#34;occupyType&#34;:0},&#34;success&#34;:true}
func (r *TaobaotrainpurchaseorderpayurlAPIRequest) SetChangeApplyId(_changeApplyId int64) error {
	r._changeApplyId = _changeApplyId
	r.Set("change_apply_id", _changeApplyId)
	return nil
}

// GetChangeApplyId ChangeApplyId Getter
func (r TaobaotrainpurchaseorderpayurlAPIRequest) GetChangeApplyId() int64 {
	return r._changeApplyId
}
