package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthtracecodesellermilktracetosourceadddataAPIRequest 奶粉溯源-同步数据 API请求
// alibaba.alihealth.tracecodeseller.milk.trace.tosource.add.data
//
// 奶粉溯源-同步数据
type AlibabaalihealthtracecodesellermilktracetosourceadddataAPIRequest struct {
	model.Params
	// 奶粉品牌ID
	_entId string
	// 奶粉数据
	_jsonStr string
}

// NewAlibabaalihealthtracecodesellermilktracetosourceadddataRequest 初始化AlibabaalihealthtracecodesellermilktracetosourceadddataAPIRequest对象
func NewAlibabaalihealthtracecodesellermilktracetosourceadddataRequest() *AlibabaalihealthtracecodesellermilktracetosourceadddataAPIRequest {
	return &AlibabaalihealthtracecodesellermilktracetosourceadddataAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthtracecodesellermilktracetosourceadddataAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.tracecodeseller.milk.trace.tosource.add.data"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthtracecodesellermilktracetosourceadddataAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthtracecodesellermilktracetosourceadddataAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEntId is EntId Setter
// 奶粉品牌ID
func (r *AlibabaalihealthtracecodesellermilktracetosourceadddataAPIRequest) SetEntId(_entId string) error {
	r._entId = _entId
	r.Set("ent_id", _entId)
	return nil
}

// GetEntId EntId Getter
func (r AlibabaalihealthtracecodesellermilktracetosourceadddataAPIRequest) GetEntId() string {
	return r._entId
}

// SetJsonStr is JsonStr Setter
// 奶粉数据
func (r *AlibabaalihealthtracecodesellermilktracetosourceadddataAPIRequest) SetJsonStr(_jsonStr string) error {
	r._jsonStr = _jsonStr
	r.Set("json_str", _jsonStr)
	return nil
}

// GetJsonStr JsonStr Getter
func (r AlibabaalihealthtracecodesellermilktracetosourceadddataAPIRequest) GetJsonStr() string {
	return r._jsonStr
}
