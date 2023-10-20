package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidletenderaftersaleorderperformAPIRequest 闲鱼帮卖售后订单履约 API请求
// alibaba.idle.tender.aftersale.order.perform
//
// 闲鱼帮卖售后订单履约
type AlibabaidletenderaftersaleorderperformAPIRequest struct {
	model.Params
	// 实参
	_tenderAfterSaleParam *TenderAfterSaleParam
}

// NewAlibabaidletenderaftersaleorderperformRequest 初始化AlibabaidletenderaftersaleorderperformAPIRequest对象
func NewAlibabaidletenderaftersaleorderperformRequest() *AlibabaidletenderaftersaleorderperformAPIRequest {
	return &AlibabaidletenderaftersaleorderperformAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidletenderaftersaleorderperformAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.tender.aftersale.order.perform"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidletenderaftersaleorderperformAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidletenderaftersaleorderperformAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenderAfterSaleParam is TenderAfterSaleParam Setter
// 实参
func (r *AlibabaidletenderaftersaleorderperformAPIRequest) SetTenderAfterSaleParam(_tenderAfterSaleParam *TenderAfterSaleParam) error {
	r._tenderAfterSaleParam = _tenderAfterSaleParam
	r.Set("tender_after_sale_param", _tenderAfterSaleParam)
	return nil
}

// GetTenderAfterSaleParam TenderAfterSaleParam Getter
func (r AlibabaidletenderaftersaleorderperformAPIRequest) GetTenderAfterSaleParam() *TenderAfterSaleParam {
	return r._tenderAfterSaleParam
}
