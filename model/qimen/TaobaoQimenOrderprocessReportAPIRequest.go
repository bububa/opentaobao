package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenOrderprocessReportAPIRequest 订单流水通知接口 API请求
// taobao.qimen.orderprocess.report
//
// WMS调用奇门的接口,将订单在仓库的状态回传给ERP；场景说明：仓库仓内操作状态回传给ERP, 比如打包操作完成时, 回传一个打 包完成的状态给到ERP, ERP自行决定如何处理。
type TaobaoQimenOrderprocessReportAPIRequest struct {
	model.Params
	//
	_request *OrderProcessReportRequest
}

// NewTaobaoQimenOrderprocessReportRequest 初始化TaobaoQimenOrderprocessReportAPIRequest对象
func NewTaobaoQimenOrderprocessReportRequest() *TaobaoQimenOrderprocessReportAPIRequest {
	return &TaobaoQimenOrderprocessReportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenOrderprocessReportAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.orderprocess.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenOrderprocessReportAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Request Setter
//
func (r *TaobaoQimenOrderprocessReportAPIRequest) SetRequest(_request *OrderProcessReportRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// Get Request Getter
func (r TaobaoQimenOrderprocessReportAPIRequest) GetRequest() *OrderProcessReportRequest {
	return r._request
}
