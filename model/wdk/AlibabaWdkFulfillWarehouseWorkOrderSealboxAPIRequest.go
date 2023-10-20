package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkfulfillwarehouseworkordersealboxAPIRequest 仓封箱回告 API请求
// alibaba.wdk.fulfill.warehouse.work.order.sealbox
//
// 仓封箱回告箱与包裹的关系
type AlibabawdkfulfillwarehouseworkordersealboxAPIRequest struct {
	model.Params
	// 同城交付物箱
	_sameTownBox *SameTownBox
}

// NewAlibabawdkfulfillwarehouseworkordersealboxRequest 初始化AlibabawdkfulfillwarehouseworkordersealboxAPIRequest对象
func NewAlibabawdkfulfillwarehouseworkordersealboxRequest() *AlibabawdkfulfillwarehouseworkordersealboxAPIRequest {
	return &AlibabawdkfulfillwarehouseworkordersealboxAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkfulfillwarehouseworkordersealboxAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.fulfill.warehouse.work.order.sealbox"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkfulfillwarehouseworkordersealboxAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkfulfillwarehouseworkordersealboxAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSameTownBox is SameTownBox Setter
// 同城交付物箱
func (r *AlibabawdkfulfillwarehouseworkordersealboxAPIRequest) SetSameTownBox(_sameTownBox *SameTownBox) error {
	r._sameTownBox = _sameTownBox
	r.Set("same_town_box", _sameTownBox)
	return nil
}

// GetSameTownBox SameTownBox Getter
func (r AlibabawdkfulfillwarehouseworkordersealboxAPIRequest) GetSameTownBox() *SameTownBox {
	return r._sameTownBox
}
