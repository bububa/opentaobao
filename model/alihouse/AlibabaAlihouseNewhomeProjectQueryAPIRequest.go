package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeprojectqueryAPIRequest 查询楼盘相关信息 API请求
// alibaba.alihouse.newhome.project.query
//
// 根据outerid查询楼盘相关信息
type AlibabaalihousenewhomeprojectqueryAPIRequest struct {
	model.Params
	// 外部楼盘/小区id
	_outerId string
	// 商品id
	_itemId string
	// 门店ID
	_outerStoreId string
}

// NewAlibabaalihousenewhomeprojectqueryRequest 初始化AlibabaalihousenewhomeprojectqueryAPIRequest对象
func NewAlibabaalihousenewhomeprojectqueryRequest() *AlibabaalihousenewhomeprojectqueryAPIRequest {
	return &AlibabaalihousenewhomeprojectqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomeprojectqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomeprojectqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomeprojectqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterId is OuterId Setter
// 外部楼盘/小区id
func (r *AlibabaalihousenewhomeprojectqueryAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r AlibabaalihousenewhomeprojectqueryAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetItemId is ItemId Setter
// 商品id
func (r *AlibabaalihousenewhomeprojectqueryAPIRequest) SetItemId(_itemId string) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlibabaalihousenewhomeprojectqueryAPIRequest) GetItemId() string {
	return r._itemId
}

// SetOuterStoreId is OuterStoreId Setter
// 门店ID
func (r *AlibabaalihousenewhomeprojectqueryAPIRequest) SetOuterStoreId(_outerStoreId string) error {
	r._outerStoreId = _outerStoreId
	r.Set("outer_store_id", _outerStoreId)
	return nil
}

// GetOuterStoreId OuterStoreId Getter
func (r AlibabaalihousenewhomeprojectqueryAPIRequest) GetOuterStoreId() string {
	return r._outerStoreId
}
