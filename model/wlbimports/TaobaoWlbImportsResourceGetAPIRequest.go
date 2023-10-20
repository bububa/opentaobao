package wlbimports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlbimportsresourcegetAPIRequest 获取所有服务列表 API请求
// taobao.wlb.imports.resource.get
//
// 一般进口TOP接口，获取所有服务列表。
type TaobaowlbimportsresourcegetAPIRequest struct {
	model.Params
	// 卖家发货地区域ID
	_fromId int64
	// 买家收货地信息
	_toAddress *ReciverAddressDo
}

// NewTaobaowlbimportsresourcegetRequest 初始化TaobaowlbimportsresourcegetAPIRequest对象
func NewTaobaowlbimportsresourcegetRequest() *TaobaowlbimportsresourcegetAPIRequest {
	return &TaobaowlbimportsresourcegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowlbimportsresourcegetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.imports.resource.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowlbimportsresourcegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowlbimportsresourcegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFromId is FromId Setter
// 卖家发货地区域ID
func (r *TaobaowlbimportsresourcegetAPIRequest) SetFromId(_fromId int64) error {
	r._fromId = _fromId
	r.Set("from_id", _fromId)
	return nil
}

// GetFromId FromId Getter
func (r TaobaowlbimportsresourcegetAPIRequest) GetFromId() int64 {
	return r._fromId
}

// SetToAddress is ToAddress Setter
// 买家收货地信息
func (r *TaobaowlbimportsresourcegetAPIRequest) SetToAddress(_toAddress *ReciverAddressDo) error {
	r._toAddress = _toAddress
	r.Set("to_address", _toAddress)
	return nil
}

// GetToAddress ToAddress Getter
func (r TaobaowlbimportsresourcegetAPIRequest) GetToAddress() *ReciverAddressDo {
	return r._toAddress
}
