package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopFeedlistGetAPIRequest 获取对话流列表 API请求
// taobao.ailab.aicloud.top.feedlist.get
//
// 获取指定应用的对话流信息
type TaobaoAilabAicloudTopFeedlistGetAPIRequest struct {
	model.Params
	// 用户信息
	_param0 *OpenBaseInfo
	// 设备id
	_param1 string
	// 最后一条对话的key
	_param2 string
	// 单页的条目数，注意，是String类型！
	_param3 string
}

// NewTaobaoAilabAicloudTopFeedlistGetRequest 初始化TaobaoAilabAicloudTopFeedlistGetAPIRequest对象
func NewTaobaoAilabAicloudTopFeedlistGetRequest() *TaobaoAilabAicloudTopFeedlistGetAPIRequest {
	return &TaobaoAilabAicloudTopFeedlistGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopFeedlistGetAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.feedlist.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopFeedlistGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam0 is Param0 Setter
// 用户信息
func (r *TaobaoAilabAicloudTopFeedlistGetAPIRequest) SetParam0(_param0 *OpenBaseInfo) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaoAilabAicloudTopFeedlistGetAPIRequest) GetParam0() *OpenBaseInfo {
	return r._param0
}

// SetParam1 is Param1 Setter
// 设备id
func (r *TaobaoAilabAicloudTopFeedlistGetAPIRequest) SetParam1(_param1 string) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r TaobaoAilabAicloudTopFeedlistGetAPIRequest) GetParam1() string {
	return r._param1
}

// SetParam2 is Param2 Setter
// 最后一条对话的key
func (r *TaobaoAilabAicloudTopFeedlistGetAPIRequest) SetParam2(_param2 string) error {
	r._param2 = _param2
	r.Set("param2", _param2)
	return nil
}

// GetParam2 Param2 Getter
func (r TaobaoAilabAicloudTopFeedlistGetAPIRequest) GetParam2() string {
	return r._param2
}

// SetParam3 is Param3 Setter
// 单页的条目数，注意，是String类型！
func (r *TaobaoAilabAicloudTopFeedlistGetAPIRequest) SetParam3(_param3 string) error {
	r._param3 = _param3
	r.Set("param3", _param3)
	return nil
}

// GetParam3 Param3 Getter
func (r TaobaoAilabAicloudTopFeedlistGetAPIRequest) GetParam3() string {
	return r._param3
}
