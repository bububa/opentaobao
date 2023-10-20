package jstinteractive

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstInteractiveTaskQueryAPIRequest 互动任务列表查询接口 API请求
// taobao.jst.interactive.task.query
//
// 查询用户的互动任务列表
type TaobaoJstInteractiveTaskQueryAPIRequest struct {
	model.Params
	// 小程序ID
	_mcGwSourceMiniAppId string
	// 小程序appkey
	_mcGwSourceAppKey string
}

// NewTaobaoJstInteractiveTaskQueryRequest 初始化TaobaoJstInteractiveTaskQueryAPIRequest对象
func NewTaobaoJstInteractiveTaskQueryRequest() *TaobaoJstInteractiveTaskQueryAPIRequest {
	return &TaobaoJstInteractiveTaskQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstInteractiveTaskQueryAPIRequest) GetApiMethodName() string {
	return "taobao.jst.interactive.task.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstInteractiveTaskQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJstInteractiveTaskQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMcGwSourceMiniAppId is McGwSourceMiniAppId Setter
// 小程序ID
func (r *TaobaoJstInteractiveTaskQueryAPIRequest) SetMcGwSourceMiniAppId(_mcGwSourceMiniAppId string) error {
	r._mcGwSourceMiniAppId = _mcGwSourceMiniAppId
	r.Set("mc_gw_source_mini_app_id", _mcGwSourceMiniAppId)
	return nil
}

// GetMcGwSourceMiniAppId McGwSourceMiniAppId Getter
func (r TaobaoJstInteractiveTaskQueryAPIRequest) GetMcGwSourceMiniAppId() string {
	return r._mcGwSourceMiniAppId
}

// SetMcGwSourceAppKey is McGwSourceAppKey Setter
// 小程序appkey
func (r *TaobaoJstInteractiveTaskQueryAPIRequest) SetMcGwSourceAppKey(_mcGwSourceAppKey string) error {
	r._mcGwSourceAppKey = _mcGwSourceAppKey
	r.Set("mc_gw_source_app_key", _mcGwSourceAppKey)
	return nil
}

// GetMcGwSourceAppKey McGwSourceAppKey Getter
func (r TaobaoJstInteractiveTaskQueryAPIRequest) GetMcGwSourceAppKey() string {
	return r._mcGwSourceAppKey
}
