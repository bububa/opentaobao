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
    mobile   string
    // 会员积分
    point   int64
    // 会员等级
    level   int64
    // 会员拓展信息
    extend   string
    // 该次同步的版本信息（建议使用时间戳）
    version   int64
    // 混淆昵称
    mixNick   string
    // 昵称
    nick   string
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
func (r *TmallMeiCrmMemberSyncRequest) SetMobile(mobile string) error {
    r.mobile = mobile
    r.Set("mobile", mobile)
    return nil
}

// Mobile Getter
func (r TmallMeiCrmMemberSyncRequest) GetMobile() string {
    return r.mobile
}
// Point Setter
// 会员积分
func (r *TmallMeiCrmMemberSyncRequest) SetPoint(point int64) error {
    r.point = point
    r.Set("point", point)
    return nil
}

// Point Getter
func (r TmallMeiCrmMemberSyncRequest) GetPoint() int64 {
    return r.point
}
// Level Setter
// 会员等级
func (r *TmallMeiCrmMemberSyncRequest) SetLevel(level int64) error {
    r.level = level
    r.Set("level", level)
    return nil
}

// Level Getter
func (r TmallMeiCrmMemberSyncRequest) GetLevel() int64 {
    return r.level
}
// Extend Setter
// 会员拓展信息
func (r *TmallMeiCrmMemberSyncRequest) SetExtend(extend string) error {
    r.extend = extend
    r.Set("extend", extend)
    return nil
}

// Extend Getter
func (r TmallMeiCrmMemberSyncRequest) GetExtend() string {
    return r.extend
}
// Version Setter
// 该次同步的版本信息（建议使用时间戳）
func (r *TmallMeiCrmMemberSyncRequest) SetVersion(version int64) error {
    r.version = version
    r.Set("version", version)
    return nil
}

// Version Getter
func (r TmallMeiCrmMemberSyncRequest) GetVersion() int64 {
    return r.version
}
// MixNick Setter
// 混淆昵称
func (r *TmallMeiCrmMemberSyncRequest) SetMixNick(mixNick string) error {
    r.mixNick = mixNick
    r.Set("mix_nick", mixNick)
    return nil
}

// MixNick Getter
func (r TmallMeiCrmMemberSyncRequest) GetMixNick() string {
    return r.mixNick
}
// Nick Setter
// 昵称
func (r *TmallMeiCrmMemberSyncRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TmallMeiCrmMemberSyncRequest) GetNick() string {
    return r.nick
}
