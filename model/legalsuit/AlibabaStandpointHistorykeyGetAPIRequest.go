package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabastandpointhistorykeygetAPIRequest 查询历史数据 API请求
// alibaba.standpoint.historykey.get
//
// 查询历史数据
type AlibabastandpointhistorykeygetAPIRequest struct {
	model.Params
	// 工号
	_userId string
	// 系统标识
	_inputSystemCode string
}

// NewAlibabastandpointhistorykeygetRequest 初始化AlibabastandpointhistorykeygetAPIRequest对象
func NewAlibabastandpointhistorykeygetRequest() *AlibabastandpointhistorykeygetAPIRequest {
	return &AlibabastandpointhistorykeygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabastandpointhistorykeygetAPIRequest) GetApiMethodName() string {
	return "alibaba.standpoint.historykey.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabastandpointhistorykeygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabastandpointhistorykeygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserId is UserId Setter
// 工号
func (r *AlibabastandpointhistorykeygetAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabastandpointhistorykeygetAPIRequest) GetUserId() string {
	return r._userId
}

// SetInputSystemCode is InputSystemCode Setter
// 系统标识
func (r *AlibabastandpointhistorykeygetAPIRequest) SetInputSystemCode(_inputSystemCode string) error {
	r._inputSystemCode = _inputSystemCode
	r.Set("input_system_code", _inputSystemCode)
	return nil
}

// GetInputSystemCode InputSystemCode Getter
func (r AlibabastandpointhistorykeygetAPIRequest) GetInputSystemCode() string {
	return r._inputSystemCode
}
