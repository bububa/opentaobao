package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowtttradeservicegetAPIRequest 获取网厅号卡垂直标信息 API请求
// taobao.wtt.trade.service.get
//
// 查询网厅订单信息
type TaobaowtttradeservicegetAPIRequest struct {
	model.Params
	// 订单ID
	_bizOrder int64
}

// NewTaobaowtttradeservicegetRequest 初始化TaobaowtttradeservicegetAPIRequest对象
func NewTaobaowtttradeservicegetRequest() *TaobaowtttradeservicegetAPIRequest {
	return &TaobaowtttradeservicegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowtttradeservicegetAPIRequest) GetApiMethodName() string {
	return "taobao.wtt.trade.service.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowtttradeservicegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowtttradeservicegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizOrder is BizOrder Setter
// 订单ID
func (r *TaobaowtttradeservicegetAPIRequest) SetBizOrder(_bizOrder int64) error {
	r._bizOrder = _bizOrder
	r.Set("biz_order", _bizOrder)
	return nil
}

// GetBizOrder BizOrder Getter
func (r TaobaowtttradeservicegetAPIRequest) GetBizOrder() int64 {
	return r._bizOrder
}
