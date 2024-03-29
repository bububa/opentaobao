package ascp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTianmaoInventoryQueryAPIRequest 阿里巴巴.天猫.aic库存.查询 API请求
// alibaba.tianmao.inventory.query
//
// 阿里巴巴.天猫.aic库存.查询
type AlibabaTianmaoInventoryQueryAPIRequest struct {
	model.Params
	// 库存查询模型
	_hiErpQueryInventoryReq *HiErpQueryInventoryReq
}

// NewAlibabaTianmaoInventoryQueryRequest 初始化AlibabaTianmaoInventoryQueryAPIRequest对象
func NewAlibabaTianmaoInventoryQueryRequest() *AlibabaTianmaoInventoryQueryAPIRequest {
	return &AlibabaTianmaoInventoryQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTianmaoInventoryQueryAPIRequest) Reset() {
	r._hiErpQueryInventoryReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTianmaoInventoryQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.tianmao.inventory.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTianmaoInventoryQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTianmaoInventoryQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHiErpQueryInventoryReq is HiErpQueryInventoryReq Setter
// 库存查询模型
func (r *AlibabaTianmaoInventoryQueryAPIRequest) SetHiErpQueryInventoryReq(_hiErpQueryInventoryReq *HiErpQueryInventoryReq) error {
	r._hiErpQueryInventoryReq = _hiErpQueryInventoryReq
	r.Set("hi_erp_query_inventory_req", _hiErpQueryInventoryReq)
	return nil
}

// GetHiErpQueryInventoryReq HiErpQueryInventoryReq Getter
func (r AlibabaTianmaoInventoryQueryAPIRequest) GetHiErpQueryInventoryReq() *HiErpQueryInventoryReq {
	return r._hiErpQueryInventoryReq
}

var poolAlibabaTianmaoInventoryQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTianmaoInventoryQueryRequest()
	},
}

// GetAlibabaTianmaoInventoryQueryRequest 从 sync.Pool 获取 AlibabaTianmaoInventoryQueryAPIRequest
func GetAlibabaTianmaoInventoryQueryAPIRequest() *AlibabaTianmaoInventoryQueryAPIRequest {
	return poolAlibabaTianmaoInventoryQueryAPIRequest.Get().(*AlibabaTianmaoInventoryQueryAPIRequest)
}

// ReleaseAlibabaTianmaoInventoryQueryAPIRequest 将 AlibabaTianmaoInventoryQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaTianmaoInventoryQueryAPIRequest(v *AlibabaTianmaoInventoryQueryAPIRequest) {
	v.Reset()
	poolAlibabaTianmaoInventoryQueryAPIRequest.Put(v)
}
