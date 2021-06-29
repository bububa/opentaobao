package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
小程序添加用户到指定的活动 API请求
taobao.jst.miniapp.crowd.user.add

小程序添加用户到指定的活动
*/
type TaobaoJstMiniappCrowdUserAddRequest struct {
    model.Params
    // 活动code
    crowdCode   string
    // 小程序id
    mcGwSourceMiniAppId   string
    // 小程序appkey
    mcGwSourceAppKey   string
}

// 初始化TaobaoJstMiniappCrowdUserAddRequest对象
func NewTaobaoJstMiniappCrowdUserAddRequest() *TaobaoJstMiniappCrowdUserAddRequest{
    return &TaobaoJstMiniappCrowdUserAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJstMiniappCrowdUserAddRequest) GetApiMethodName() string {
    return "taobao.jst.miniapp.crowd.user.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJstMiniappCrowdUserAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CrowdCode Setter
// 活动code
func (r *TaobaoJstMiniappCrowdUserAddRequest) SetCrowdCode(crowdCode string) error {
    r.crowdCode = crowdCode
    r.Set("crowd_code", crowdCode)
    return nil
}

// CrowdCode Getter
func (r TaobaoJstMiniappCrowdUserAddRequest) GetCrowdCode() string {
    return r.crowdCode
}
// McGwSourceMiniAppId Setter
// 小程序id
func (r *TaobaoJstMiniappCrowdUserAddRequest) SetMcGwSourceMiniAppId(mcGwSourceMiniAppId string) error {
    r.mcGwSourceMiniAppId = mcGwSourceMiniAppId
    r.Set("mc_gw_source_mini_app_id", mcGwSourceMiniAppId)
    return nil
}

// McGwSourceMiniAppId Getter
func (r TaobaoJstMiniappCrowdUserAddRequest) GetMcGwSourceMiniAppId() string {
    return r.mcGwSourceMiniAppId
}
// McGwSourceAppKey Setter
// 小程序appkey
func (r *TaobaoJstMiniappCrowdUserAddRequest) SetMcGwSourceAppKey(mcGwSourceAppKey string) error {
    r.mcGwSourceAppKey = mcGwSourceAppKey
    r.Set("mc_gw_source_app_key", mcGwSourceAppKey)
    return nil
}

// McGwSourceAppKey Getter
func (r TaobaoJstMiniappCrowdUserAddRequest) GetMcGwSourceAppKey() string {
    return r.mcGwSourceAppKey
}
