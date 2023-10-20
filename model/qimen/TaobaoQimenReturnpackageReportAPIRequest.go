package qimen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenReturnpackageReportAPIRequest 退货包裹状态通知接口 API请求
// taobao.qimen.returnpackage.report
//
// 退货包裹状态通知接口
type TaobaoQimenReturnpackageReportAPIRequest struct {
	model.Params
	//
	_request *TaobaoQimenReturnpackageReportRequest
}

// NewTaobaoQimenReturnpackageReportRequest 初始化TaobaoQimenReturnpackageReportAPIRequest对象
func NewTaobaoQimenReturnpackageReportRequest() *TaobaoQimenReturnpackageReportAPIRequest {
	return &TaobaoQimenReturnpackageReportAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQimenReturnpackageReportAPIRequest) Reset() {
	r._request = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenReturnpackageReportAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.returnpackage.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenReturnpackageReportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenReturnpackageReportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenReturnpackageReportAPIRequest) SetRequest(_request *TaobaoQimenReturnpackageReportRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenReturnpackageReportAPIRequest) GetRequest() *TaobaoQimenReturnpackageReportRequest {
	return r._request
}

var poolTaobaoQimenReturnpackageReportAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQimenReturnpackageReportRequest()
	},
}

// GetTaobaoQimenReturnpackageReportRequest 从 sync.Pool 获取 TaobaoQimenReturnpackageReportAPIRequest
func GetTaobaoQimenReturnpackageReportAPIRequest() *TaobaoQimenReturnpackageReportAPIRequest {
	return poolTaobaoQimenReturnpackageReportAPIRequest.Get().(*TaobaoQimenReturnpackageReportAPIRequest)
}

// ReleaseTaobaoQimenReturnpackageReportAPIRequest 将 TaobaoQimenReturnpackageReportAPIRequest 放入 sync.Pool
func ReleaseTaobaoQimenReturnpackageReportAPIRequest(v *TaobaoQimenReturnpackageReportAPIRequest) {
	v.Reset()
	poolTaobaoQimenReturnpackageReportAPIRequest.Put(v)
}
