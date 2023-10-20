package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterreserveconddeleteAPIRequest 删除主动预约开通条件 API请求
// tmall.servicecenter.reservecond.delete
//
// 删除主动预约开通条件
type TmallservicecenterreserveconddeleteAPIRequest struct {
	model.Params
	// 主动预约条件删除
	_reserveOpenConditionDelDto *ReserveOpenConditionDelDto
}

// NewTmallservicecenterreserveconddeleteRequest 初始化TmallservicecenterreserveconddeleteAPIRequest对象
func NewTmallservicecenterreserveconddeleteRequest() *TmallservicecenterreserveconddeleteAPIRequest {
	return &TmallservicecenterreserveconddeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenterreserveconddeleteAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.reservecond.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenterreserveconddeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenterreserveconddeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReserveOpenConditionDelDto is ReserveOpenConditionDelDto Setter
// 主动预约条件删除
func (r *TmallservicecenterreserveconddeleteAPIRequest) SetReserveOpenConditionDelDto(_reserveOpenConditionDelDto *ReserveOpenConditionDelDto) error {
	r._reserveOpenConditionDelDto = _reserveOpenConditionDelDto
	r.Set("reserve_open_condition_del_dto", _reserveOpenConditionDelDto)
	return nil
}

// GetReserveOpenConditionDelDto ReserveOpenConditionDelDto Getter
func (r TmallservicecenterreserveconddeleteAPIRequest) GetReserveOpenConditionDelDto() *ReserveOpenConditionDelDto {
	return r._reserveOpenConditionDelDto
}
