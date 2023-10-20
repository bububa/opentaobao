package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripagentflightintentionlistAPIRequest 意向单列表 API请求
// alitrip.agent.flight.intention.list
//
// 意向单列表
type AlitripagentflightintentionlistAPIRequest struct {
	model.Params
	// 入参对象
	_intentionListRequestDto *IntentionListRequestDto
}

// NewAlitripagentflightintentionlistRequest 初始化AlitripagentflightintentionlistAPIRequest对象
func NewAlitripagentflightintentionlistRequest() *AlitripagentflightintentionlistAPIRequest {
	return &AlitripagentflightintentionlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripagentflightintentionlistAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.flight.intention.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripagentflightintentionlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripagentflightintentionlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIntentionListRequestDto is IntentionListRequestDto Setter
// 入参对象
func (r *AlitripagentflightintentionlistAPIRequest) SetIntentionListRequestDto(_intentionListRequestDto *IntentionListRequestDto) error {
	r._intentionListRequestDto = _intentionListRequestDto
	r.Set("intention_list_request_dto", _intentionListRequestDto)
	return nil
}

// GetIntentionListRequestDto IntentionListRequestDto Getter
func (r AlitripagentflightintentionlistAPIRequest) GetIntentionListRequestDto() *IntentionListRequestDto {
	return r._intentionListRequestDto
}
