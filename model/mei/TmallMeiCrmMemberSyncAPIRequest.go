package mei

import (
	"net/url"
	"sync"

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
	// 会员拓展信息
	_extend string
	// 成长值字段，支持整数或小数(小数位最多两位),当level_type=2时，level_point必传
	_levelPoint string
	// 混淆昵称
	_mixNick string
	// 昵称
	_nick string
	// 会员积分
	_point int64
	// 会员等级
	_level int64
	// 该次同步的版本信息（建议使用时间戳）
	_version int64
	// 毫秒级别时间戳，可不传，表示会员等级有效期
	_levelExpireTime int64
	// 1(表示传统等级模式), 2(成长值等级模式)
	_levelType int64
}

// NewTmallMeiCrmMemberSyncRequest 初始化TmallMeiCrmMemberSyncAPIRequest对象
func NewTmallMeiCrmMemberSyncRequest() *TmallMeiCrmMemberSyncAPIRequest {
	return &TmallMeiCrmMemberSyncAPIRequest{
		Params: model.NewParams(10),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallMeiCrmMemberSyncAPIRequest) Reset() {
	r._mobile = ""
	r._extend = ""
	r._levelPoint = ""
	r._mixNick = ""
	r._nick = ""
	r._point = 0
	r._level = 0
	r._version = 0
	r._levelExpireTime = 0
	r._levelType = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallMeiCrmMemberSyncAPIRequest) GetApiMethodName() string {
	return "tmall.mei.crm.member.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallMeiCrmMemberSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallMeiCrmMemberSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMobile is Mobile Setter
// 会员手机号码
func (r *TmallMeiCrmMemberSyncAPIRequest) SetMobile(_mobile string) error {
	r._mobile = _mobile
	r.Set("mobile", _mobile)
	return nil
}

// GetMobile Mobile Getter
func (r TmallMeiCrmMemberSyncAPIRequest) GetMobile() string {
	return r._mobile
}

// SetExtend is Extend Setter
// 会员拓展信息
func (r *TmallMeiCrmMemberSyncAPIRequest) SetExtend(_extend string) error {
	r._extend = _extend
	r.Set("extend", _extend)
	return nil
}

// GetExtend Extend Getter
func (r TmallMeiCrmMemberSyncAPIRequest) GetExtend() string {
	return r._extend
}

// SetLevelPoint is LevelPoint Setter
// 成长值字段，支持整数或小数(小数位最多两位),当level_type=2时，level_point必传
func (r *TmallMeiCrmMemberSyncAPIRequest) SetLevelPoint(_levelPoint string) error {
	r._levelPoint = _levelPoint
	r.Set("level_point", _levelPoint)
	return nil
}

// GetLevelPoint LevelPoint Getter
func (r TmallMeiCrmMemberSyncAPIRequest) GetLevelPoint() string {
	return r._levelPoint
}

// SetMixNick is MixNick Setter
// 混淆昵称
func (r *TmallMeiCrmMemberSyncAPIRequest) SetMixNick(_mixNick string) error {
	r._mixNick = _mixNick
	r.Set("mix_nick", _mixNick)
	return nil
}

// GetMixNick MixNick Getter
func (r TmallMeiCrmMemberSyncAPIRequest) GetMixNick() string {
	return r._mixNick
}

// SetNick is Nick Setter
// 昵称
func (r *TmallMeiCrmMemberSyncAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TmallMeiCrmMemberSyncAPIRequest) GetNick() string {
	return r._nick
}

// SetPoint is Point Setter
// 会员积分
func (r *TmallMeiCrmMemberSyncAPIRequest) SetPoint(_point int64) error {
	r._point = _point
	r.Set("point", _point)
	return nil
}

// GetPoint Point Getter
func (r TmallMeiCrmMemberSyncAPIRequest) GetPoint() int64 {
	return r._point
}

// SetLevel is Level Setter
// 会员等级
func (r *TmallMeiCrmMemberSyncAPIRequest) SetLevel(_level int64) error {
	r._level = _level
	r.Set("level", _level)
	return nil
}

// GetLevel Level Getter
func (r TmallMeiCrmMemberSyncAPIRequest) GetLevel() int64 {
	return r._level
}

// SetVersion is Version Setter
// 该次同步的版本信息（建议使用时间戳）
func (r *TmallMeiCrmMemberSyncAPIRequest) SetVersion(_version int64) error {
	r._version = _version
	r.Set("version", _version)
	return nil
}

// GetVersion Version Getter
func (r TmallMeiCrmMemberSyncAPIRequest) GetVersion() int64 {
	return r._version
}

// SetLevelExpireTime is LevelExpireTime Setter
// 毫秒级别时间戳，可不传，表示会员等级有效期
func (r *TmallMeiCrmMemberSyncAPIRequest) SetLevelExpireTime(_levelExpireTime int64) error {
	r._levelExpireTime = _levelExpireTime
	r.Set("level_expire_time", _levelExpireTime)
	return nil
}

// GetLevelExpireTime LevelExpireTime Getter
func (r TmallMeiCrmMemberSyncAPIRequest) GetLevelExpireTime() int64 {
	return r._levelExpireTime
}

// SetLevelType is LevelType Setter
// 1(表示传统等级模式), 2(成长值等级模式)
func (r *TmallMeiCrmMemberSyncAPIRequest) SetLevelType(_levelType int64) error {
	r._levelType = _levelType
	r.Set("level_type", _levelType)
	return nil
}

// GetLevelType LevelType Getter
func (r TmallMeiCrmMemberSyncAPIRequest) GetLevelType() int64 {
	return r._levelType
}

var poolTmallMeiCrmMemberSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallMeiCrmMemberSyncRequest()
	},
}

// GetTmallMeiCrmMemberSyncRequest 从 sync.Pool 获取 TmallMeiCrmMemberSyncAPIRequest
func GetTmallMeiCrmMemberSyncAPIRequest() *TmallMeiCrmMemberSyncAPIRequest {
	return poolTmallMeiCrmMemberSyncAPIRequest.Get().(*TmallMeiCrmMemberSyncAPIRequest)
}

// ReleaseTmallMeiCrmMemberSyncAPIRequest 将 TmallMeiCrmMemberSyncAPIRequest 放入 sync.Pool
func ReleaseTmallMeiCrmMemberSyncAPIRequest(v *TmallMeiCrmMemberSyncAPIRequest) {
	v.Reset()
	poolTmallMeiCrmMemberSyncAPIRequest.Put(v)
}
