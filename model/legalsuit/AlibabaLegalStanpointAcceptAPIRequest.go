package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegalstanpointacceptAPIRequest 采纳口径 API请求
// alibaba.legal.stanpoint.accept
//
// 采纳口径
type AlibabalegalstanpointacceptAPIRequest struct {
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

// NewAlibabalegalstanpointacceptRequest 初始化AlibabalegalstanpointacceptAPIRequest对象
func NewAlibabalegalstanpointacceptRequest() *AlibabalegalstanpointacceptAPIRequest {
	return &AlibabalegalstanpointacceptAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalegalstanpointacceptAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.stanpoint.accept"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalegalstanpointacceptAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalegalstanpointacceptAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStandpointIds is StandpointIds Setter
// 采纳口径id字符串
func (r *AlibabalegalstanpointacceptAPIRequest) SetStandpointIds(_standpointIds string) error {
	r._standpointIds = _standpointIds
	r.Set("standpoint_ids", _standpointIds)
	return nil
}

// GetStandpointIds StandpointIds Getter
func (r AlibabalegalstanpointacceptAPIRequest) GetStandpointIds() string {
	return r._standpointIds
}

// SetBusiId is BusiId Setter
// 业务实体ID
func (r *AlibabalegalstanpointacceptAPIRequest) SetBusiId(_busiId string) error {
	r._busiId = _busiId
	r.Set("busi_id", _busiId)
	return nil
}

// GetBusiId BusiId Getter
func (r AlibabalegalstanpointacceptAPIRequest) GetBusiId() string {
	return r._busiId
}

// SetUserId is UserId Setter
// 工号
func (r *AlibabalegalstanpointacceptAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabalegalstanpointacceptAPIRequest) GetUserId() string {
	return r._userId
}

// SetInputSystemCode is InputSystemCode Setter
// 系统标识
func (r *AlibabalegalstanpointacceptAPIRequest) SetInputSystemCode(_inputSystemCode string) error {
	r._inputSystemCode = _inputSystemCode
	r.Set("input_system_code", _inputSystemCode)
	return nil
}

// GetInputSystemCode InputSystemCode Getter
func (r AlibabalegalstanpointacceptAPIRequest) GetInputSystemCode() string {
	return r._inputSystemCode
}
