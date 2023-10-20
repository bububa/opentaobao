package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallitemhscodedetailgetAPIRequest 通过hscode获取计量单位 API请求
// tmall.item.hscode.detail.get
//
// 通过hscode获取计量单位和销售单位
type TmallitemhscodedetailgetAPIRequest struct {
	model.Params
	// hscode
	_hscode string
}

// NewTmallitemhscodedetailgetRequest 初始化TmallitemhscodedetailgetAPIRequest对象
func NewTmallitemhscodedetailgetRequest() *TmallitemhscodedetailgetAPIRequest {
	return &TmallitemhscodedetailgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallitemhscodedetailgetAPIRequest) GetApiMethodName() string {
	return "tmall.item.hscode.detail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallitemhscodedetailgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallitemhscodedetailgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHscode is Hscode Setter
// hscode
func (r *TmallitemhscodedetailgetAPIRequest) SetHscode(_hscode string) error {
	r._hscode = _hscode
	r.Set("hscode", _hscode)
	return nil
}

// GetHscode Hscode Getter
func (r TmallitemhscodedetailgetAPIRequest) GetHscode() string {
	return r._hscode
}
