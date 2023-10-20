package ascp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangScitemDeleteAPIRequest 货品删除 API请求
// alibaba.dchain.aoxiang.scitem.delete
//
// 货品删除
type AlibabaDchainAoxiangScitemDeleteAPIRequest struct {
	model.Params
	// 入参结构体
	_deleteScItemRequest *DeleteScItemRequest
}

// NewAlibabaDchainAoxiangScitemDeleteRequest 初始化AlibabaDchainAoxiangScitemDeleteAPIRequest对象
func NewAlibabaDchainAoxiangScitemDeleteRequest() *AlibabaDchainAoxiangScitemDeleteAPIRequest {
	return &AlibabaDchainAoxiangScitemDeleteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDchainAoxiangScitemDeleteAPIRequest) Reset() {
	r._deleteScItemRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangScitemDeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.scitem.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangScitemDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDchainAoxiangScitemDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeleteScItemRequest is DeleteScItemRequest Setter
// 入参结构体
func (r *AlibabaDchainAoxiangScitemDeleteAPIRequest) SetDeleteScItemRequest(_deleteScItemRequest *DeleteScItemRequest) error {
	r._deleteScItemRequest = _deleteScItemRequest
	r.Set("delete_sc_item_request", _deleteScItemRequest)
	return nil
}

// GetDeleteScItemRequest DeleteScItemRequest Getter
func (r AlibabaDchainAoxiangScitemDeleteAPIRequest) GetDeleteScItemRequest() *DeleteScItemRequest {
	return r._deleteScItemRequest
}

var poolAlibabaDchainAoxiangScitemDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDchainAoxiangScitemDeleteRequest()
	},
}

// GetAlibabaDchainAoxiangScitemDeleteRequest 从 sync.Pool 获取 AlibabaDchainAoxiangScitemDeleteAPIRequest
func GetAlibabaDchainAoxiangScitemDeleteAPIRequest() *AlibabaDchainAoxiangScitemDeleteAPIRequest {
	return poolAlibabaDchainAoxiangScitemDeleteAPIRequest.Get().(*AlibabaDchainAoxiangScitemDeleteAPIRequest)
}

// ReleaseAlibabaDchainAoxiangScitemDeleteAPIRequest 将 AlibabaDchainAoxiangScitemDeleteAPIRequest 放入 sync.Pool
func ReleaseAlibabaDchainAoxiangScitemDeleteAPIRequest(v *AlibabaDchainAoxiangScitemDeleteAPIRequest) {
	v.Reset()
	poolAlibabaDchainAoxiangScitemDeleteAPIRequest.Put(v)
}
