package mei

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
同步推送会员信息 API请求
tmall.mei.crm.member.sync

品牌方通过该api主动推送会员信息。使用场景包括
1.用户在线上注册后，线下补充信息后，同步到线上。
2.其他情况的主动推送变更。
*/
type TmallMeiCrmMemberSyncRequest struct {
    model.Params
    // 会员手机号码
    _mobile   string
    // 会员积分
    _point   int64
    // 会员等级
    _level   int64
    // 会员拓展信息
    _extend   string
    // 该次同步的版本信息（建议使用时间戳）
    _version   int64
    // 混淆昵称
    _mixNick   string
    // 昵称
    _nick   string
}

// 初始化TmallMeiCrmMemberSyncRequest对象
func NewTmallMeiCrmMemberSyncRequest() *TmallMeiCrmMemberSyncRequest{
    return &TmallMeiCrmMemberSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallMeiCrmMemberSyncRequest) GetApiMethodName() string {
    return "tmall.mei.crm.member.sync"
}

// IRequest interface 方法, 获取API参数
func (r TmallMeiCrmMemberSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Mobile Setter
// 会员手机号码
func (r *TmallMeiCrmMemberSyncRequest) SetMobile(_mobile string) error {
    r._mobile = _mobile
    r.Set("mobile", _mobile)
    return nil
}

// Mobile Getter
func (r TmallMeiCrmMemberSyncRequest) GetMobile() string {
    return r._mobile
}
// Point Setter
// 会员积分
func (r *TmallMeiCrmMemberSyncRequest) SetPoint(_point int64) error {
    r._point = _point
    r.Set("point", _point)
    return nil
}

// Point Getter
func (r TmallMeiCrmMemberSyncRequest) GetPoint() int64 {
    return r._point
}
// Level Setter
// 会员等级
func (r *TmallMeiCrmMemberSyncRequest) SetLevel(_level int64) error {
    r._level = _level
    r.Set("level", _level)
    return nil
}

// Level Getter
func (r TmallMeiCrmMemberSyncRequest) GetLevel() int64 {
    return r._level
}
// Extend Setter
// 会员拓展信息
func (r *TmallMeiCrmMemberSyncRequest) SetExtend(_extend string) error {
    r._extend = _extend
    r.Set("extend", _extend)
    return nil
}

// Extend Getter
func (r TmallMeiCrmMemberSyncRequest) GetExtend() string {
    return r._extend
}
// Version Setter
// 该次同步的版本信息（建议使用时间戳）
func (r *TmallMeiCrmMemberSyncRequest) SetVersion(_version int64) error {
    r._version = _version
    r.Set("version", _version)
    return nil
}

// Version Getter
func (r TmallMeiCrmMemberSyncRequest) GetVersion() int64 {
    return r._version
}
// MixNick Setter
// 混淆昵称
func (r *TmallMeiCrmMemberSyncRequest) SetMixNick(_mixNick string) error {
    r._mixNick = _mixNick
    r.Set("mix_nick", _mixNick)
    return nil
}

// MixNick Getter
func (r TmallMeiCrmMemberSyncRequest) GetMixNick() string {
    return r._mixNick
}
// Nick Setter
// 昵称
func (r *TmallMeiCrmMemberSyncRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TmallMeiCrmMemberSyncRequest) GetNick() string {
    return r._nick
}
