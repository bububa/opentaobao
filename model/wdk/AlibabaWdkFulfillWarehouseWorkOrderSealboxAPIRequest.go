package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIRequest 仓封箱回告 API请求
// alibaba.wdk.fulfill.warehouse.work.order.sealbox
//
// 仓封箱回告箱与包裹的关系
type AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIRequest struct {
	model.Params
	// 同城交付物箱
	_sameTownBox *SameTownBox
}

// NewAlibabaWdkFulfillWarehouseWorkOrderSealboxRequest 初始化AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIRequest对象
func NewAlibabaWdkFulfillWarehouseWorkOrderSealboxRequest() *AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIRequest {
	return &AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIRequest) Reset() {
	r._sameTownBox = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.fulfill.warehouse.work.order.sealbox"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSameTownBox is SameTownBox Setter
// 同城交付物箱
func (r *AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIRequest) SetSameTownBox(_sameTownBox *SameTownBox) error {
	r._sameTownBox = _sameTownBox
	r.Set("same_town_box", _sameTownBox)
	return nil
}

// GetSameTownBox SameTownBox Getter
func (r AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIRequest) GetSameTownBox() *SameTownBox {
	return r._sameTownBox
}

var poolAlibabaWdkFulfillWarehouseWorkOrderSealboxAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkFulfillWarehouseWorkOrderSealboxRequest()
	},
}

// GetAlibabaWdkFulfillWarehouseWorkOrderSealboxRequest 从 sync.Pool 获取 AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIRequest
func GetAlibabaWdkFulfillWarehouseWorkOrderSealboxAPIRequest() *AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIRequest {
	return poolAlibabaWdkFulfillWarehouseWorkOrderSealboxAPIRequest.Get().(*AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIRequest)
}

// ReleaseAlibabaWdkFulfillWarehouseWorkOrderSealboxAPIRequest 将 AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkFulfillWarehouseWorkOrderSealboxAPIRequest(v *AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIRequest) {
	v.Reset()
	poolAlibabaWdkFulfillWarehouseWorkOrderSealboxAPIRequest.Put(v)
}
