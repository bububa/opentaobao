package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌营销活动用户上传 API请求
tmall.ccf.crowd.activityuser.upload

搜集ISV的活动用户信息，将其沉淀为活动人群数据
*/
type TmallCcfCrowdActivityuserUploadRequest struct {
    model.Params
    // 活动id
    _activityId   int64
    // 人群类型
    _crowdTypes   []string
    // 淘宝小程序的openid
    _taobaoOpenId   string
    // 小程序对应的appKey
    _taobaoAppKey   string
}

// 初始化TmallCcfCrowdActivityuserUploadRequest对象
func NewTmallCcfCrowdActivityuserUploadRequest() *TmallCcfCrowdActivityuserUploadRequest{
    return &TmallCcfCrowdActivityuserUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallCcfCrowdActivityuserUploadRequest) GetApiMethodName() string {
    return "tmall.ccf.crowd.activityuser.upload"
}

// IRequest interface 方法, 获取API参数
func (r TmallCcfCrowdActivityuserUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ActivityId Setter
// 活动id
func (r *TmallCcfCrowdActivityuserUploadRequest) SetActivityId(_activityId int64) error {
    r._activityId = _activityId
    r.Set("activity_id", _activityId)
    return nil
}

// ActivityId Getter
func (r TmallCcfCrowdActivityuserUploadRequest) GetActivityId() int64 {
    return r._activityId
}
// CrowdTypes Setter
// 人群类型
func (r *TmallCcfCrowdActivityuserUploadRequest) SetCrowdTypes(_crowdTypes []string) error {
    r._crowdTypes = _crowdTypes
    r.Set("crowd_types", _crowdTypes)
    return nil
}

// CrowdTypes Getter
func (r TmallCcfCrowdActivityuserUploadRequest) GetCrowdTypes() []string {
    return r._crowdTypes
}
// TaobaoOpenId Setter
// 淘宝小程序的openid
func (r *TmallCcfCrowdActivityuserUploadRequest) SetTaobaoOpenId(_taobaoOpenId string) error {
    r._taobaoOpenId = _taobaoOpenId
    r.Set("taobao_open_id", _taobaoOpenId)
    return nil
}

// TaobaoOpenId Getter
func (r TmallCcfCrowdActivityuserUploadRequest) GetTaobaoOpenId() string {
    return r._taobaoOpenId
}
// TaobaoAppKey Setter
// 小程序对应的appKey
func (r *TmallCcfCrowdActivityuserUploadRequest) SetTaobaoAppKey(_taobaoAppKey string) error {
    r._taobaoAppKey = _taobaoAppKey
    r.Set("taobao_app_key", _taobaoAppKey)
    return nil
}

// TaobaoAppKey Getter
func (r TmallCcfCrowdActivityuserUploadRequest) GetTaobaoAppKey() string {
    return r._taobaoAppKey
}
