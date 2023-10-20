package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenstockchangereportAPIRequest 库存异动通知接口 API请求
// taobao.qimen.stockchange.report
//
// taobao.qimen.stockchange.report
type TaobaoqimenstockchangereportAPIRequest struct {
	model.Params
	//
	_request *StockChangeReportRequest
}

// NewTaobaoqimenstockchangereportRequest 初始化TaobaoqimenstockchangereportAPIRequest对象
func NewTaobaoqimenstockchangereportRequest() *TaobaoqimenstockchangereportAPIRequest {
	return &TaobaoqimenstockchangereportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimenstockchangereportAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.stockchange.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimenstockchangereportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimenstockchangereportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimenstockchangereportAPIRequest) SetRequest(_request *StockChangeReportRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimenstockchangereportAPIRequest) GetRequest() *StockChangeReportRequest {
	return r._request
}
