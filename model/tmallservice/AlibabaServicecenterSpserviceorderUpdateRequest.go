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
    _spServiceOrderId   int64
    // 新设备sn.当action填写addSn、changeSn时必填
    _action   string
    // 新设备sn.当action填写addSn、changeSn时必填
    _newSn   string
    // 旧设备sn，当action填写changeSn时必填
    _oldSn   string
    // 服务有效期开始时间
    _gmtEffect   string
    // 服务过期时间
    _gmtExpire   string
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
func (r *AlibabaServicecenterSpserviceorderUpdateRequest) SetSpServiceOrderId(_spServiceOrderId int64) error {
    r._spServiceOrderId = _spServiceOrderId
    r.Set("sp_service_order_id", _spServiceOrderId)
    return nil
}

// SpServiceOrderId Getter
func (r AlibabaServicecenterSpserviceorderUpdateRequest) GetSpServiceOrderId() int64 {
    return r._spServiceOrderId
}
// Action Setter
// 新设备sn.当action填写addSn、changeSn时必填
func (r *AlibabaServicecenterSpserviceorderUpdateRequest) SetAction(_action string) error {
    r._action = _action
    r.Set("action", _action)
    return nil
}

// Action Getter
func (r AlibabaServicecenterSpserviceorderUpdateRequest) GetAction() string {
    return r._action
}
// NewSn Setter
// 新设备sn.当action填写addSn、changeSn时必填
func (r *AlibabaServicecenterSpserviceorderUpdateRequest) SetNewSn(_newSn string) error {
    r._newSn = _newSn
    r.Set("new_sn", _newSn)
    return nil
}

// NewSn Getter
func (r AlibabaServicecenterSpserviceorderUpdateRequest) GetNewSn() string {
    return r._newSn
}
// OldSn Setter
// 旧设备sn，当action填写changeSn时必填
func (r *AlibabaServicecenterSpserviceorderUpdateRequest) SetOldSn(_oldSn string) error {
    r._oldSn = _oldSn
    r.Set("old_sn", _oldSn)
    return nil
}

// OldSn Getter
func (r AlibabaServicecenterSpserviceorderUpdateRequest) GetOldSn() string {
    return r._oldSn
}
// GmtEffect Setter
// 服务有效期开始时间
func (r *AlibabaServicecenterSpserviceorderUpdateRequest) SetGmtEffect(_gmtEffect string) error {
    r._gmtEffect = _gmtEffect
    r.Set("gmt_effect", _gmtEffect)
    return nil
}

// GmtEffect Getter
func (r AlibabaServicecenterSpserviceorderUpdateRequest) GetGmtEffect() string {
    return r._gmtEffect
}
// GmtExpire Setter
// 服务过期时间
func (r *AlibabaServicecenterSpserviceorderUpdateRequest) SetGmtExpire(_gmtExpire string) error {
    r._gmtExpire = _gmtExpire
    r.Set("gmt_expire", _gmtExpire)
    return nil
}

// GmtExpire Getter
func (r AlibabaServicecenterSpserviceorderUpdateRequest) GetGmtExpire() string {
    return r._gmtExpire
}
