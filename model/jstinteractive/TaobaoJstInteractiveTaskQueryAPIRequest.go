package jstinteractive

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojstinteractivetaskqueryAPIRequest 互动任务列表查询接口 API请求
// taobao.jst.interactive.task.query
//
// 查询用户的互动任务列表
type TaobaojstinteractivetaskqueryAPIRequest struct {
	model.Params
	// 小程序ID
	_mcGwSourceMiniAppId string
	// 小程序appkey
	_mcGwSourceAppKey string
}

// NewTaobaojstinteractivetaskqueryRequest 初始化TaobaojstinteractivetaskqueryAPIRequest对象
func NewTaobaojstinteractivetaskqueryRequest() *TaobaojstinteractivetaskqueryAPIRequest {
	return &TaobaojstinteractivetaskqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojstinteractivetaskqueryAPIRequest) GetApiMethodName() string {
	return "taobao.jst.interactive.task.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojstinteractivetaskqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojstinteractivetaskqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMcGwSourceMiniAppId is McGwSourceMiniAppId Setter
// 小程序ID
func (r *TaobaojstinteractivetaskqueryAPIRequest) SetMcGwSourceMiniAppId(_mcGwSourceMiniAppId string) error {
	r._mcGwSourceMiniAppId = _mcGwSourceMiniAppId
	r.Set("mc_gw_source_mini_app_id", _mcGwSourceMiniAppId)
	return nil
}

// GetMcGwSourceMiniAppId McGwSourceMiniAppId Getter
func (r TaobaojstinteractivetaskqueryAPIRequest) GetMcGwSourceMiniAppId() string {
	return r._mcGwSourceMiniAppId
}

// SetMcGwSourceAppKey is McGwSourceAppKey Setter
// 小程序appkey
func (r *TaobaojstinteractivetaskqueryAPIRequest) SetMcGwSourceAppKey(_mcGwSourceAppKey string) error {
	r._mcGwSourceAppKey = _mcGwSourceAppKey
	r.Set("mc_gw_source_app_key", _mcGwSourceAppKey)
	return nil
}

// GetMcGwSourceAppKey McGwSourceAppKey Getter
func (r TaobaojstinteractivetaskqueryAPIRequest) GetMcGwSourceAppKey() string {
	return r._mcGwSourceAppKey
}
