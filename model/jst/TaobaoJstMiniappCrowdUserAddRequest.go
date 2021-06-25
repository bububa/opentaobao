package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
小程序添加用户到指定的活动 APIRequest
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

func NewTaobaoJstMiniappCrowdUserAddRequest() *TaobaoJstMiniappCrowdUserAddRequest{
    return &TaobaoJstMiniappCrowdUserAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJstMiniappCrowdUserAddRequest) GetApiMethodName() string {
    return "taobao.jst.miniapp.crowd.user.add"
}

func (r TaobaoJstMiniappCrowdUserAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoJstMiniappCrowdUserAddRequest) SetCrowdCode(crowdCode string) error {
    r.crowdCode = crowdCode
    r.Set("crowd_code", crowdCode)
    return nil
}

func (r TaobaoJstMiniappCrowdUserAddRequest) GetCrowdCode() string {
    return r.crowdCode
}

func (r *TaobaoJstMiniappCrowdUserAddRequest) SetMcGwSourceMiniAppId(mcGwSourceMiniAppId string) error {
    r.mcGwSourceMiniAppId = mcGwSourceMiniAppId
    r.Set("mc_gw_source_mini_app_id", mcGwSourceMiniAppId)
    return nil
}

func (r TaobaoJstMiniappCrowdUserAddRequest) GetMcGwSourceMiniAppId() string {
    return r.mcGwSourceMiniAppId
}

func (r *TaobaoJstMiniappCrowdUserAddRequest) SetMcGwSourceAppKey(mcGwSourceAppKey string) error {
    r.mcGwSourceAppKey = mcGwSourceAppKey
    r.Set("mc_gw_source_app_key", mcGwSourceAppKey)
    return nil
}

func (r TaobaoJstMiniappCrowdUserAddRequest) GetMcGwSourceAppKey() string {
    return r.mcGwSourceAppKey
}

