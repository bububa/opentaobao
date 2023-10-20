package qimen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenWavenumReportAPIRequest 发货单波次通知接口 API请求
// taobao.qimen.wavenum.report
//
// WMS调用奇门的接口,该接口用来给ERP或者OMS回传波次号及对应的发货单号，以支持商家货票同行、波次内包裹的处理等需求
type TaobaoQimenWavenumReportAPIRequest struct {
	model.Params
	//
	_request *WaveNumReportRequest
}

// NewTaobaoQimenWavenumReportRequest 初始化TaobaoQimenWavenumReportAPIRequest对象
func NewTaobaoQimenWavenumReportRequest() *TaobaoQimenWavenumReportAPIRequest {
	return &TaobaoQimenWavenumReportAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQimenWavenumReportAPIRequest) Reset() {
	r._request = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenWavenumReportAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.wavenum.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenWavenumReportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenWavenumReportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenWavenumReportAPIRequest) SetRequest(_request *WaveNumReportRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenWavenumReportAPIRequest) GetRequest() *WaveNumReportRequest {
	return r._request
}

var poolTaobaoQimenWavenumReportAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQimenWavenumReportRequest()
	},
}

// GetTaobaoQimenWavenumReportRequest 从 sync.Pool 获取 TaobaoQimenWavenumReportAPIRequest
func GetTaobaoQimenWavenumReportAPIRequest() *TaobaoQimenWavenumReportAPIRequest {
	return poolTaobaoQimenWavenumReportAPIRequest.Get().(*TaobaoQimenWavenumReportAPIRequest)
}

// ReleaseTaobaoQimenWavenumReportAPIRequest 将 TaobaoQimenWavenumReportAPIRequest 放入 sync.Pool
func ReleaseTaobaoQimenWavenumReportAPIRequest(v *TaobaoQimenWavenumReportAPIRequest) {
	v.Reset()
	poolTaobaoQimenWavenumReportAPIRequest.Put(v)
}
