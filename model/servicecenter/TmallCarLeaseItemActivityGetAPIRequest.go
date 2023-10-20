package servicecenter

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCarLeaseItemActivityGetAPIRequest 查询汽车租赁活动信息 API请求
// tmall.car.lease.item.activity.get
//
// 查询汽车租赁活动信息
type TmallCarLeaseItemActivityGetAPIRequest struct {
	model.Params
}

// NewTmallCarLeaseItemActivityGetRequest 初始化TmallCarLeaseItemActivityGetAPIRequest对象
func NewTmallCarLeaseItemActivityGetRequest() *TmallCarLeaseItemActivityGetAPIRequest {
	return &TmallCarLeaseItemActivityGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallCarLeaseItemActivityGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCarLeaseItemActivityGetAPIRequest) GetApiMethodName() string {
	return "tmall.car.lease.item.activity.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCarLeaseItemActivityGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallCarLeaseItemActivityGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTmallCarLeaseItemActivityGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallCarLeaseItemActivityGetRequest()
	},
}

// GetTmallCarLeaseItemActivityGetRequest 从 sync.Pool 获取 TmallCarLeaseItemActivityGetAPIRequest
func GetTmallCarLeaseItemActivityGetAPIRequest() *TmallCarLeaseItemActivityGetAPIRequest {
	return poolTmallCarLeaseItemActivityGetAPIRequest.Get().(*TmallCarLeaseItemActivityGetAPIRequest)
}

// ReleaseTmallCarLeaseItemActivityGetAPIRequest 将 TmallCarLeaseItemActivityGetAPIRequest 放入 sync.Pool
func ReleaseTmallCarLeaseItemActivityGetAPIRequest(v *TmallCarLeaseItemActivityGetAPIRequest) {
	v.Reset()
	poolTmallCarLeaseItemActivityGetAPIRequest.Put(v)
}
