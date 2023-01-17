package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWeikeEserviceSchedulePutAPIRequest 提交客服排班信息 API请求
// taobao.weike.eservice.schedule.put
//
// 添加、更新、删除排班信息
type TaobaoWeikeEserviceSchedulePutAPIRequest struct {
	model.Params
	// 按天排班信息
	_csSchedulings []CsSchedulingOneDayDto
	// 订单ID
	_orderId int64
}

// NewTaobaoWeikeEserviceSchedulePutRequest 初始化TaobaoWeikeEserviceSchedulePutAPIRequest对象
func NewTaobaoWeikeEserviceSchedulePutRequest() *TaobaoWeikeEserviceSchedulePutAPIRequest {
	return &TaobaoWeikeEserviceSchedulePutAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWeikeEserviceSchedulePutAPIRequest) GetApiMethodName() string {
	return "taobao.weike.eservice.schedule.put"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWeikeEserviceSchedulePutAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWeikeEserviceSchedulePutAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCsSchedulings is CsSchedulings Setter
// 按天排班信息
func (r *TaobaoWeikeEserviceSchedulePutAPIRequest) SetCsSchedulings(_csSchedulings []CsSchedulingOneDayDto) error {
	r._csSchedulings = _csSchedulings
	r.Set("cs_schedulings", _csSchedulings)
	return nil
}

// GetCsSchedulings CsSchedulings Getter
func (r TaobaoWeikeEserviceSchedulePutAPIRequest) GetCsSchedulings() []CsSchedulingOneDayDto {
	return r._csSchedulings
}

// SetOrderId is OrderId Setter
// 订单ID
func (r *TaobaoWeikeEserviceSchedulePutAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaoWeikeEserviceSchedulePutAPIRequest) GetOrderId() int64 {
	return r._orderId
}
