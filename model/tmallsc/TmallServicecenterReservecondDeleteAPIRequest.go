package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterReservecondDeleteAPIRequest
删除主动预约开通条件 API请求
tmall.servicecenter.reservecond.delete

删除主动预约开通条件 */
type TmallServicecenterReservecondDeleteAPIRequest struct {
	model.Params
	// 主动预约条件删除
	_reserveOpenConditionDelDto *ReserveOpenConditionDelDto
}

// NewTmallServicecenterReservecondDeleteRequest 初始化TmallServicecenterReservecondDeleteAPIRequest对象
func NewTmallServicecenterReservecondDeleteRequest() *TmallServicecenterReservecondDeleteAPIRequest {
	return &TmallServicecenterReservecondDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterReservecondDeleteAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.reservecond.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterReservecondDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ReserveOpenConditionDelDto Setter
// 主动预约条件删除
func (r *TmallServicecenterReservecondDeleteAPIRequest) SetReserveOpenConditionDelDto(_reserveOpenConditionDelDto *ReserveOpenConditionDelDto) error {
	r._reserveOpenConditionDelDto = _reserveOpenConditionDelDto
	r.Set("reserve_open_condition_del_dto", _reserveOpenConditionDelDto)
	return nil
}

// Get ReserveOpenConditionDelDto Getter
func (r TmallServicecenterReservecondDeleteAPIRequest) GetReserveOpenConditionDelDto() *ReserveOpenConditionDelDto {
	return r._reserveOpenConditionDelDto
}
