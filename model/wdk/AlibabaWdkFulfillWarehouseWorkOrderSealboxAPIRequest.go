package wdk

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.fulfill.warehouse.work.order.sealbox"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
