package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJstMiniappCrowdUserAddAPIRequest
小程序添加用户到指定的活动 API请求
taobao.jst.miniapp.crowd.user.add

小程序添加用户到指定的活动 */
type TaobaoJstMiniappCrowdUserAddAPIRequest struct {
	model.Params
	// 活动code
	_crowdCode string
	// 小程序id
	_mcGwSourceMiniAppId string
	// 小程序appkey
	_mcGwSourceAppKey string
}

// NewTaobaoJstMiniappCrowdUserAddRequest 初始化TaobaoJstMiniappCrowdUserAddAPIRequest对象
func NewTaobaoJstMiniappCrowdUserAddRequest() *TaobaoJstMiniappCrowdUserAddAPIRequest {
	return &TaobaoJstMiniappCrowdUserAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstMiniappCrowdUserAddAPIRequest) GetApiMethodName() string {
	return "taobao.jst.miniapp.crowd.user.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstMiniappCrowdUserAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CrowdCode Setter
// 活动code
func (r *TaobaoJstMiniappCrowdUserAddAPIRequest) SetCrowdCode(_crowdCode string) error {
	r._crowdCode = _crowdCode
	r.Set("crowd_code", _crowdCode)
	return nil
}

// Get CrowdCode Getter
func (r TaobaoJstMiniappCrowdUserAddAPIRequest) GetCrowdCode() string {
	return r._crowdCode
}

// Set is McGwSourceMiniAppId Setter
// 小程序id
func (r *TaobaoJstMiniappCrowdUserAddAPIRequest) SetMcGwSourceMiniAppId(_mcGwSourceMiniAppId string) error {
	r._mcGwSourceMiniAppId = _mcGwSourceMiniAppId
	r.Set("mc_gw_source_mini_app_id", _mcGwSourceMiniAppId)
	return nil
}

// Get McGwSourceMiniAppId Getter
func (r TaobaoJstMiniappCrowdUserAddAPIRequest) GetMcGwSourceMiniAppId() string {
	return r._mcGwSourceMiniAppId
}

// Set is McGwSourceAppKey Setter
// 小程序appkey
func (r *TaobaoJstMiniappCrowdUserAddAPIRequest) SetMcGwSourceAppKey(_mcGwSourceAppKey string) error {
	r._mcGwSourceAppKey = _mcGwSourceAppKey
	r.Set("mc_gw_source_app_key", _mcGwSourceAppKey)
	return nil
}

// Get McGwSourceAppKey Getter
func (r TaobaoJstMiniappCrowdUserAddAPIRequest) GetMcGwSourceAppKey() string {
	return r._mcGwSourceAppKey
}
