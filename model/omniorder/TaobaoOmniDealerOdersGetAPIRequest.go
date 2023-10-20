package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoomnidealerodersgetAPIRequest 获取单笔全渠道经销商订单的详细信息 API请求
// taobao.omni.dealer.oders.get
//
// 全渠道经销商获取单笔订单的详细信息
type TaobaoomnidealerodersgetAPIRequest struct {
	model.Params
	// 主订单ID
	_mainOrderId int64
}

// NewTaobaoomnidealerodersgetRequest 初始化TaobaoomnidealerodersgetAPIRequest对象
func NewTaobaoomnidealerodersgetRequest() *TaobaoomnidealerodersgetAPIRequest {
	return &TaobaoomnidealerodersgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoomnidealerodersgetAPIRequest) GetApiMethodName() string {
	return "taobao.omni.dealer.oders.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoomnidealerodersgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoomnidealerodersgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMainOrderId is MainOrderId Setter
// 主订单ID
func (r *TaobaoomnidealerodersgetAPIRequest) SetMainOrderId(_mainOrderId int64) error {
	r._mainOrderId = _mainOrderId
	r.Set("main_order_id", _mainOrderId)
	return nil
}

// GetMainOrderId MainOrderId Getter
func (r TaobaoomnidealerodersgetAPIRequest) GetMainOrderId() int64 {
	return r._mainOrderId
}
