package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterworkcardlogisticsinfocreateAPIRequest 创建服务履约物流单 API请求
// tmall.servicecenter.workcard.logisticsinfo.create
//
// 创建服务履约物流单
type TmallservicecenterworkcardlogisticsinfocreateAPIRequest struct {
	model.Params
	// 创建服务履约物流单信息
	_createLogisticsOrderRequest *CreateLogisticsOrderRequest
}

// NewTmallservicecenterworkcardlogisticsinfocreateRequest 初始化TmallservicecenterworkcardlogisticsinfocreateAPIRequest对象
func NewTmallservicecenterworkcardlogisticsinfocreateRequest() *TmallservicecenterworkcardlogisticsinfocreateAPIRequest {
	return &TmallservicecenterworkcardlogisticsinfocreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenterworkcardlogisticsinfocreateAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.logisticsinfo.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenterworkcardlogisticsinfocreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenterworkcardlogisticsinfocreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreateLogisticsOrderRequest is CreateLogisticsOrderRequest Setter
// 创建服务履约物流单信息
func (r *TmallservicecenterworkcardlogisticsinfocreateAPIRequest) SetCreateLogisticsOrderRequest(_createLogisticsOrderRequest *CreateLogisticsOrderRequest) error {
	r._createLogisticsOrderRequest = _createLogisticsOrderRequest
	r.Set("create_logistics_order_request", _createLogisticsOrderRequest)
	return nil
}

// GetCreateLogisticsOrderRequest CreateLogisticsOrderRequest Getter
func (r TmallservicecenterworkcardlogisticsinfocreateAPIRequest) GetCreateLogisticsOrderRequest() *CreateLogisticsOrderRequest {
	return r._createLogisticsOrderRequest
}
