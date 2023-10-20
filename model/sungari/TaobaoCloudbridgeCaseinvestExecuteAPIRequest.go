package sungari

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCloudbridgeCaseinvestExecuteAPIRequest 红盾云桥案件协查服务 API请求
// taobao.cloudbridge.caseinvest.execute
//
// 通过API接口直接提供政府部门录入及查询函件服务
type TaobaoCloudbridgeCaseinvestExecuteAPIRequest struct {
	model.Params
	// 方法名称
	_apiName string
	// 方法参数
	_data string
}

// NewTaobaoCloudbridgeCaseinvestExecuteRequest 初始化TaobaoCloudbridgeCaseinvestExecuteAPIRequest对象
func NewTaobaoCloudbridgeCaseinvestExecuteRequest() *TaobaoCloudbridgeCaseinvestExecuteAPIRequest {
	return &TaobaoCloudbridgeCaseinvestExecuteAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoCloudbridgeCaseinvestExecuteAPIRequest) Reset() {
	r._apiName = ""
	r._data = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCloudbridgeCaseinvestExecuteAPIRequest) GetApiMethodName() string {
	return "taobao.cloudbridge.caseinvest.execute"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCloudbridgeCaseinvestExecuteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoCloudbridgeCaseinvestExecuteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApiName is ApiName Setter
// 方法名称
func (r *TaobaoCloudbridgeCaseinvestExecuteAPIRequest) SetApiName(_apiName string) error {
	r._apiName = _apiName
	r.Set("api_name", _apiName)
	return nil
}

// GetApiName ApiName Getter
func (r TaobaoCloudbridgeCaseinvestExecuteAPIRequest) GetApiName() string {
	return r._apiName
}

// SetData is Data Setter
// 方法参数
func (r *TaobaoCloudbridgeCaseinvestExecuteAPIRequest) SetData(_data string) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// GetData Data Getter
func (r TaobaoCloudbridgeCaseinvestExecuteAPIRequest) GetData() string {
	return r._data
}

var poolTaobaoCloudbridgeCaseinvestExecuteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoCloudbridgeCaseinvestExecuteRequest()
	},
}

// GetTaobaoCloudbridgeCaseinvestExecuteRequest 从 sync.Pool 获取 TaobaoCloudbridgeCaseinvestExecuteAPIRequest
func GetTaobaoCloudbridgeCaseinvestExecuteAPIRequest() *TaobaoCloudbridgeCaseinvestExecuteAPIRequest {
	return poolTaobaoCloudbridgeCaseinvestExecuteAPIRequest.Get().(*TaobaoCloudbridgeCaseinvestExecuteAPIRequest)
}

// ReleaseTaobaoCloudbridgeCaseinvestExecuteAPIRequest 将 TaobaoCloudbridgeCaseinvestExecuteAPIRequest 放入 sync.Pool
func ReleaseTaobaoCloudbridgeCaseinvestExecuteAPIRequest(v *TaobaoCloudbridgeCaseinvestExecuteAPIRequest) {
	v.Reset()
	poolTaobaoCloudbridgeCaseinvestExecuteAPIRequest.Put(v)
}
