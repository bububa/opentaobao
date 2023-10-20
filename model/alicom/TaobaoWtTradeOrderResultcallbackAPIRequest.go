package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowttradeorderresultcallbackAPIRequest 商家回调接口 API请求
// taobao.wt.trade.order.resultcallback
//
// 阿里通信定制服务，商家发货后进行调用该接口，用于自动发货并确认收货
type TaobaowttradeorderresultcallbackAPIRequest struct {
	model.Params
	// 系统自动生成
	_param0 *OrderResultDto
}

// NewTaobaowttradeorderresultcallbackRequest 初始化TaobaowttradeorderresultcallbackAPIRequest对象
func NewTaobaowttradeorderresultcallbackRequest() *TaobaowttradeorderresultcallbackAPIRequest {
	return &TaobaowttradeorderresultcallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowttradeorderresultcallbackAPIRequest) GetApiMethodName() string {
	return "taobao.wt.trade.order.resultcallback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowttradeorderresultcallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowttradeorderresultcallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 系统自动生成
func (r *TaobaowttradeorderresultcallbackAPIRequest) SetParam0(_param0 *OrderResultDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaowttradeorderresultcallbackAPIRequest) GetParam0() *OrderResultDto {
	return r._param0
}
