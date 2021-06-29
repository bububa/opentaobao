package xiamitrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
曲库开放平台内容行为上报接口 API请求
xiami.content.resource.action.report

合作方对接入的曲库开放内容上报行为日志
*/
type XiamiContentResourceActionReportRequest struct {
    model.Params
    // 资源ID
    _resourceId   string
    // 行为数量
    _num   int64
    // 资源类型（可枚举）: song(歌曲)
    _resourceType   string
    // 行为类型（可枚举）：LISTEN（主动试听）、PASSIVE_LISTEN（被动试听）
    _action   string
    // 来源id，如歌单id
    _fromId   string
    // 1推荐2歌单3标签
    _fromType   int64
    // 用户id
    _openId   string
    // 用户设备id
    _utdid   string
}

// 初始化XiamiContentResourceActionReportRequest对象
func NewXiamiContentResourceActionReportRequest() *XiamiContentResourceActionReportRequest{
    return &XiamiContentResourceActionReportRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r XiamiContentResourceActionReportRequest) GetApiMethodName() string {
    return "xiami.content.resource.action.report"
}

// IRequest interface 方法, 获取API参数
func (r XiamiContentResourceActionReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ResourceId Setter
// 资源ID
func (r *XiamiContentResourceActionReportRequest) SetResourceId(_resourceId string) error {
    r._resourceId = _resourceId
    r.Set("resource_id", _resourceId)
    return nil
}

// ResourceId Getter
func (r XiamiContentResourceActionReportRequest) GetResourceId() string {
    return r._resourceId
}
// Num Setter
// 行为数量
func (r *XiamiContentResourceActionReportRequest) SetNum(_num int64) error {
    r._num = _num
    r.Set("num", _num)
    return nil
}

// Num Getter
func (r XiamiContentResourceActionReportRequest) GetNum() int64 {
    return r._num
}
// ResourceType Setter
// 资源类型（可枚举）: song(歌曲)
func (r *XiamiContentResourceActionReportRequest) SetResourceType(_resourceType string) error {
    r._resourceType = _resourceType
    r.Set("resource_type", _resourceType)
    return nil
}

// ResourceType Getter
func (r XiamiContentResourceActionReportRequest) GetResourceType() string {
    return r._resourceType
}
// Action Setter
// 行为类型（可枚举）：LISTEN（主动试听）、PASSIVE_LISTEN（被动试听）
func (r *XiamiContentResourceActionReportRequest) SetAction(_action string) error {
    r._action = _action
    r.Set("action", _action)
    return nil
}

// Action Getter
func (r XiamiContentResourceActionReportRequest) GetAction() string {
    return r._action
}
// FromId Setter
// 来源id，如歌单id
func (r *XiamiContentResourceActionReportRequest) SetFromId(_fromId string) error {
    r._fromId = _fromId
    r.Set("from_id", _fromId)
    return nil
}

// FromId Getter
func (r XiamiContentResourceActionReportRequest) GetFromId() string {
    return r._fromId
}
// FromType Setter
// 1推荐2歌单3标签
func (r *XiamiContentResourceActionReportRequest) SetFromType(_fromType int64) error {
    r._fromType = _fromType
    r.Set("from_type", _fromType)
    return nil
}

// FromType Getter
func (r XiamiContentResourceActionReportRequest) GetFromType() int64 {
    return r._fromType
}
// OpenId Setter
// 用户id
func (r *XiamiContentResourceActionReportRequest) SetOpenId(_openId string) error {
    r._openId = _openId
    r.Set("open_id", _openId)
    return nil
}

// OpenId Getter
func (r XiamiContentResourceActionReportRequest) GetOpenId() string {
    return r._openId
}
// Utdid Setter
// 用户设备id
func (r *XiamiContentResourceActionReportRequest) SetUtdid(_utdid string) error {
    r._utdid = _utdid
    r.Set("utdid", _utdid)
    return nil
}

// Utdid Getter
func (r XiamiContentResourceActionReportRequest) GetUtdid() string {
    return r._utdid
}
