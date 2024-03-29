package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallSscSupplyplatformCapacityEditAPIRequest 容量编辑 API请求
// tmall.ssc.supplyplatform.capacity.edit
//
// 容量编辑
type TmallSscSupplyplatformCapacityEditAPIRequest struct {
	model.Params
	// 容量调整
	_paramServiceCapacityAdjustReqDto *ServiceCapacityAdjustReqDto
}

// NewTmallSscSupplyplatformCapacityEditRequest 初始化TmallSscSupplyplatformCapacityEditAPIRequest对象
func NewTmallSscSupplyplatformCapacityEditRequest() *TmallSscSupplyplatformCapacityEditAPIRequest {
	return &TmallSscSupplyplatformCapacityEditAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallSscSupplyplatformCapacityEditAPIRequest) Reset() {
	r._paramServiceCapacityAdjustReqDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallSscSupplyplatformCapacityEditAPIRequest) GetApiMethodName() string {
	return "tmall.ssc.supplyplatform.capacity.edit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallSscSupplyplatformCapacityEditAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallSscSupplyplatformCapacityEditAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamServiceCapacityAdjustReqDto is ParamServiceCapacityAdjustReqDto Setter
// 容量调整
func (r *TmallSscSupplyplatformCapacityEditAPIRequest) SetParamServiceCapacityAdjustReqDto(_paramServiceCapacityAdjustReqDto *ServiceCapacityAdjustReqDto) error {
	r._paramServiceCapacityAdjustReqDto = _paramServiceCapacityAdjustReqDto
	r.Set("param_service_capacity_adjust_req_dto", _paramServiceCapacityAdjustReqDto)
	return nil
}

// GetParamServiceCapacityAdjustReqDto ParamServiceCapacityAdjustReqDto Getter
func (r TmallSscSupplyplatformCapacityEditAPIRequest) GetParamServiceCapacityAdjustReqDto() *ServiceCapacityAdjustReqDto {
	return r._paramServiceCapacityAdjustReqDto
}

var poolTmallSscSupplyplatformCapacityEditAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallSscSupplyplatformCapacityEditRequest()
	},
}

// GetTmallSscSupplyplatformCapacityEditRequest 从 sync.Pool 获取 TmallSscSupplyplatformCapacityEditAPIRequest
func GetTmallSscSupplyplatformCapacityEditAPIRequest() *TmallSscSupplyplatformCapacityEditAPIRequest {
	return poolTmallSscSupplyplatformCapacityEditAPIRequest.Get().(*TmallSscSupplyplatformCapacityEditAPIRequest)
}

// ReleaseTmallSscSupplyplatformCapacityEditAPIRequest 将 TmallSscSupplyplatformCapacityEditAPIRequest 放入 sync.Pool
func ReleaseTmallSscSupplyplatformCapacityEditAPIRequest(v *TmallSscSupplyplatformCapacityEditAPIRequest) {
	v.Reset()
	poolTmallSscSupplyplatformCapacityEditAPIRequest.Put(v)
}
