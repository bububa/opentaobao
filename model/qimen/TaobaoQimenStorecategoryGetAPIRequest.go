package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenstorecategorygetAPIRequest 门店类目获取接口 API请求
// taobao.qimen.storecategory.get
//
// 商家在ERP中调用该接口，获取门店类目
type TaobaoqimenstorecategorygetAPIRequest struct {
	model.Params
	// 备注
	_remark string
}

// NewTaobaoqimenstorecategorygetRequest 初始化TaobaoqimenstorecategorygetAPIRequest对象
func NewTaobaoqimenstorecategorygetRequest() *TaobaoqimenstorecategorygetAPIRequest {
	return &TaobaoqimenstorecategorygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimenstorecategorygetAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.storecategory.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimenstorecategorygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimenstorecategorygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRemark is Remark Setter
// 备注
func (r *TaobaoqimenstorecategorygetAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// GetRemark Remark Getter
func (r TaobaoqimenstorecategorygetAPIRequest) GetRemark() string {
	return r._remark
}
