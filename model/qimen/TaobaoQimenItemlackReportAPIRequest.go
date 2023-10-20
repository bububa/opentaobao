package qimen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenItemlackReportAPIRequest 发货单缺货通知接口 API请求
// taobao.qimen.itemlack.report
//
// WMS调用奇门的接口,将商家在库某商品缺货的信息回传给ERP
type TaobaoQimenItemlackReportAPIRequest struct {
	model.Params
	//
	_request *ItemLackReportRequest
}

// NewTaobaoQimenItemlackReportRequest 初始化TaobaoQimenItemlackReportAPIRequest对象
func NewTaobaoQimenItemlackReportRequest() *TaobaoQimenItemlackReportAPIRequest {
	return &TaobaoQimenItemlackReportAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQimenItemlackReportAPIRequest) Reset() {
	r._request = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenItemlackReportAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.itemlack.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenItemlackReportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenItemlackReportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenItemlackReportAPIRequest) SetRequest(_request *ItemLackReportRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenItemlackReportAPIRequest) GetRequest() *ItemLackReportRequest {
	return r._request
}

var poolTaobaoQimenItemlackReportAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQimenItemlackReportRequest()
	},
}

// GetTaobaoQimenItemlackReportRequest 从 sync.Pool 获取 TaobaoQimenItemlackReportAPIRequest
func GetTaobaoQimenItemlackReportAPIRequest() *TaobaoQimenItemlackReportAPIRequest {
	return poolTaobaoQimenItemlackReportAPIRequest.Get().(*TaobaoQimenItemlackReportAPIRequest)
}

// ReleaseTaobaoQimenItemlackReportAPIRequest 将 TaobaoQimenItemlackReportAPIRequest 放入 sync.Pool
func ReleaseTaobaoQimenItemlackReportAPIRequest(v *TaobaoQimenItemlackReportAPIRequest) {
	v.Reset()
	poolTaobaoQimenItemlackReportAPIRequest.Put(v)
}
