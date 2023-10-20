package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaxianyutenderorderperformAPIRequest 闲鱼暗拍订单履约 API请求
// alibaba.xianyu.tender.order.perform
//
// 闲鱼暗拍订单履约
type AlibabaxianyutenderorderperformAPIRequest struct {
	model.Params
	// 入参
	_param0 *TenderOrderSynDto
}

// NewAlibabaxianyutenderorderperformRequest 初始化AlibabaxianyutenderorderperformAPIRequest对象
func NewAlibabaxianyutenderorderperformRequest() *AlibabaxianyutenderorderperformAPIRequest {
	return &AlibabaxianyutenderorderperformAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaxianyutenderorderperformAPIRequest) GetApiMethodName() string {
	return "alibaba.xianyu.tender.order.perform"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaxianyutenderorderperformAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaxianyutenderorderperformAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 入参
func (r *AlibabaxianyutenderorderperformAPIRequest) SetParam0(_param0 *TenderOrderSynDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaxianyutenderorderperformAPIRequest) GetParam0() *TenderOrderSynDto {
	return r._param0
}
