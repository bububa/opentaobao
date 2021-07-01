package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenWavenumReportAPIRequest
发货单波次通知接口 API请求
taobao.qimen.wavenum.report

WMS调用奇门的接口,该接口用来给ERP或者OMS回传波次号及对应的发货单号，以支持商家货票同行、波次内包裹的处理等需求 */
type TaobaoQimenWavenumReportAPIRequest struct {
	model.Params
	//
	_request *WaveNumReportRequest
}

// NewTaobaoQimenWavenumReportRequest 初始化TaobaoQimenWavenumReportAPIRequest对象
func NewTaobaoQimenWavenumReportRequest() *TaobaoQimenWavenumReportAPIRequest {
	return &TaobaoQimenWavenumReportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenWavenumReportAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.wavenum.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenWavenumReportAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Request Setter
//
func (r *TaobaoQimenWavenumReportAPIRequest) SetRequest(_request *WaveNumReportRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// Get Request Getter
func (r TaobaoQimenWavenumReportAPIRequest) GetRequest() *WaveNumReportRequest {
	return r._request
}
