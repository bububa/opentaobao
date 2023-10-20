package dutyfree

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDutyfreeStockQueryAPIRequest 对外库存查询接口 API请求
// alibaba.dutyfree.stock.query
//
// 对外部服务提供库存查询接口
type AlibabaDutyfreeStockQueryAPIRequest struct {
	model.Params
	// 条形码
	_barCode string
}

// NewAlibabaDutyfreeStockQueryRequest 初始化AlibabaDutyfreeStockQueryAPIRequest对象
func NewAlibabaDutyfreeStockQueryRequest() *AlibabaDutyfreeStockQueryAPIRequest {
	return &AlibabaDutyfreeStockQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDutyfreeStockQueryAPIRequest) Reset() {
	r._barCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDutyfreeStockQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.dutyfree.stock.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDutyfreeStockQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDutyfreeStockQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBarCode is BarCode Setter
// 条形码
func (r *AlibabaDutyfreeStockQueryAPIRequest) SetBarCode(_barCode string) error {
	r._barCode = _barCode
	r.Set("bar_code", _barCode)
	return nil
}

// GetBarCode BarCode Getter
func (r AlibabaDutyfreeStockQueryAPIRequest) GetBarCode() string {
	return r._barCode
}

var poolAlibabaDutyfreeStockQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDutyfreeStockQueryRequest()
	},
}

// GetAlibabaDutyfreeStockQueryRequest 从 sync.Pool 获取 AlibabaDutyfreeStockQueryAPIRequest
func GetAlibabaDutyfreeStockQueryAPIRequest() *AlibabaDutyfreeStockQueryAPIRequest {
	return poolAlibabaDutyfreeStockQueryAPIRequest.Get().(*AlibabaDutyfreeStockQueryAPIRequest)
}

// ReleaseAlibabaDutyfreeStockQueryAPIRequest 将 AlibabaDutyfreeStockQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaDutyfreeStockQueryAPIRequest(v *AlibabaDutyfreeStockQueryAPIRequest) {
	v.Reset()
	poolAlibabaDutyfreeStockQueryAPIRequest.Put(v)
}
