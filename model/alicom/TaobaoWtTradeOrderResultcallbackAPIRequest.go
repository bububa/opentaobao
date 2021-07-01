package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWtTradeOrderResultcallbackAPIRequest
商家回调接口 API请求
taobao.wt.trade.order.resultcallback

阿里通信定制服务，商家发货后进行调用该接口，用于自动发货并确认收货 */
type TaobaoWtTradeOrderResultcallbackAPIRequest struct {
	model.Params
	// 系统自动生成
	_param0 *OrderResultDto
}

// NewTaobaoWtTradeOrderResultcallbackRequest 初始化TaobaoWtTradeOrderResultcallbackAPIRequest对象
func NewTaobaoWtTradeOrderResultcallbackRequest() *TaobaoWtTradeOrderResultcallbackAPIRequest {
	return &TaobaoWtTradeOrderResultcallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWtTradeOrderResultcallbackAPIRequest) GetApiMethodName() string {
	return "taobao.wt.trade.order.resultcallback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWtTradeOrderResultcallbackAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param0 Setter
// 系统自动生成
func (r *TaobaoWtTradeOrderResultcallbackAPIRequest) SetParam0(_param0 *OrderResultDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// Get Param0 Getter
func (r TaobaoWtTradeOrderResultcallbackAPIRequest) GetParam0() *OrderResultDto {
	return r._param0
}
