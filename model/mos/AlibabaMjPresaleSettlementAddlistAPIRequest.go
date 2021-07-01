package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMjPresaleSettlementAddlistAPIRequest
预售结算数据回传 API请求
alibaba.mj.presale.settlement.addlist

用于预售活动结算数据的回传。 */
type AlibabaMjPresaleSettlementAddlistAPIRequest struct {
	model.Params
	// 订单json格式数据
	_preSaleRefundJson string
}

// NewAlibabaMjPresaleSettlementAddlistRequest 初始化AlibabaMjPresaleSettlementAddlistAPIRequest对象
func NewAlibabaMjPresaleSettlementAddlistRequest() *AlibabaMjPresaleSettlementAddlistAPIRequest {
	return &AlibabaMjPresaleSettlementAddlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMjPresaleSettlementAddlistAPIRequest) GetApiMethodName() string {
	return "alibaba.mj.presale.settlement.addlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMjPresaleSettlementAddlistAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is PreSaleRefundJson Setter
// 订单json格式数据
func (r *AlibabaMjPresaleSettlementAddlistAPIRequest) SetPreSaleRefundJson(_preSaleRefundJson string) error {
	r._preSaleRefundJson = _preSaleRefundJson
	r.Set("pre_sale_refund_json", _preSaleRefundJson)
	return nil
}

// Get PreSaleRefundJson Getter
func (r AlibabaMjPresaleSettlementAddlistAPIRequest) GetPreSaleRefundJson() string {
	return r._preSaleRefundJson
}
