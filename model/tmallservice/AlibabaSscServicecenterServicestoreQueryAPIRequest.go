package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSscServicecenterServicestoreQueryAPIRequest 根据天猫id查询门店信息 API请求
// alibaba.ssc.servicecenter.servicestore.query
//
// 根据天猫id查询门店信息
type AlibabaSscServicecenterServicestoreQueryAPIRequest struct {
	model.Params
	// 天猫id
	_id int64
}

// NewAlibabaSscServicecenterServicestoreQueryRequest 初始化AlibabaSscServicecenterServicestoreQueryAPIRequest对象
func NewAlibabaSscServicecenterServicestoreQueryRequest() *AlibabaSscServicecenterServicestoreQueryAPIRequest {
	return &AlibabaSscServicecenterServicestoreQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSscServicecenterServicestoreQueryAPIRequest) Reset() {
	r._id = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSscServicecenterServicestoreQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.ssc.servicecenter.servicestore.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSscServicecenterServicestoreQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSscServicecenterServicestoreQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// 天猫id
func (r *AlibabaSscServicecenterServicestoreQueryAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r AlibabaSscServicecenterServicestoreQueryAPIRequest) GetId() int64 {
	return r._id
}

var poolAlibabaSscServicecenterServicestoreQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSscServicecenterServicestoreQueryRequest()
	},
}

// GetAlibabaSscServicecenterServicestoreQueryRequest 从 sync.Pool 获取 AlibabaSscServicecenterServicestoreQueryAPIRequest
func GetAlibabaSscServicecenterServicestoreQueryAPIRequest() *AlibabaSscServicecenterServicestoreQueryAPIRequest {
	return poolAlibabaSscServicecenterServicestoreQueryAPIRequest.Get().(*AlibabaSscServicecenterServicestoreQueryAPIRequest)
}

// ReleaseAlibabaSscServicecenterServicestoreQueryAPIRequest 将 AlibabaSscServicecenterServicestoreQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaSscServicecenterServicestoreQueryAPIRequest(v *AlibabaSscServicecenterServicestoreQueryAPIRequest) {
	v.Reset()
	poolAlibabaSscServicecenterServicestoreQueryAPIRequest.Put(v)
}
