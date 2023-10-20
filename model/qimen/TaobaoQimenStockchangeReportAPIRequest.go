package qimen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenStockchangeReportAPIRequest 库存异动通知接口 API请求
// taobao.qimen.stockchange.report
//
// taobao.qimen.stockchange.report
type TaobaoQimenStockchangeReportAPIRequest struct {
	model.Params
	//
	_request *StockChangeReportRequest
}

// NewTaobaoQimenStockchangeReportRequest 初始化TaobaoQimenStockchangeReportAPIRequest对象
func NewTaobaoQimenStockchangeReportRequest() *TaobaoQimenStockchangeReportAPIRequest {
	return &TaobaoQimenStockchangeReportAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQimenStockchangeReportAPIRequest) Reset() {
	r._request = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenStockchangeReportAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.stockchange.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenStockchangeReportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenStockchangeReportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenStockchangeReportAPIRequest) SetRequest(_request *StockChangeReportRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenStockchangeReportAPIRequest) GetRequest() *StockChangeReportRequest {
	return r._request
}

var poolTaobaoQimenStockchangeReportAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQimenStockchangeReportRequest()
	},
}

// GetTaobaoQimenStockchangeReportRequest 从 sync.Pool 获取 TaobaoQimenStockchangeReportAPIRequest
func GetTaobaoQimenStockchangeReportAPIRequest() *TaobaoQimenStockchangeReportAPIRequest {
	return poolTaobaoQimenStockchangeReportAPIRequest.Get().(*TaobaoQimenStockchangeReportAPIRequest)
}

// ReleaseTaobaoQimenStockchangeReportAPIRequest 将 TaobaoQimenStockchangeReportAPIRequest 放入 sync.Pool
func ReleaseTaobaoQimenStockchangeReportAPIRequest(v *TaobaoQimenStockchangeReportAPIRequest) {
	v.Reset()
	poolTaobaoQimenStockchangeReportAPIRequest.Put(v)
}
