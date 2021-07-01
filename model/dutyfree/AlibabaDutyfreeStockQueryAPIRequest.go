package dutyfree

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDutyfreeStockQueryAPIRequest
对外库存查询接口 API请求
alibaba.dutyfree.stock.query

对外部服务提供库存查询接口 */
type AlibabaDutyfreeStockQueryAPIRequest struct {
	model.Params
	// 条形码
	_barCode string
}

// NewAlibabaDutyfreeStockQueryRequest 初始化AlibabaDutyfreeStockQueryAPIRequest对象
func NewAlibabaDutyfreeStockQueryRequest() *AlibabaDutyfreeStockQueryAPIRequest {
	return &AlibabaDutyfreeStockQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDutyfreeStockQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.dutyfree.stock.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDutyfreeStockQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is BarCode Setter
// 条形码
func (r *AlibabaDutyfreeStockQueryAPIRequest) SetBarCode(_barCode string) error {
	r._barCode = _barCode
	r.Set("bar_code", _barCode)
	return nil
}

// Get BarCode Getter
func (r AlibabaDutyfreeStockQueryAPIRequest) GetBarCode() string {
	return r._barCode
}
