package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaStandpointHistorykeyGetAPIRequest 查询历史数据 API请求
// alibaba.standpoint.historykey.get
//
// 查询历史数据
type AlibabaStandpointHistorykeyGetAPIRequest struct {
	model.Params
	// 工号
	_userId string
	// 系统标识
	_inputSystemCode string
}

// NewAlibabaStandpointHistorykeyGetRequest 初始化AlibabaStandpointHistorykeyGetAPIRequest对象
func NewAlibabaStandpointHistorykeyGetRequest() *AlibabaStandpointHistorykeyGetAPIRequest {
	return &AlibabaStandpointHistorykeyGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaStandpointHistorykeyGetAPIRequest) GetApiMethodName() string {
	return "alibaba.standpoint.historykey.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaStandpointHistorykeyGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaStandpointHistorykeyGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserId is UserId Setter
// 工号
func (r *AlibabaStandpointHistorykeyGetAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaStandpointHistorykeyGetAPIRequest) GetUserId() string {
	return r._userId
}

// SetInputSystemCode is InputSystemCode Setter
// 系统标识
func (r *AlibabaStandpointHistorykeyGetAPIRequest) SetInputSystemCode(_inputSystemCode string) error {
	r._inputSystemCode = _inputSystemCode
	r.Set("input_system_code", _inputSystemCode)
	return nil
}

// GetInputSystemCode InputSystemCode Getter
func (r AlibabaStandpointHistorykeyGetAPIRequest) GetInputSystemCode() string {
	return r._inputSystemCode
}
