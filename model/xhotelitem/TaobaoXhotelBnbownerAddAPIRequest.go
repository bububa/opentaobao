package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelbnbowneraddAPIRequest 民宿房东信息添加 API请求
// taobao.xhotel.bnbowner.add
//
// 添加和更新民宿房东的信息
type TaobaoxhotelbnbowneraddAPIRequest struct {
	model.Params
	// 添加房东信息的对象
	_addOwnerParam *AddOwnerParam
}

// NewTaobaoxhotelbnbowneraddRequest 初始化TaobaoxhotelbnbowneraddAPIRequest对象
func NewTaobaoxhotelbnbowneraddRequest() *TaobaoxhotelbnbowneraddAPIRequest {
	return &TaobaoxhotelbnbowneraddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelbnbowneraddAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.bnbowner.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelbnbowneraddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelbnbowneraddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAddOwnerParam is AddOwnerParam Setter
// 添加房东信息的对象
func (r *TaobaoxhotelbnbowneraddAPIRequest) SetAddOwnerParam(_addOwnerParam *AddOwnerParam) error {
	r._addOwnerParam = _addOwnerParam
	r.Set("add_owner_param", _addOwnerParam)
	return nil
}

// GetAddOwnerParam AddOwnerParam Getter
func (r TaobaoxhotelbnbowneraddAPIRequest) GetAddOwnerParam() *AddOwnerParam {
	return r._addOwnerParam
}
