package flight

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripAgentCoordinateRejectAPIRequest 慧飞商家协同单拒绝接口 API请求
// alitrip.agent.coordinate.reject
//
// 慧飞商家协同单拒绝接口
type AlitripAgentCoordinateRejectAPIRequest struct {
	model.Params
	// 协同单拒绝参数
	_rejectDto *RejectDto
}

// NewAlitripAgentCoordinateRejectRequest 初始化AlitripAgentCoordinateRejectAPIRequest对象
func NewAlitripAgentCoordinateRejectRequest() *AlitripAgentCoordinateRejectAPIRequest {
	return &AlitripAgentCoordinateRejectAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripAgentCoordinateRejectAPIRequest) Reset() {
	r._rejectDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripAgentCoordinateRejectAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.coordinate.reject"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripAgentCoordinateRejectAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripAgentCoordinateRejectAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRejectDto is RejectDto Setter
// 协同单拒绝参数
func (r *AlitripAgentCoordinateRejectAPIRequest) SetRejectDto(_rejectDto *RejectDto) error {
	r._rejectDto = _rejectDto
	r.Set("reject_dto", _rejectDto)
	return nil
}

// GetRejectDto RejectDto Getter
func (r AlitripAgentCoordinateRejectAPIRequest) GetRejectDto() *RejectDto {
	return r._rejectDto
}

var poolAlitripAgentCoordinateRejectAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripAgentCoordinateRejectRequest()
	},
}

// GetAlitripAgentCoordinateRejectRequest 从 sync.Pool 获取 AlitripAgentCoordinateRejectAPIRequest
func GetAlitripAgentCoordinateRejectAPIRequest() *AlitripAgentCoordinateRejectAPIRequest {
	return poolAlitripAgentCoordinateRejectAPIRequest.Get().(*AlitripAgentCoordinateRejectAPIRequest)
}

// ReleaseAlitripAgentCoordinateRejectAPIRequest 将 AlitripAgentCoordinateRejectAPIRequest 放入 sync.Pool
func ReleaseAlitripAgentCoordinateRejectAPIRequest(v *AlitripAgentCoordinateRejectAPIRequest) {
	v.Reset()
	poolAlitripAgentCoordinateRejectAPIRequest.Put(v)
}
