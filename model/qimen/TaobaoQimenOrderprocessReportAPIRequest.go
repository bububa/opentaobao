package qimen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenOrderprocessReportAPIRequest 订单流水通知接口 API请求
// taobao.qimen.orderprocess.report
//
// taobao.qimen.orderprocess.report
type TaobaoQimenOrderprocessReportAPIRequest struct {
	model.Params
	//
	_request *OrderProcessReportRequest
}

// NewTaobaoQimenOrderprocessReportRequest 初始化TaobaoQimenOrderprocessReportAPIRequest对象
func NewTaobaoQimenOrderprocessReportRequest() *TaobaoQimenOrderprocessReportAPIRequest {
	return &TaobaoQimenOrderprocessReportAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQimenOrderprocessReportAPIRequest) Reset() {
	r._request = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenOrderprocessReportAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.orderprocess.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenOrderprocessReportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenOrderprocessReportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenOrderprocessReportAPIRequest) SetRequest(_request *OrderProcessReportRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenOrderprocessReportAPIRequest) GetRequest() *OrderProcessReportRequest {
	return r._request
}

var poolTaobaoQimenOrderprocessReportAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQimenOrderprocessReportRequest()
	},
}

// GetTaobaoQimenOrderprocessReportRequest 从 sync.Pool 获取 TaobaoQimenOrderprocessReportAPIRequest
func GetTaobaoQimenOrderprocessReportAPIRequest() *TaobaoQimenOrderprocessReportAPIRequest {
	return poolTaobaoQimenOrderprocessReportAPIRequest.Get().(*TaobaoQimenOrderprocessReportAPIRequest)
}

// ReleaseTaobaoQimenOrderprocessReportAPIRequest 将 TaobaoQimenOrderprocessReportAPIRequest 放入 sync.Pool
func ReleaseTaobaoQimenOrderprocessReportAPIRequest(v *TaobaoQimenOrderprocessReportAPIRequest) {
	v.Reset()
	poolTaobaoQimenOrderprocessReportAPIRequest.Put(v)
}
