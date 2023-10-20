package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomecommunitybrokersubmitAPIRequest 提交小区专家 API请求
// alibaba.alihouse.existinghome.community.broker.submit
//
// 提交小区专家
type AlibabaalihouseexistinghomecommunitybrokersubmitAPIRequest struct {
	model.Params
	// 小区专家关系
	_relationDto *SubmitCommunityAgentRelationDto
}

// NewAlibabaalihouseexistinghomecommunitybrokersubmitRequest 初始化AlibabaalihouseexistinghomecommunitybrokersubmitAPIRequest对象
func NewAlibabaalihouseexistinghomecommunitybrokersubmitRequest() *AlibabaalihouseexistinghomecommunitybrokersubmitAPIRequest {
	return &AlibabaalihouseexistinghomecommunitybrokersubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseexistinghomecommunitybrokersubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.community.broker.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseexistinghomecommunitybrokersubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseexistinghomecommunitybrokersubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRelationDto is RelationDto Setter
// 小区专家关系
func (r *AlibabaalihouseexistinghomecommunitybrokersubmitAPIRequest) SetRelationDto(_relationDto *SubmitCommunityAgentRelationDto) error {
	r._relationDto = _relationDto
	r.Set("relation_dto", _relationDto)
	return nil
}

// GetRelationDto RelationDto Getter
func (r AlibabaalihouseexistinghomecommunitybrokersubmitAPIRequest) GetRelationDto() *SubmitCommunityAgentRelationDto {
	return r._relationDto
}
