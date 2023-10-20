package wlbimports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlbimportsordercancelAPIRequest 一般进口取消物流订单 API请求
// taobao.wlb.imports.order.cancel
//
// 商家在发货后，需要对订单进行取消，如果仓库已经收货则无法取消。
type TaobaowlbimportsordercancelAPIRequest struct {
	model.Params
	// 物流订单编号
	_lgorderCode string
}

// NewTaobaowlbimportsordercancelRequest 初始化TaobaowlbimportsordercancelAPIRequest对象
func NewTaobaowlbimportsordercancelRequest() *TaobaowlbimportsordercancelAPIRequest {
	return &TaobaowlbimportsordercancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowlbimportsordercancelAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.imports.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowlbimportsordercancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowlbimportsordercancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLgorderCode is LgorderCode Setter
// 物流订单编号
func (r *TaobaowlbimportsordercancelAPIRequest) SetLgorderCode(_lgorderCode string) error {
	r._lgorderCode = _lgorderCode
	r.Set("lgorder_code", _lgorderCode)
	return nil
}

// GetLgorderCode LgorderCode Getter
func (r TaobaowlbimportsordercancelAPIRequest) GetLgorderCode() string {
	return r._lgorderCode
}
