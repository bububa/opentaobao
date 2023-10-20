package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenwavenumreportAPIRequest 发货单波次通知接口 API请求
// taobao.qimen.wavenum.report
//
// WMS调用奇门的接口,该接口用来给ERP或者OMS回传波次号及对应的发货单号，以支持商家货票同行、波次内包裹的处理等需求
type TaobaoqimenwavenumreportAPIRequest struct {
	model.Params
	//
	_request *WaveNumReportRequest
}

// NewTaobaoqimenwavenumreportRequest 初始化TaobaoqimenwavenumreportAPIRequest对象
func NewTaobaoqimenwavenumreportRequest() *TaobaoqimenwavenumreportAPIRequest {
	return &TaobaoqimenwavenumreportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimenwavenumreportAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.wavenum.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimenwavenumreportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimenwavenumreportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimenwavenumreportAPIRequest) SetRequest(_request *WaveNumReportRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimenwavenumreportAPIRequest) GetRequest() *WaveNumReportRequest {
	return r._request
}
