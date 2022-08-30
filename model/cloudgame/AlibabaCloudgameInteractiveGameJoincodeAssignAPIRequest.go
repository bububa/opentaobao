package cloudgame

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCloudgameInteractiveGameJoincodeAssignAPIRequest 分配joinCode API请求
// alibaba.cloudgame.interactive.game.joincode.assign
//
// 分配joinCode
type AlibabaCloudgameInteractiveGameJoincodeAssignAPIRequest struct {
	model.Params
	// 请求入参
	_joinCodeAssignRequest *JoinCodeAssignRequest
}

// NewAlibabaCloudgameInteractiveGameJoincodeAssignRequest 初始化AlibabaCloudgameInteractiveGameJoincodeAssignAPIRequest对象
func NewAlibabaCloudgameInteractiveGameJoincodeAssignRequest() *AlibabaCloudgameInteractiveGameJoincodeAssignAPIRequest {
	return &AlibabaCloudgameInteractiveGameJoincodeAssignAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCloudgameInteractiveGameJoincodeAssignAPIRequest) GetApiMethodName() string {
	return "alibaba.cloudgame.interactive.game.joincode.assign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCloudgameInteractiveGameJoincodeAssignAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetJoinCodeAssignRequest is JoinCodeAssignRequest Setter
// 请求入参
func (r *AlibabaCloudgameInteractiveGameJoincodeAssignAPIRequest) SetJoinCodeAssignRequest(_joinCodeAssignRequest *JoinCodeAssignRequest) error {
	r._joinCodeAssignRequest = _joinCodeAssignRequest
	r.Set("join_code_assign_request", _joinCodeAssignRequest)
	return nil
}

// GetJoinCodeAssignRequest JoinCodeAssignRequest Getter
func (r AlibabaCloudgameInteractiveGameJoincodeAssignAPIRequest) GetJoinCodeAssignRequest() *JoinCodeAssignRequest {
	return r._joinCodeAssignRequest
}
