package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaomodifyorderordercheckAPIRequest 自助改单服务发货订单校验 API请求
// taobao.modifyorder.order.check
//
// 自助改单服务发货后订单回传接口
type TaobaomodifyorderordercheckAPIRequest struct {
	model.Params
	// 商品skuId
	_finalSkuId string
	// 订单id
	_bizOrderId string
	// 商品skuId
	_finalOuterId string
	// 子订单id
	_subBizOrderId string
	// 地址oaid
	_oaid string
}

// NewTaobaomodifyorderordercheckRequest 初始化TaobaomodifyorderordercheckAPIRequest对象
func NewTaobaomodifyorderordercheckRequest() *TaobaomodifyorderordercheckAPIRequest {
	return &TaobaomodifyorderordercheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaomodifyorderordercheckAPIRequest) GetApiMethodName() string {
	return "taobao.modifyorder.order.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaomodifyorderordercheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaomodifyorderordercheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFinalSkuId is FinalSkuId Setter
// 商品skuId
func (r *TaobaomodifyorderordercheckAPIRequest) SetFinalSkuId(_finalSkuId string) error {
	r._finalSkuId = _finalSkuId
	r.Set("final_sku_id", _finalSkuId)
	return nil
}

// GetFinalSkuId FinalSkuId Getter
func (r TaobaomodifyorderordercheckAPIRequest) GetFinalSkuId() string {
	return r._finalSkuId
}

// SetBizOrderId is BizOrderId Setter
// 订单id
func (r *TaobaomodifyorderordercheckAPIRequest) SetBizOrderId(_bizOrderId string) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r TaobaomodifyorderordercheckAPIRequest) GetBizOrderId() string {
	return r._bizOrderId
}

// SetFinalOuterId is FinalOuterId Setter
// 商品skuId
func (r *TaobaomodifyorderordercheckAPIRequest) SetFinalOuterId(_finalOuterId string) error {
	r._finalOuterId = _finalOuterId
	r.Set("final_outer_id", _finalOuterId)
	return nil
}

// GetFinalOuterId FinalOuterId Getter
func (r TaobaomodifyorderordercheckAPIRequest) GetFinalOuterId() string {
	return r._finalOuterId
}

// SetSubBizOrderId is SubBizOrderId Setter
// 子订单id
func (r *TaobaomodifyorderordercheckAPIRequest) SetSubBizOrderId(_subBizOrderId string) error {
	r._subBizOrderId = _subBizOrderId
	r.Set("sub_biz_order_id", _subBizOrderId)
	return nil
}

// GetSubBizOrderId SubBizOrderId Getter
func (r TaobaomodifyorderordercheckAPIRequest) GetSubBizOrderId() string {
	return r._subBizOrderId
}

// SetOaid is Oaid Setter
// 地址oaid
func (r *TaobaomodifyorderordercheckAPIRequest) SetOaid(_oaid string) error {
	r._oaid = _oaid
	r.Set("oaid", _oaid)
	return nil
}

// GetOaid Oaid Getter
func (r TaobaomodifyorderordercheckAPIRequest) GetOaid() string {
	return r._oaid
}
