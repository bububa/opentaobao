package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallItemHscodeDetailGetAPIRequest 通过hscode获取计量单位 API请求
// tmall.item.hscode.detail.get
//
// 通过hscode获取计量单位和销售单位
type TmallItemHscodeDetailGetAPIRequest struct {
	model.Params
	// hscode
	_hscode string
}

// NewTmallItemHscodeDetailGetRequest 初始化TmallItemHscodeDetailGetAPIRequest对象
func NewTmallItemHscodeDetailGetRequest() *TmallItemHscodeDetailGetAPIRequest {
	return &TmallItemHscodeDetailGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemHscodeDetailGetAPIRequest) GetApiMethodName() string {
	return "tmall.item.hscode.detail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemHscodeDetailGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallItemHscodeDetailGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHscode is Hscode Setter
// hscode
func (r *TmallItemHscodeDetailGetAPIRequest) SetHscode(_hscode string) error {
	r._hscode = _hscode
	r.Set("hscode", _hscode)
	return nil
}

// GetHscode Hscode Getter
func (r TmallItemHscodeDetailGetAPIRequest) GetHscode() string {
	return r._hscode
}
