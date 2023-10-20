package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamjpresalesettlementaddlistAPIRequest 预售结算数据回传 API请求
// alibaba.mj.presale.settlement.addlist
//
// 用于预售活动结算数据的回传。
type AlibabamjpresalesettlementaddlistAPIRequest struct {
	model.Params
	// 订单json格式数据
	_preSaleRefundJson string
}

// NewAlibabamjpresalesettlementaddlistRequest 初始化AlibabamjpresalesettlementaddlistAPIRequest对象
func NewAlibabamjpresalesettlementaddlistRequest() *AlibabamjpresalesettlementaddlistAPIRequest {
	return &AlibabamjpresalesettlementaddlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamjpresalesettlementaddlistAPIRequest) GetApiMethodName() string {
	return "alibaba.mj.presale.settlement.addlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamjpresalesettlementaddlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamjpresalesettlementaddlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPreSaleRefundJson is PreSaleRefundJson Setter
// 订单json格式数据
func (r *AlibabamjpresalesettlementaddlistAPIRequest) SetPreSaleRefundJson(_preSaleRefundJson string) error {
	r._preSaleRefundJson = _preSaleRefundJson
	r.Set("pre_sale_refund_json", _preSaleRefundJson)
	return nil
}

// GetPreSaleRefundJson PreSaleRefundJson Getter
func (r AlibabamjpresalesettlementaddlistAPIRequest) GetPreSaleRefundJson() string {
	return r._preSaleRefundJson
}
