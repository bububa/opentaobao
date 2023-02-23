package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenOrderSnReportAPIRequest 订单SN通知接口 API请求
// taobao.qimen.order.sn.report
//
// WMS调用奇门的接口,在出库、发货、入库等场景下，ERP和WMS之间同步操作的SN列表
type TaobaoQimenOrderSnReportAPIRequest struct {
	model.Params
	//
	_request *TaobaoQimenOrderSnReportRequest
}

// NewTaobaoQimenOrderSnReportRequest 初始化TaobaoQimenOrderSnReportAPIRequest对象
func NewTaobaoQimenOrderSnReportRequest() *TaobaoQimenOrderSnReportAPIRequest {
	return &TaobaoQimenOrderSnReportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenOrderSnReportAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.order.sn.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenOrderSnReportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenOrderSnReportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenOrderSnReportAPIRequest) SetRequest(_request *TaobaoQimenOrderSnReportRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenOrderSnReportAPIRequest) GetRequest() *TaobaoQimenOrderSnReportRequest {
	return r._request
}
