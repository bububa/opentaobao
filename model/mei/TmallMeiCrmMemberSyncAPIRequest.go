package mei

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallMeiCrmMemberSyncAPIRequest 同步推送会员信息 API请求
// tmall.mei.crm.member.sync
//
// 品牌方通过该api主动推送会员信息。使用场景包括
// 1.用户在线上注册后，线下补充信息后，同步到线上。
// 2.其他情况的主动推送变更。
type TmallMeiCrmMemberSyncAPIRequest struct {
	model.Params
	// 会员手机号码
	_mobile string
	// 会员积分
	_point int64
	// 会员等级
	_level int64
	// 会员拓展信息
	_extend string
	// 该次同步的版本信息（建议使用时间戳）
	_version int64
	// 混淆昵称
	_mixNick string
	// 昵称
	_nick string
}

// NewTmallMeiCrmMemberSyncRequest 初始化TmallMeiCrmMemberSyncAPIRequest对象
func NewTmallMeiCrmMemberSyncRequest() *TmallMeiCrmMemberSyncAPIRequest {
	return &TmallMeiCrmMemberSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallMeiCrmMemberSyncAPIRequest) GetApiMethodName() string {
	return "tmall.mei.crm.member.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallMeiCrmMemberSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Mobile Setter
// 会员手机号码
func (r *TmallMeiCrmMemberSyncAPIRequest) SetMobile(_mobile string) error {
	r._mobile = _mobile
	r.Set("mobile", _mobile)
	return nil
}

// Get Mobile Getter
func (r TmallMeiCrmMemberSyncAPIRequest) GetMobile() string {
	return r._mobile
}

// Set is Point Setter
// 会员积分
func (r *TmallMeiCrmMemberSyncAPIRequest) SetPoint(_point int64) error {
	r._point = _point
	r.Set("point", _point)
	return nil
}

// Get Point Getter
func (r TmallMeiCrmMemberSyncAPIRequest) GetPoint() int64 {
	return r._point
}

// Set is Level Setter
// 会员等级
func (r *TmallMeiCrmMemberSyncAPIRequest) SetLevel(_level int64) error {
	r._level = _level
	r.Set("level", _level)
	return nil
}

// Get Level Getter
func (r TmallMeiCrmMemberSyncAPIRequest) GetLevel() int64 {
	return r._level
}

// Set is Extend Setter
// 会员拓展信息
func (r *TmallMeiCrmMemberSyncAPIRequest) SetExtend(_extend string) error {
	r._extend = _extend
	r.Set("extend", _extend)
	return nil
}

// Get Extend Getter
func (r TmallMeiCrmMemberSyncAPIRequest) GetExtend() string {
	return r._extend
}

// Set is Version Setter
// 该次同步的版本信息（建议使用时间戳）
func (r *TmallMeiCrmMemberSyncAPIRequest) SetVersion(_version int64) error {
	r._version = _version
	r.Set("version", _version)
	return nil
}

// Get Version Getter
func (r TmallMeiCrmMemberSyncAPIRequest) GetVersion() int64 {
	return r._version
}

// Set is MixNick Setter
// 混淆昵称
func (r *TmallMeiCrmMemberSyncAPIRequest) SetMixNick(_mixNick string) error {
	r._mixNick = _mixNick
	r.Set("mix_nick", _mixNick)
	return nil
}

// Get MixNick Getter
func (r TmallMeiCrmMemberSyncAPIRequest) GetMixNick() string {
	return r._mixNick
}

// Set is Nick Setter
// 昵称
func (r *TmallMeiCrmMemberSyncAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// Get Nick Getter
func (r TmallMeiCrmMemberSyncAPIRequest) GetNick() string {
	return r._nick
}
