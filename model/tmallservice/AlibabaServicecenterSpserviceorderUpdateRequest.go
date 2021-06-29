package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务供应链服务单更新 API请求
alibaba.servicecenter.spserviceorder.update

服务供应链服务单更新，服务商通过此接口将商品的sn等信息推送到服务单中
*/
type AlibabaServicecenterSpserviceorderUpdateRequest struct {
    model.Params
    // 服务单id
    spServiceOrderId   int64
    // 新设备sn.当action填写addSn、changeSn时必填
    action   string
    // 新设备sn.当action填写addSn、changeSn时必填
    newSn   string
    // 旧设备sn，当action填写changeSn时必填
    oldSn   string
    // 服务有效期开始时间
    gmtEffect   string
    // 服务过期时间
    gmtExpire   string
}

// 初始化AlibabaServicecenterSpserviceorderUpdateRequest对象
func NewAlibabaServicecenterSpserviceorderUpdateRequest() *AlibabaServicecenterSpserviceorderUpdateRequest{
    return &AlibabaServicecenterSpserviceorderUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaServicecenterSpserviceorderUpdateRequest) GetApiMethodName() string {
    return "alibaba.servicecenter.spserviceorder.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaServicecenterSpserviceorderUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SpServiceOrderId Setter
// 服务单id
func (r *AlibabaServicecenterSpserviceorderUpdateRequest) SetSpServiceOrderId(spServiceOrderId int64) error {
    r.spServiceOrderId = spServiceOrderId
    r.Set("sp_service_order_id", spServiceOrderId)
    return nil
}

// SpServiceOrderId Getter
func (r AlibabaServicecenterSpserviceorderUpdateRequest) GetSpServiceOrderId() int64 {
    return r.spServiceOrderId
}
// Action Setter
// 新设备sn.当action填写addSn、changeSn时必填
func (r *AlibabaServicecenterSpserviceorderUpdateRequest) SetAction(action string) error {
    r.action = action
    r.Set("action", action)
    return nil
}

// Action Getter
func (r AlibabaServicecenterSpserviceorderUpdateRequest) GetAction() string {
    return r.action
}
// NewSn Setter
// 新设备sn.当action填写addSn、changeSn时必填
func (r *AlibabaServicecenterSpserviceorderUpdateRequest) SetNewSn(newSn string) error {
    r.newSn = newSn
    r.Set("new_sn", newSn)
    return nil
}

// NewSn Getter
func (r AlibabaServicecenterSpserviceorderUpdateRequest) GetNewSn() string {
    return r.newSn
}
// OldSn Setter
// 旧设备sn，当action填写changeSn时必填
func (r *AlibabaServicecenterSpserviceorderUpdateRequest) SetOldSn(oldSn string) error {
    r.oldSn = oldSn
    r.Set("old_sn", oldSn)
    return nil
}

// OldSn Getter
func (r AlibabaServicecenterSpserviceorderUpdateRequest) GetOldSn() string {
    return r.oldSn
}
// GmtEffect Setter
// 服务有效期开始时间
func (r *AlibabaServicecenterSpserviceorderUpdateRequest) SetGmtEffect(gmtEffect string) error {
    r.gmtEffect = gmtEffect
    r.Set("gmt_effect", gmtEffect)
    return nil
}

// GmtEffect Getter
func (r AlibabaServicecenterSpserviceorderUpdateRequest) GetGmtEffect() string {
    return r.gmtEffect
}
// GmtExpire Setter
// 服务过期时间
func (r *AlibabaServicecenterSpserviceorderUpdateRequest) SetGmtExpire(gmtExpire string) error {
    r.gmtExpire = gmtExpire
    r.Set("gmt_expire", gmtExpire)
    return nil
}

// GmtExpire Getter
func (r AlibabaServicecenterSpserviceorderUpdateRequest) GetGmtExpire() string {
    return r.gmtExpire
}
