package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidletradecarperformAPIRequest 二手车寄卖履约接口 API请求
// alibaba.idle.trade.car.perform
//
// 二手车寄卖履约接口
type AlibabaidletradecarperformAPIRequest struct {
	model.Params
	// 履约参数
	_carConsignmentParam *CarConsignmentParam
}

// NewAlibabaidletradecarperformRequest 初始化AlibabaidletradecarperformAPIRequest对象
func NewAlibabaidletradecarperformRequest() *AlibabaidletradecarperformAPIRequest {
	return &AlibabaidletradecarperformAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidletradecarperformAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.trade.car.perform"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidletradecarperformAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidletradecarperformAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCarConsignmentParam is CarConsignmentParam Setter
// 履约参数
func (r *AlibabaidletradecarperformAPIRequest) SetCarConsignmentParam(_carConsignmentParam *CarConsignmentParam) error {
	r._carConsignmentParam = _carConsignmentParam
	r.Set("car_consignment_param", _carConsignmentParam)
	return nil
}

// GetCarConsignmentParam CarConsignmentParam Getter
func (r AlibabaidletradecarperformAPIRequest) GetCarConsignmentParam() *CarConsignmentParam {
	return r._carConsignmentParam
}
