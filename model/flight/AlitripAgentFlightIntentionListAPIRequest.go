package flight

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripAgentFlightIntentionListAPIRequest 意向单列表 API请求
// alitrip.agent.flight.intention.list
//
// 意向单列表
type AlitripAgentFlightIntentionListAPIRequest struct {
	model.Params
	// 入参对象
	_intentionListRequestDto *IntentionListRequestDto
}

// NewAlitripAgentFlightIntentionListRequest 初始化AlitripAgentFlightIntentionListAPIRequest对象
func NewAlitripAgentFlightIntentionListRequest() *AlitripAgentFlightIntentionListAPIRequest {
	return &AlitripAgentFlightIntentionListAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripAgentFlightIntentionListAPIRequest) Reset() {
	r._intentionListRequestDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripAgentFlightIntentionListAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.flight.intention.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripAgentFlightIntentionListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripAgentFlightIntentionListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIntentionListRequestDto is IntentionListRequestDto Setter
// 入参对象
func (r *AlitripAgentFlightIntentionListAPIRequest) SetIntentionListRequestDto(_intentionListRequestDto *IntentionListRequestDto) error {
	r._intentionListRequestDto = _intentionListRequestDto
	r.Set("intention_list_request_dto", _intentionListRequestDto)
	return nil
}

// GetIntentionListRequestDto IntentionListRequestDto Getter
func (r AlitripAgentFlightIntentionListAPIRequest) GetIntentionListRequestDto() *IntentionListRequestDto {
	return r._intentionListRequestDto
}

var poolAlitripAgentFlightIntentionListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripAgentFlightIntentionListRequest()
	},
}

// GetAlitripAgentFlightIntentionListRequest 从 sync.Pool 获取 AlitripAgentFlightIntentionListAPIRequest
func GetAlitripAgentFlightIntentionListAPIRequest() *AlitripAgentFlightIntentionListAPIRequest {
	return poolAlitripAgentFlightIntentionListAPIRequest.Get().(*AlitripAgentFlightIntentionListAPIRequest)
}

// ReleaseAlitripAgentFlightIntentionListAPIRequest 将 AlitripAgentFlightIntentionListAPIRequest 放入 sync.Pool
func ReleaseAlitripAgentFlightIntentionListAPIRequest(v *AlitripAgentFlightIntentionListAPIRequest) {
	v.Reset()
	poolAlitripAgentFlightIntentionListAPIRequest.Put(v)
}
