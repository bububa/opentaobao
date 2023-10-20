package mos

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMjOcOnlineTicketnoGetAPIRequest 线上小票号获取 API请求
// alibaba.mj.oc.online.ticketno.get
//
// 线上小票号获取
type AlibabaMjOcOnlineTicketnoGetAPIRequest struct {
	model.Params
	// 外部门店号
	_outStoreNo string
}

// NewAlibabaMjOcOnlineTicketnoGetRequest 初始化AlibabaMjOcOnlineTicketnoGetAPIRequest对象
func NewAlibabaMjOcOnlineTicketnoGetRequest() *AlibabaMjOcOnlineTicketnoGetAPIRequest {
	return &AlibabaMjOcOnlineTicketnoGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMjOcOnlineTicketnoGetAPIRequest) Reset() {
	r._outStoreNo = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMjOcOnlineTicketnoGetAPIRequest) GetApiMethodName() string {
	return "alibaba.mj.oc.online.ticketno.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMjOcOnlineTicketnoGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMjOcOnlineTicketnoGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutStoreNo is OutStoreNo Setter
// 外部门店号
func (r *AlibabaMjOcOnlineTicketnoGetAPIRequest) SetOutStoreNo(_outStoreNo string) error {
	r._outStoreNo = _outStoreNo
	r.Set("out_store_no", _outStoreNo)
	return nil
}

// GetOutStoreNo OutStoreNo Getter
func (r AlibabaMjOcOnlineTicketnoGetAPIRequest) GetOutStoreNo() string {
	return r._outStoreNo
}

var poolAlibabaMjOcOnlineTicketnoGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMjOcOnlineTicketnoGetRequest()
	},
}

// GetAlibabaMjOcOnlineTicketnoGetRequest 从 sync.Pool 获取 AlibabaMjOcOnlineTicketnoGetAPIRequest
func GetAlibabaMjOcOnlineTicketnoGetAPIRequest() *AlibabaMjOcOnlineTicketnoGetAPIRequest {
	return poolAlibabaMjOcOnlineTicketnoGetAPIRequest.Get().(*AlibabaMjOcOnlineTicketnoGetAPIRequest)
}

// ReleaseAlibabaMjOcOnlineTicketnoGetAPIRequest 将 AlibabaMjOcOnlineTicketnoGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaMjOcOnlineTicketnoGetAPIRequest(v *AlibabaMjOcOnlineTicketnoGetAPIRequest) {
	v.Reset()
	poolAlibabaMjOcOnlineTicketnoGetAPIRequest.Put(v)
}
