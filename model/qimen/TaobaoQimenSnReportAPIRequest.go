package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimensnreportAPIRequest 发货单SN通知接口 API请求
// taobao.qimen.sn.report
//
// WMS调用奇门的接口,在仓库出库单后, 把SN信息回传给ERP
type TaobaoqimensnreportAPIRequest struct {
	model.Params
	//
	_request *SnReportRequest
}

// NewTaobaoqimensnreportRequest 初始化TaobaoqimensnreportAPIRequest对象
func NewTaobaoqimensnreportRequest() *TaobaoqimensnreportAPIRequest {
	return &TaobaoqimensnreportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimensnreportAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.sn.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimensnreportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimensnreportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimensnreportAPIRequest) SetRequest(_request *SnReportRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimensnreportAPIRequest) GetRequest() *SnReportRequest {
	return r._request
}
