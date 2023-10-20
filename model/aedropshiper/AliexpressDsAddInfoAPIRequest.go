package aedropshiper

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressdsaddinfoAPIRequest 上报DS信息 API请求
// aliexpress.ds.add.info
//
// ISV用户上报下游DS信息
type AliexpressdsaddinfoAPIRequest struct {
	model.Params
	// Request object
	_param0 *DropShipperReq
}

// NewAliexpressdsaddinfoRequest 初始化AliexpressdsaddinfoAPIRequest对象
func NewAliexpressdsaddinfoRequest() *AliexpressdsaddinfoAPIRequest {
	return &AliexpressdsaddinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressdsaddinfoAPIRequest) GetApiMethodName() string {
	return "aliexpress.ds.add.info"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressdsaddinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressdsaddinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// Request object
func (r *AliexpressdsaddinfoAPIRequest) SetParam0(_param0 *DropShipperReq) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AliexpressdsaddinfoAPIRequest) GetParam0() *DropShipperReq {
	return r._param0
}
