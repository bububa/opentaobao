package ascp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangScitemQueryAPIRequest 货品查询 API请求
// alibaba.dchain.aoxiang.scitem.query
//
// 货品查询
type AlibabaDchainAoxiangScitemQueryAPIRequest struct {
	model.Params
	// 货品查询入参
	_queryScitemRequest *QueryScItemRequest
}

// NewAlibabaDchainAoxiangScitemQueryRequest 初始化AlibabaDchainAoxiangScitemQueryAPIRequest对象
func NewAlibabaDchainAoxiangScitemQueryRequest() *AlibabaDchainAoxiangScitemQueryAPIRequest {
	return &AlibabaDchainAoxiangScitemQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDchainAoxiangScitemQueryAPIRequest) Reset() {
	r._queryScitemRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangScitemQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.scitem.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangScitemQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDchainAoxiangScitemQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryScitemRequest is QueryScitemRequest Setter
// 货品查询入参
func (r *AlibabaDchainAoxiangScitemQueryAPIRequest) SetQueryScitemRequest(_queryScitemRequest *QueryScItemRequest) error {
	r._queryScitemRequest = _queryScitemRequest
	r.Set("query_scitem_request", _queryScitemRequest)
	return nil
}

// GetQueryScitemRequest QueryScitemRequest Getter
func (r AlibabaDchainAoxiangScitemQueryAPIRequest) GetQueryScitemRequest() *QueryScItemRequest {
	return r._queryScitemRequest
}

var poolAlibabaDchainAoxiangScitemQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDchainAoxiangScitemQueryRequest()
	},
}

// GetAlibabaDchainAoxiangScitemQueryRequest 从 sync.Pool 获取 AlibabaDchainAoxiangScitemQueryAPIRequest
func GetAlibabaDchainAoxiangScitemQueryAPIRequest() *AlibabaDchainAoxiangScitemQueryAPIRequest {
	return poolAlibabaDchainAoxiangScitemQueryAPIRequest.Get().(*AlibabaDchainAoxiangScitemQueryAPIRequest)
}

// ReleaseAlibabaDchainAoxiangScitemQueryAPIRequest 将 AlibabaDchainAoxiangScitemQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaDchainAoxiangScitemQueryAPIRequest(v *AlibabaDchainAoxiangScitemQueryAPIRequest) {
	v.Reset()
	poolAlibabaDchainAoxiangScitemQueryAPIRequest.Put(v)
}
