package iot

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopMessageAddtextAPIRequest 精灵代说 API请求
// taobao.ailab.aicloud.top.message.addtext
//
// 精灵代说
type TaobaoAilabAicloudTopMessageAddtextAPIRequest struct {
	model.Params
	// 设备id
	_param1 string
	// 代说文本
	_param2 string
	// 扩展信息，可以为空
	_param3 string
	// 用户信息
	_param0 *OpenBaseInfo
}

// NewTaobaoAilabAicloudTopMessageAddtextRequest 初始化TaobaoAilabAicloudTopMessageAddtextAPIRequest对象
func NewTaobaoAilabAicloudTopMessageAddtextRequest() *TaobaoAilabAicloudTopMessageAddtextAPIRequest {
	return &TaobaoAilabAicloudTopMessageAddtextAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAilabAicloudTopMessageAddtextAPIRequest) Reset() {
	r._param1 = ""
	r._param2 = ""
	r._param3 = ""
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopMessageAddtextAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.message.addtext"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopMessageAddtextAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAilabAicloudTopMessageAddtextAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam1 is Param1 Setter
// 设备id
func (r *TaobaoAilabAicloudTopMessageAddtextAPIRequest) SetParam1(_param1 string) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r TaobaoAilabAicloudTopMessageAddtextAPIRequest) GetParam1() string {
	return r._param1
}

// SetParam2 is Param2 Setter
// 代说文本
func (r *TaobaoAilabAicloudTopMessageAddtextAPIRequest) SetParam2(_param2 string) error {
	r._param2 = _param2
	r.Set("param2", _param2)
	return nil
}

// GetParam2 Param2 Getter
func (r TaobaoAilabAicloudTopMessageAddtextAPIRequest) GetParam2() string {
	return r._param2
}

// SetParam3 is Param3 Setter
// 扩展信息，可以为空
func (r *TaobaoAilabAicloudTopMessageAddtextAPIRequest) SetParam3(_param3 string) error {
	r._param3 = _param3
	r.Set("param3", _param3)
	return nil
}

// GetParam3 Param3 Getter
func (r TaobaoAilabAicloudTopMessageAddtextAPIRequest) GetParam3() string {
	return r._param3
}

// SetParam0 is Param0 Setter
// 用户信息
func (r *TaobaoAilabAicloudTopMessageAddtextAPIRequest) SetParam0(_param0 *OpenBaseInfo) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaoAilabAicloudTopMessageAddtextAPIRequest) GetParam0() *OpenBaseInfo {
	return r._param0
}

var poolTaobaoAilabAicloudTopMessageAddtextAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAilabAicloudTopMessageAddtextRequest()
	},
}

// GetTaobaoAilabAicloudTopMessageAddtextRequest 从 sync.Pool 获取 TaobaoAilabAicloudTopMessageAddtextAPIRequest
func GetTaobaoAilabAicloudTopMessageAddtextAPIRequest() *TaobaoAilabAicloudTopMessageAddtextAPIRequest {
	return poolTaobaoAilabAicloudTopMessageAddtextAPIRequest.Get().(*TaobaoAilabAicloudTopMessageAddtextAPIRequest)
}

// ReleaseTaobaoAilabAicloudTopMessageAddtextAPIRequest 将 TaobaoAilabAicloudTopMessageAddtextAPIRequest 放入 sync.Pool
func ReleaseTaobaoAilabAicloudTopMessageAddtextAPIRequest(v *TaobaoAilabAicloudTopMessageAddtextAPIRequest) {
	v.Reset()
	poolTaobaoAilabAicloudTopMessageAddtextAPIRequest.Put(v)
}
