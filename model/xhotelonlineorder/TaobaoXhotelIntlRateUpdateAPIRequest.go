package xhotelonlineorder

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelIntlRateUpdateAPIRequest 不落库商家推送更新酒店rate API请求
// taobao.xhotel.intl.rate.update
//
// 商家主动推送不落库商品的酒店信息
type TaobaoXhotelIntlRateUpdateAPIRequest struct {
	model.Params
	// rate更新参数
	_updateRateParam *UpdateRateParam
}

// NewTaobaoXhotelIntlRateUpdateRequest 初始化TaobaoXhotelIntlRateUpdateAPIRequest对象
func NewTaobaoXhotelIntlRateUpdateRequest() *TaobaoXhotelIntlRateUpdateAPIRequest {
	return &TaobaoXhotelIntlRateUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelIntlRateUpdateAPIRequest) Reset() {
	r._updateRateParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelIntlRateUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.intl.rate.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelIntlRateUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelIntlRateUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUpdateRateParam is UpdateRateParam Setter
// rate更新参数
func (r *TaobaoXhotelIntlRateUpdateAPIRequest) SetUpdateRateParam(_updateRateParam *UpdateRateParam) error {
	r._updateRateParam = _updateRateParam
	r.Set("update_rate_param", _updateRateParam)
	return nil
}

// GetUpdateRateParam UpdateRateParam Getter
func (r TaobaoXhotelIntlRateUpdateAPIRequest) GetUpdateRateParam() *UpdateRateParam {
	return r._updateRateParam
}

var poolTaobaoXhotelIntlRateUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelIntlRateUpdateRequest()
	},
}

// GetTaobaoXhotelIntlRateUpdateRequest 从 sync.Pool 获取 TaobaoXhotelIntlRateUpdateAPIRequest
func GetTaobaoXhotelIntlRateUpdateAPIRequest() *TaobaoXhotelIntlRateUpdateAPIRequest {
	return poolTaobaoXhotelIntlRateUpdateAPIRequest.Get().(*TaobaoXhotelIntlRateUpdateAPIRequest)
}

// ReleaseTaobaoXhotelIntlRateUpdateAPIRequest 将 TaobaoXhotelIntlRateUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelIntlRateUpdateAPIRequest(v *TaobaoXhotelIntlRateUpdateAPIRequest) {
	v.Reset()
	poolTaobaoXhotelIntlRateUpdateAPIRequest.Put(v)
}
