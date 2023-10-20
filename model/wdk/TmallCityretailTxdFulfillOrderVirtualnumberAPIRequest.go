package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallCityretailTxdFulfillOrderVirtualnumberAPIRequest 淘鲜达虚拟号服务接口 API请求
// tmall.cityretail.txd.fulfill.order.virtualnumber
//
// 虚拟小号绑定接口，只有开通了虚拟号服务的淘鲜达商家能绑定。
type TmallCityretailTxdFulfillOrderVirtualnumberAPIRequest struct {
	model.Params
	// 淘鲜达交易单号
	_sourceOrderId string
	// 查询类型
	_type string
}

// NewTmallCityretailTxdFulfillOrderVirtualnumberRequest 初始化TmallCityretailTxdFulfillOrderVirtualnumberAPIRequest对象
func NewTmallCityretailTxdFulfillOrderVirtualnumberRequest() *TmallCityretailTxdFulfillOrderVirtualnumberAPIRequest {
	return &TmallCityretailTxdFulfillOrderVirtualnumberAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCityretailTxdFulfillOrderVirtualnumberAPIRequest) GetApiMethodName() string {
	return "tmall.cityretail.txd.fulfill.order.virtualnumber"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCityretailTxdFulfillOrderVirtualnumberAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallCityretailTxdFulfillOrderVirtualnumberAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSourceOrderId is SourceOrderId Setter
// 淘鲜达交易单号
func (r *TmallCityretailTxdFulfillOrderVirtualnumberAPIRequest) SetSourceOrderId(_sourceOrderId string) error {
	r._sourceOrderId = _sourceOrderId
	r.Set("source_order_id", _sourceOrderId)
	return nil
}

// GetSourceOrderId SourceOrderId Getter
func (r TmallCityretailTxdFulfillOrderVirtualnumberAPIRequest) GetSourceOrderId() string {
	return r._sourceOrderId
}

// SetType is Type Setter
// 查询类型
func (r *TmallCityretailTxdFulfillOrderVirtualnumberAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TmallCityretailTxdFulfillOrderVirtualnumberAPIRequest) GetType() string {
	return r._type
}
