package xhotel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelBnbpromoUpdateAPIRequest 民宿营销活动更新 API请求
// taobao.xhotel.bnbpromo.update
//
// 全量更新对应外部活动code相关的营销活动信息
type TaobaoXhotelBnbpromoUpdateAPIRequest struct {
	model.Params
	// 更新营销活动的入参
	_updatePromoParam *UpdatePromoParam
}

// NewTaobaoXhotelBnbpromoUpdateRequest 初始化TaobaoXhotelBnbpromoUpdateAPIRequest对象
func NewTaobaoXhotelBnbpromoUpdateRequest() *TaobaoXhotelBnbpromoUpdateAPIRequest {
	return &TaobaoXhotelBnbpromoUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelBnbpromoUpdateAPIRequest) Reset() {
	r._updatePromoParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelBnbpromoUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.bnbpromo.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelBnbpromoUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelBnbpromoUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUpdatePromoParam is UpdatePromoParam Setter
// 更新营销活动的入参
func (r *TaobaoXhotelBnbpromoUpdateAPIRequest) SetUpdatePromoParam(_updatePromoParam *UpdatePromoParam) error {
	r._updatePromoParam = _updatePromoParam
	r.Set("update_promo_param", _updatePromoParam)
	return nil
}

// GetUpdatePromoParam UpdatePromoParam Getter
func (r TaobaoXhotelBnbpromoUpdateAPIRequest) GetUpdatePromoParam() *UpdatePromoParam {
	return r._updatePromoParam
}

var poolTaobaoXhotelBnbpromoUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelBnbpromoUpdateRequest()
	},
}

// GetTaobaoXhotelBnbpromoUpdateRequest 从 sync.Pool 获取 TaobaoXhotelBnbpromoUpdateAPIRequest
func GetTaobaoXhotelBnbpromoUpdateAPIRequest() *TaobaoXhotelBnbpromoUpdateAPIRequest {
	return poolTaobaoXhotelBnbpromoUpdateAPIRequest.Get().(*TaobaoXhotelBnbpromoUpdateAPIRequest)
}

// ReleaseTaobaoXhotelBnbpromoUpdateAPIRequest 将 TaobaoXhotelBnbpromoUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelBnbpromoUpdateAPIRequest(v *TaobaoXhotelBnbpromoUpdateAPIRequest) {
	v.Reset()
	poolTaobaoXhotelBnbpromoUpdateAPIRequest.Put(v)
}
