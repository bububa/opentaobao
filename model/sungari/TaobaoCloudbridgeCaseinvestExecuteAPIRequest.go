package sungari

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaocloudbridgecaseinvestexecuteAPIRequest 红盾云桥案件协查服务 API请求
// taobao.cloudbridge.caseinvest.execute
//
// 通过API接口直接提供政府部门录入及查询函件服务
type TaobaocloudbridgecaseinvestexecuteAPIRequest struct {
	model.Params
	// 方法名称
	_apiName string
	// 方法参数
	_data string
}

// NewTaobaocloudbridgecaseinvestexecuteRequest 初始化TaobaocloudbridgecaseinvestexecuteAPIRequest对象
func NewTaobaocloudbridgecaseinvestexecuteRequest() *TaobaocloudbridgecaseinvestexecuteAPIRequest {
	return &TaobaocloudbridgecaseinvestexecuteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaocloudbridgecaseinvestexecuteAPIRequest) GetApiMethodName() string {
	return "taobao.cloudbridge.caseinvest.execute"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaocloudbridgecaseinvestexecuteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaocloudbridgecaseinvestexecuteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApiName is ApiName Setter
// 方法名称
func (r *TaobaocloudbridgecaseinvestexecuteAPIRequest) SetApiName(_apiName string) error {
	r._apiName = _apiName
	r.Set("api_name", _apiName)
	return nil
}

// GetApiName ApiName Getter
func (r TaobaocloudbridgecaseinvestexecuteAPIRequest) GetApiName() string {
	return r._apiName
}

// SetData is Data Setter
// 方法参数
func (r *TaobaocloudbridgecaseinvestexecuteAPIRequest) SetData(_data string) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// GetData Data Getter
func (r TaobaocloudbridgecaseinvestexecuteAPIRequest) GetData() string {
	return r._data
}
