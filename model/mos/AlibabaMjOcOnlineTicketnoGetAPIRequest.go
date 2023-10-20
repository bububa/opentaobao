package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamjoconlineticketnogetAPIRequest 线上小票号获取 API请求
// alibaba.mj.oc.online.ticketno.get
//
// 线上小票号获取
type AlibabamjoconlineticketnogetAPIRequest struct {
	model.Params
	// 外部门店号
	_outStoreNo string
}

// NewAlibabamjoconlineticketnogetRequest 初始化AlibabamjoconlineticketnogetAPIRequest对象
func NewAlibabamjoconlineticketnogetRequest() *AlibabamjoconlineticketnogetAPIRequest {
	return &AlibabamjoconlineticketnogetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamjoconlineticketnogetAPIRequest) GetApiMethodName() string {
	return "alibaba.mj.oc.online.ticketno.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamjoconlineticketnogetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamjoconlineticketnogetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutStoreNo is OutStoreNo Setter
// 外部门店号
func (r *AlibabamjoconlineticketnogetAPIRequest) SetOutStoreNo(_outStoreNo string) error {
	r._outStoreNo = _outStoreNo
	r.Set("out_store_no", _outStoreNo)
	return nil
}

// GetOutStoreNo OutStoreNo Getter
func (r AlibabamjoconlineticketnogetAPIRequest) GetOutStoreNo() string {
	return r._outStoreNo
}
