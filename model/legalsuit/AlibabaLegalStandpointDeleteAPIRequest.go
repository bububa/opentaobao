package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegalstandpointdeleteAPIRequest 删除关联口径 API请求
// alibaba.legal.standpoint.delete
//
// 删除关联口径
type AlibabalegalstandpointdeleteAPIRequest struct {
	model.Params
	// 工号
	_userId string
	// 系统标识
	_inputSystemCode string
	// 关联id
	_refrenceId int64
}

// NewAlibabalegalstandpointdeleteRequest 初始化AlibabalegalstandpointdeleteAPIRequest对象
func NewAlibabalegalstandpointdeleteRequest() *AlibabalegalstandpointdeleteAPIRequest {
	return &AlibabalegalstandpointdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalegalstandpointdeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.standpoint.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalegalstandpointdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalegalstandpointdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserId is UserId Setter
// 工号
func (r *AlibabalegalstandpointdeleteAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabalegalstandpointdeleteAPIRequest) GetUserId() string {
	return r._userId
}

// SetInputSystemCode is InputSystemCode Setter
// 系统标识
func (r *AlibabalegalstandpointdeleteAPIRequest) SetInputSystemCode(_inputSystemCode string) error {
	r._inputSystemCode = _inputSystemCode
	r.Set("input_system_code", _inputSystemCode)
	return nil
}

// GetInputSystemCode InputSystemCode Getter
func (r AlibabalegalstandpointdeleteAPIRequest) GetInputSystemCode() string {
	return r._inputSystemCode
}

// SetRefrenceId is RefrenceId Setter
// 关联id
func (r *AlibabalegalstandpointdeleteAPIRequest) SetRefrenceId(_refrenceId int64) error {
	r._refrenceId = _refrenceId
	r.Set("refrence_id", _refrenceId)
	return nil
}

// GetRefrenceId RefrenceId Getter
func (r AlibabalegalstandpointdeleteAPIRequest) GetRefrenceId() int64 {
	return r._refrenceId
}
