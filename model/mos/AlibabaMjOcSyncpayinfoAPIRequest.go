package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMjOcSyncpayinfoAPIRequest 支付参考号回传 API请求
// alibaba.mj.oc.syncpayinfo
//
// 支付参考号同步到oc
type AlibabaMjOcSyncpayinfoAPIRequest struct {
	model.Params
	// 支付参考号信息
	_posPay *PosPayDto
}

// NewAlibabaMjOcSyncpayinfoRequest 初始化AlibabaMjOcSyncpayinfoAPIRequest对象
func NewAlibabaMjOcSyncpayinfoRequest() *AlibabaMjOcSyncpayinfoAPIRequest {
	return &AlibabaMjOcSyncpayinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMjOcSyncpayinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.mj.oc.syncpayinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMjOcSyncpayinfoAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetPosPay is PosPay Setter
// 支付参考号信息
func (r *AlibabaMjOcSyncpayinfoAPIRequest) SetPosPay(_posPay *PosPayDto) error {
	r._posPay = _posPay
	r.Set("pos_pay", _posPay)
	return nil
}

// GetPosPay PosPay Getter
func (r AlibabaMjOcSyncpayinfoAPIRequest) GetPosPay() *PosPayDto {
	return r._posPay
}
