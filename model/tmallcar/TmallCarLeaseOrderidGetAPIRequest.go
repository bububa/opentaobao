package tmallcar

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCarLeaseOrderidGetAPIRequest 天猫开新车查询订单id API请求
// tmall.car.lease.orderid.get
//
// 天猫开新车查询订单id
type TmallCarLeaseOrderidGetAPIRequest struct {
	model.Params
	// openid
	_openId string
}

// NewTmallCarLeaseOrderidGetRequest 初始化TmallCarLeaseOrderidGetAPIRequest对象
func NewTmallCarLeaseOrderidGetRequest() *TmallCarLeaseOrderidGetAPIRequest {
	return &TmallCarLeaseOrderidGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallCarLeaseOrderidGetAPIRequest) Reset() {
	r._openId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCarLeaseOrderidGetAPIRequest) GetApiMethodName() string {
	return "tmall.car.lease.orderid.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCarLeaseOrderidGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallCarLeaseOrderidGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOpenId is OpenId Setter
// openid
func (r *TmallCarLeaseOrderidGetAPIRequest) SetOpenId(_openId string) error {
	r._openId = _openId
	r.Set("open_id", _openId)
	return nil
}

// GetOpenId OpenId Getter
func (r TmallCarLeaseOrderidGetAPIRequest) GetOpenId() string {
	return r._openId
}

var poolTmallCarLeaseOrderidGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallCarLeaseOrderidGetRequest()
	},
}

// GetTmallCarLeaseOrderidGetRequest 从 sync.Pool 获取 TmallCarLeaseOrderidGetAPIRequest
func GetTmallCarLeaseOrderidGetAPIRequest() *TmallCarLeaseOrderidGetAPIRequest {
	return poolTmallCarLeaseOrderidGetAPIRequest.Get().(*TmallCarLeaseOrderidGetAPIRequest)
}

// ReleaseTmallCarLeaseOrderidGetAPIRequest 将 TmallCarLeaseOrderidGetAPIRequest 放入 sync.Pool
func ReleaseTmallCarLeaseOrderidGetAPIRequest(v *TmallCarLeaseOrderidGetAPIRequest) {
	v.Reset()
	poolTmallCarLeaseOrderidGetAPIRequest.Put(v)
}
