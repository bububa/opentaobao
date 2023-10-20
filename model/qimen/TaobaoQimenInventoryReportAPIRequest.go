package qimen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenInventoryReportAPIRequest 库存盘点通知接口 API请求
// taobao.qimen.inventory.report
//
// WMS调用奇门的接口,将库存盘点情况回传ERP
type TaobaoQimenInventoryReportAPIRequest struct {
	model.Params
	//
	_request *InventoryReportRequest
}

// NewTaobaoQimenInventoryReportRequest 初始化TaobaoQimenInventoryReportAPIRequest对象
func NewTaobaoQimenInventoryReportRequest() *TaobaoQimenInventoryReportAPIRequest {
	return &TaobaoQimenInventoryReportAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQimenInventoryReportAPIRequest) Reset() {
	r._request = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenInventoryReportAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.inventory.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenInventoryReportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenInventoryReportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenInventoryReportAPIRequest) SetRequest(_request *InventoryReportRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenInventoryReportAPIRequest) GetRequest() *InventoryReportRequest {
	return r._request
}

var poolTaobaoQimenInventoryReportAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQimenInventoryReportRequest()
	},
}

// GetTaobaoQimenInventoryReportRequest 从 sync.Pool 获取 TaobaoQimenInventoryReportAPIRequest
func GetTaobaoQimenInventoryReportAPIRequest() *TaobaoQimenInventoryReportAPIRequest {
	return poolTaobaoQimenInventoryReportAPIRequest.Get().(*TaobaoQimenInventoryReportAPIRequest)
}

// ReleaseTaobaoQimenInventoryReportAPIRequest 将 TaobaoQimenInventoryReportAPIRequest 放入 sync.Pool
func ReleaseTaobaoQimenInventoryReportAPIRequest(v *TaobaoQimenInventoryReportAPIRequest) {
	v.Reset()
	poolTaobaoQimenInventoryReportAPIRequest.Put(v)
}
