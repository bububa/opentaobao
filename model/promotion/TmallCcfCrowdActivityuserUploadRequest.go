package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌营销活动用户上传 APIRequest
tmall.ccf.crowd.activityuser.upload

搜集ISV的活动用户信息，将其沉淀为活动人群数据
*/
type TmallCcfCrowdActivityuserUploadRequest struct {
    model.Params

    // 活动id
    activityId   int64 

    // 人群类型
    crowdTypes   []string 

    // 淘宝小程序的openid
    taobaoOpenId   string 

    // 小程序对应的appKey
    taobaoAppKey   string 

}

func NewTmallCcfCrowdActivityuserUploadRequest() *TmallCcfCrowdActivityuserUploadRequest{
    return &TmallCcfCrowdActivityuserUploadRequest{
        Params: model.NewParams(),
    }
}

func (r TmallCcfCrowdActivityuserUploadRequest) GetApiMethodName() string {
    return "tmall.ccf.crowd.activityuser.upload"
}

func (r TmallCcfCrowdActivityuserUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallCcfCrowdActivityuserUploadRequest) SetActivityId(activityId int64) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

func (r TmallCcfCrowdActivityuserUploadRequest) GetActivityId() int64 {
    return r.activityId
}

func (r *TmallCcfCrowdActivityuserUploadRequest) SetCrowdTypes(crowdTypes []string) error {
    r.crowdTypes = crowdTypes
    r.Set("crowd_types", crowdTypes)
    return nil
}

func (r TmallCcfCrowdActivityuserUploadRequest) GetCrowdTypes() []string {
    return r.crowdTypes
}

func (r *TmallCcfCrowdActivityuserUploadRequest) SetTaobaoOpenId(taobaoOpenId string) error {
    r.taobaoOpenId = taobaoOpenId
    r.Set("taobao_open_id", taobaoOpenId)
    return nil
}

func (r TmallCcfCrowdActivityuserUploadRequest) GetTaobaoOpenId() string {
    return r.taobaoOpenId
}

func (r *TmallCcfCrowdActivityuserUploadRequest) SetTaobaoAppKey(taobaoAppKey string) error {
    r.taobaoAppKey = taobaoAppKey
    r.Set("taobao_app_key", taobaoAppKey)
    return nil
}

func (r TmallCcfCrowdActivityuserUploadRequest) GetTaobaoAppKey() string {
    return r.taobaoAppKey
}

