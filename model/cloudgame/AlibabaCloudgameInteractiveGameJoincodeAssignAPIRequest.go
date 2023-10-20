package cloudgame

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacloudgameinteractivegamejoincodeassignAPIRequest 分配joinCode API请求
// alibaba.cloudgame.interactive.game.joincode.assign
//
// 分配joinCode
type AlibabacloudgameinteractivegamejoincodeassignAPIRequest struct {
	model.Params
	// 请求入参
	_joinCodeAssignRequest *JoinCodeAssignRequest
}

// NewAlibabacloudgameinteractivegamejoincodeassignRequest 初始化AlibabacloudgameinteractivegamejoincodeassignAPIRequest对象
func NewAlibabacloudgameinteractivegamejoincodeassignRequest() *AlibabacloudgameinteractivegamejoincodeassignAPIRequest {
	return &AlibabacloudgameinteractivegamejoincodeassignAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacloudgameinteractivegamejoincodeassignAPIRequest) GetApiMethodName() string {
	return "alibaba.cloudgame.interactive.game.joincode.assign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacloudgameinteractivegamejoincodeassignAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacloudgameinteractivegamejoincodeassignAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetJoinCodeAssignRequest is JoinCodeAssignRequest Setter
// 请求入参
func (r *AlibabacloudgameinteractivegamejoincodeassignAPIRequest) SetJoinCodeAssignRequest(_joinCodeAssignRequest *JoinCodeAssignRequest) error {
	r._joinCodeAssignRequest = _joinCodeAssignRequest
	r.Set("join_code_assign_request", _joinCodeAssignRequest)
	return nil
}

// GetJoinCodeAssignRequest JoinCodeAssignRequest Getter
func (r AlibabacloudgameinteractivegamejoincodeassignAPIRequest) GetJoinCodeAssignRequest() *JoinCodeAssignRequest {
	return r._joinCodeAssignRequest
}
