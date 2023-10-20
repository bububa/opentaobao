package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallsscsupplyplatformcapacityeditAPIRequest 容量编辑 API请求
// tmall.ssc.supplyplatform.capacity.edit
//
// 容量编辑
type TmallsscsupplyplatformcapacityeditAPIRequest struct {
	model.Params
	// 容量调整
	_paramServiceCapacityAdjustReqDto *ServiceCapacityAdjustReqDto
}

// NewTmallsscsupplyplatformcapacityeditRequest 初始化TmallsscsupplyplatformcapacityeditAPIRequest对象
func NewTmallsscsupplyplatformcapacityeditRequest() *TmallsscsupplyplatformcapacityeditAPIRequest {
	return &TmallsscsupplyplatformcapacityeditAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallsscsupplyplatformcapacityeditAPIRequest) GetApiMethodName() string {
	return "tmall.ssc.supplyplatform.capacity.edit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallsscsupplyplatformcapacityeditAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallsscsupplyplatformcapacityeditAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamServiceCapacityAdjustReqDto is ParamServiceCapacityAdjustReqDto Setter
// 容量调整
func (r *TmallsscsupplyplatformcapacityeditAPIRequest) SetParamServiceCapacityAdjustReqDto(_paramServiceCapacityAdjustReqDto *ServiceCapacityAdjustReqDto) error {
	r._paramServiceCapacityAdjustReqDto = _paramServiceCapacityAdjustReqDto
	r.Set("param_service_capacity_adjust_req_dto", _paramServiceCapacityAdjustReqDto)
	return nil
}

// GetParamServiceCapacityAdjustReqDto ParamServiceCapacityAdjustReqDto Getter
func (r TmallsscsupplyplatformcapacityeditAPIRequest) GetParamServiceCapacityAdjustReqDto() *ServiceCapacityAdjustReqDto {
	return r._paramServiceCapacityAdjustReqDto
}
