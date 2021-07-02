package wlbimports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbImportsResourceGetAPIRequest 获取所有服务列表 API请求
// taobao.wlb.imports.resource.get
//
// 一般进口TOP接口，获取所有服务列表。
type TaobaoWlbImportsResourceGetAPIRequest struct {
	model.Params
	// 卖家发货地区域ID
	_fromId int64
	// 买家收货地信息
	_toAddress *ReciverAddressDo
}

// NewTaobaoWlbImportsResourceGetRequest 初始化TaobaoWlbImportsResourceGetAPIRequest对象
func NewTaobaoWlbImportsResourceGetRequest() *TaobaoWlbImportsResourceGetAPIRequest {
	return &TaobaoWlbImportsResourceGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbImportsResourceGetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.imports.resource.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbImportsResourceGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is FromId Setter
// 卖家发货地区域ID
func (r *TaobaoWlbImportsResourceGetAPIRequest) SetFromId(_fromId int64) error {
	r._fromId = _fromId
	r.Set("from_id", _fromId)
	return nil
}

// Get FromId Getter
func (r TaobaoWlbImportsResourceGetAPIRequest) GetFromId() int64 {
	return r._fromId
}

// Set is ToAddress Setter
// 买家收货地信息
func (r *TaobaoWlbImportsResourceGetAPIRequest) SetToAddress(_toAddress *ReciverAddressDo) error {
	r._toAddress = _toAddress
	r.Set("to_address", _toAddress)
	return nil
}

// Get ToAddress Getter
func (r TaobaoWlbImportsResourceGetAPIRequest) GetToAddress() *ReciverAddressDo {
	return r._toAddress
}
