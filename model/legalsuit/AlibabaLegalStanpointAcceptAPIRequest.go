package legalsuit

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalStanpointAcceptAPIRequest 采纳口径 API请求
// alibaba.legal.stanpoint.accept
//
// 采纳口径
type AlibabaLegalStanpointAcceptAPIRequest struct {
	model.Params
	// 采纳口径id字符串
	_standpointIds string
	// 业务实体ID
	_busiId string
	// 工号
	_userId string
	// 系统标识
	_inputSystemCode string
}

// NewAlibabaLegalStanpointAcceptRequest 初始化AlibabaLegalStanpointAcceptAPIRequest对象
func NewAlibabaLegalStanpointAcceptRequest() *AlibabaLegalStanpointAcceptAPIRequest {
	return &AlibabaLegalStanpointAcceptAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLegalStanpointAcceptAPIRequest) Reset() {
	r._standpointIds = ""
	r._busiId = ""
	r._userId = ""
	r._inputSystemCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalStanpointAcceptAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.stanpoint.accept"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalStanpointAcceptAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLegalStanpointAcceptAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStandpointIds is StandpointIds Setter
// 采纳口径id字符串
func (r *AlibabaLegalStanpointAcceptAPIRequest) SetStandpointIds(_standpointIds string) error {
	r._standpointIds = _standpointIds
	r.Set("standpoint_ids", _standpointIds)
	return nil
}

// GetStandpointIds StandpointIds Getter
func (r AlibabaLegalStanpointAcceptAPIRequest) GetStandpointIds() string {
	return r._standpointIds
}

// SetBusiId is BusiId Setter
// 业务实体ID
func (r *AlibabaLegalStanpointAcceptAPIRequest) SetBusiId(_busiId string) error {
	r._busiId = _busiId
	r.Set("busi_id", _busiId)
	return nil
}

// GetBusiId BusiId Getter
func (r AlibabaLegalStanpointAcceptAPIRequest) GetBusiId() string {
	return r._busiId
}

// SetUserId is UserId Setter
// 工号
func (r *AlibabaLegalStanpointAcceptAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaLegalStanpointAcceptAPIRequest) GetUserId() string {
	return r._userId
}

// SetInputSystemCode is InputSystemCode Setter
// 系统标识
func (r *AlibabaLegalStanpointAcceptAPIRequest) SetInputSystemCode(_inputSystemCode string) error {
	r._inputSystemCode = _inputSystemCode
	r.Set("input_system_code", _inputSystemCode)
	return nil
}

// GetInputSystemCode InputSystemCode Getter
func (r AlibabaLegalStanpointAcceptAPIRequest) GetInputSystemCode() string {
	return r._inputSystemCode
}

var poolAlibabaLegalStanpointAcceptAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLegalStanpointAcceptRequest()
	},
}

// GetAlibabaLegalStanpointAcceptRequest 从 sync.Pool 获取 AlibabaLegalStanpointAcceptAPIRequest
func GetAlibabaLegalStanpointAcceptAPIRequest() *AlibabaLegalStanpointAcceptAPIRequest {
	return poolAlibabaLegalStanpointAcceptAPIRequest.Get().(*AlibabaLegalStanpointAcceptAPIRequest)
}

// ReleaseAlibabaLegalStanpointAcceptAPIRequest 将 AlibabaLegalStanpointAcceptAPIRequest 放入 sync.Pool
func ReleaseAlibabaLegalStanpointAcceptAPIRequest(v *AlibabaLegalStanpointAcceptAPIRequest) {
	v.Reset()
	poolAlibabaLegalStanpointAcceptAPIRequest.Put(v)
}
