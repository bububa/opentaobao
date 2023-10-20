package cloudgame

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCloudgameInteractiveGameJoincodeAssignAPIRequest) Reset() {
	r._joinCodeAssignRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCloudgameInteractiveGameJoincodeAssignAPIRequest) GetApiMethodName() string {
	return "alibaba.cloudgame.interactive.game.joincode.assign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCloudgameInteractiveGameJoincodeAssignAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCloudgameInteractiveGameJoincodeAssignAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaCloudgameInteractiveGameJoincodeAssignAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCloudgameInteractiveGameJoincodeAssignRequest()
	},
}

// GetAlibabaCloudgameInteractiveGameJoincodeAssignRequest 从 sync.Pool 获取 AlibabaCloudgameInteractiveGameJoincodeAssignAPIRequest
func GetAlibabaCloudgameInteractiveGameJoincodeAssignAPIRequest() *AlibabaCloudgameInteractiveGameJoincodeAssignAPIRequest {
	return poolAlibabaCloudgameInteractiveGameJoincodeAssignAPIRequest.Get().(*AlibabaCloudgameInteractiveGameJoincodeAssignAPIRequest)
}

// ReleaseAlibabaCloudgameInteractiveGameJoincodeAssignAPIRequest 将 AlibabaCloudgameInteractiveGameJoincodeAssignAPIRequest 放入 sync.Pool
func ReleaseAlibabaCloudgameInteractiveGameJoincodeAssignAPIRequest(v *AlibabaCloudgameInteractiveGameJoincodeAssignAPIRequest) {
	v.Reset()
	poolAlibabaCloudgameInteractiveGameJoincodeAssignAPIRequest.Put(v)
}
