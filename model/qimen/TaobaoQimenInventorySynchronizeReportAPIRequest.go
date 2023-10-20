package qimen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenInventorySynchronizeReportAPIRequest 库存状态同步确认接口 API请求
// taobao.qimen.inventory.synchronize.report
//
// 库存状态同步确认接口
type TaobaoQimenInventorySynchronizeReportAPIRequest struct {
	model.Params
	//
	_request *TaobaoQimenInventorySynchronizeReportRequest
}

// NewTaobaoQimenInventorySynchronizeReportRequest 初始化TaobaoQimenInventorySynchronizeReportAPIRequest对象
func NewTaobaoQimenInventorySynchronizeReportRequest() *TaobaoQimenInventorySynchronizeReportAPIRequest {
	return &TaobaoQimenInventorySynchronizeReportAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQimenInventorySynchronizeReportAPIRequest) Reset() {
	r._request = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenInventorySynchronizeReportAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.inventory.synchronize.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenInventorySynchronizeReportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenInventorySynchronizeReportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenInventorySynchronizeReportAPIRequest) SetRequest(_request *TaobaoQimenInventorySynchronizeReportRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenInventorySynchronizeReportAPIRequest) GetRequest() *TaobaoQimenInventorySynchronizeReportRequest {
	return r._request
}

var poolTaobaoQimenInventorySynchronizeReportAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQimenInventorySynchronizeReportRequest()
	},
}

// GetTaobaoQimenInventorySynchronizeReportRequest 从 sync.Pool 获取 TaobaoQimenInventorySynchronizeReportAPIRequest
func GetTaobaoQimenInventorySynchronizeReportAPIRequest() *TaobaoQimenInventorySynchronizeReportAPIRequest {
	return poolTaobaoQimenInventorySynchronizeReportAPIRequest.Get().(*TaobaoQimenInventorySynchronizeReportAPIRequest)
}

// ReleaseTaobaoQimenInventorySynchronizeReportAPIRequest 将 TaobaoQimenInventorySynchronizeReportAPIRequest 放入 sync.Pool
func ReleaseTaobaoQimenInventorySynchronizeReportAPIRequest(v *TaobaoQimenInventorySynchronizeReportAPIRequest) {
	v.Reset()
	poolTaobaoQimenInventorySynchronizeReportAPIRequest.Put(v)
}
