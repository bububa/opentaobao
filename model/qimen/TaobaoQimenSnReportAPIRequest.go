package qimen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenSnReportAPIRequest 发货单SN通知接口 API请求
// taobao.qimen.sn.report
//
// WMS调用奇门的接口,在仓库出库单后, 把SN信息回传给ERP
type TaobaoQimenSnReportAPIRequest struct {
	model.Params
	//
	_request *SnReportRequest
}

// NewTaobaoQimenSnReportRequest 初始化TaobaoQimenSnReportAPIRequest对象
func NewTaobaoQimenSnReportRequest() *TaobaoQimenSnReportAPIRequest {
	return &TaobaoQimenSnReportAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQimenSnReportAPIRequest) Reset() {
	r._request = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenSnReportAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.sn.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenSnReportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenSnReportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenSnReportAPIRequest) SetRequest(_request *SnReportRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenSnReportAPIRequest) GetRequest() *SnReportRequest {
	return r._request
}

var poolTaobaoQimenSnReportAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQimenSnReportRequest()
	},
}

// GetTaobaoQimenSnReportRequest 从 sync.Pool 获取 TaobaoQimenSnReportAPIRequest
func GetTaobaoQimenSnReportAPIRequest() *TaobaoQimenSnReportAPIRequest {
	return poolTaobaoQimenSnReportAPIRequest.Get().(*TaobaoQimenSnReportAPIRequest)
}

// ReleaseTaobaoQimenSnReportAPIRequest 将 TaobaoQimenSnReportAPIRequest 放入 sync.Pool
func ReleaseTaobaoQimenSnReportAPIRequest(v *TaobaoQimenSnReportAPIRequest) {
	v.Reset()
	poolTaobaoQimenSnReportAPIRequest.Put(v)
}
