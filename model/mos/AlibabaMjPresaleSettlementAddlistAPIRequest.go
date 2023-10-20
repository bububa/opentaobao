package mos

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMjPresaleSettlementAddlistAPIRequest 预售结算数据回传 API请求
// alibaba.mj.presale.settlement.addlist
//
// 用于预售活动结算数据的回传。
type AlibabaMjPresaleSettlementAddlistAPIRequest struct {
	model.Params
	// 订单json格式数据
	_preSaleRefundJson string
}

// NewAlibabaMjPresaleSettlementAddlistRequest 初始化AlibabaMjPresaleSettlementAddlistAPIRequest对象
func NewAlibabaMjPresaleSettlementAddlistRequest() *AlibabaMjPresaleSettlementAddlistAPIRequest {
	return &AlibabaMjPresaleSettlementAddlistAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMjPresaleSettlementAddlistAPIRequest) Reset() {
	r._preSaleRefundJson = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMjPresaleSettlementAddlistAPIRequest) GetApiMethodName() string {
	return "alibaba.mj.presale.settlement.addlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMjPresaleSettlementAddlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMjPresaleSettlementAddlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPreSaleRefundJson is PreSaleRefundJson Setter
// 订单json格式数据
func (r *AlibabaMjPresaleSettlementAddlistAPIRequest) SetPreSaleRefundJson(_preSaleRefundJson string) error {
	r._preSaleRefundJson = _preSaleRefundJson
	r.Set("pre_sale_refund_json", _preSaleRefundJson)
	return nil
}

// GetPreSaleRefundJson PreSaleRefundJson Getter
func (r AlibabaMjPresaleSettlementAddlistAPIRequest) GetPreSaleRefundJson() string {
	return r._preSaleRefundJson
}

var poolAlibabaMjPresaleSettlementAddlistAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMjPresaleSettlementAddlistRequest()
	},
}

// GetAlibabaMjPresaleSettlementAddlistRequest 从 sync.Pool 获取 AlibabaMjPresaleSettlementAddlistAPIRequest
func GetAlibabaMjPresaleSettlementAddlistAPIRequest() *AlibabaMjPresaleSettlementAddlistAPIRequest {
	return poolAlibabaMjPresaleSettlementAddlistAPIRequest.Get().(*AlibabaMjPresaleSettlementAddlistAPIRequest)
}

// ReleaseAlibabaMjPresaleSettlementAddlistAPIRequest 将 AlibabaMjPresaleSettlementAddlistAPIRequest 放入 sync.Pool
func ReleaseAlibabaMjPresaleSettlementAddlistAPIRequest(v *AlibabaMjPresaleSettlementAddlistAPIRequest) {
	v.Reset()
	poolAlibabaMjPresaleSettlementAddlistAPIRequest.Put(v)
}
