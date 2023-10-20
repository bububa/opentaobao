package mei

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallmeicrmmembersyncprivyAPIRequest 同步推送会员信息(隐私号版本) API请求
// tmall.mei.crm.member.sync.privy
//
// 品牌方通过该api主动推送会员信息。使用场景包括
// 1.用户在线上注册后，线下补充信息后，同步到线上。
// 2.其他情况的主动推送变更。
type TmallmeicrmmembersyncprivyAPIRequest struct {
	model.Params
	// 会员手机号码
	_mobile string
	// 会员拓展信息
	_extend string
	// ouid
	_ouid string
	// 成长值字段，支持整数或小数(小数位最多两位),当level_type=2时，level_point必传
	_levelPoint string
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

// NewTmallmeicrmmembersyncprivyRequest 初始化TmallmeicrmmembersyncprivyAPIRequest对象
func NewTmallmeicrmmembersyncprivyRequest() *TmallmeicrmmembersyncprivyAPIRequest {
	return &TmallmeicrmmembersyncprivyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallmeicrmmembersyncprivyAPIRequest) GetApiMethodName() string {
	return "tmall.mei.crm.member.sync.privy"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallmeicrmmembersyncprivyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallmeicrmmembersyncprivyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMobile is Mobile Setter
// 会员手机号码
func (r *TmallmeicrmmembersyncprivyAPIRequest) SetMobile(_mobile string) error {
	r._mobile = _mobile
	r.Set("mobile", _mobile)
	return nil
}

// GetMobile Mobile Getter
func (r TmallmeicrmmembersyncprivyAPIRequest) GetMobile() string {
	return r._mobile
}

// SetExtend is Extend Setter
// 会员拓展信息
func (r *TmallmeicrmmembersyncprivyAPIRequest) SetExtend(_extend string) error {
	r._extend = _extend
	r.Set("extend", _extend)
	return nil
}

// GetExtend Extend Getter
func (r TmallmeicrmmembersyncprivyAPIRequest) GetExtend() string {
	return r._extend
}

// SetOuid is Ouid Setter
// ouid
func (r *TmallmeicrmmembersyncprivyAPIRequest) SetOuid(_ouid string) error {
	r._ouid = _ouid
	r.Set("ouid", _ouid)
	return nil
}

// GetOuid Ouid Getter
func (r TmallmeicrmmembersyncprivyAPIRequest) GetOuid() string {
	return r._ouid
}

// SetLevelPoint is LevelPoint Setter
// 成长值字段，支持整数或小数(小数位最多两位),当level_type=2时，level_point必传
func (r *TmallmeicrmmembersyncprivyAPIRequest) SetLevelPoint(_levelPoint string) error {
	r._levelPoint = _levelPoint
	r.Set("level_point", _levelPoint)
	return nil
}

// GetLevelPoint LevelPoint Getter
func (r TmallmeicrmmembersyncprivyAPIRequest) GetLevelPoint() string {
	return r._levelPoint
}

// SetPoint is Point Setter
// 会员积分
func (r *TmallmeicrmmembersyncprivyAPIRequest) SetPoint(_point int64) error {
	r._point = _point
	r.Set("point", _point)
	return nil
}

// GetPoint Point Getter
func (r TmallmeicrmmembersyncprivyAPIRequest) GetPoint() int64 {
	return r._point
}

// SetLevel is Level Setter
// 会员等级
func (r *TmallmeicrmmembersyncprivyAPIRequest) SetLevel(_level int64) error {
	r._level = _level
	r.Set("level", _level)
	return nil
}

// GetLevel Level Getter
func (r TmallmeicrmmembersyncprivyAPIRequest) GetLevel() int64 {
	return r._level
}

// SetVersion is Version Setter
// 该次同步的版本信息（建议使用时间戳）
func (r *TmallmeicrmmembersyncprivyAPIRequest) SetVersion(_version int64) error {
	r._version = _version
	r.Set("version", _version)
	return nil
}

// GetVersion Version Getter
func (r TmallmeicrmmembersyncprivyAPIRequest) GetVersion() int64 {
	return r._version
}

// SetLevelExpireTime is LevelExpireTime Setter
// 毫秒级别时间戳，可不传，表示会员等级有效期
func (r *TmallmeicrmmembersyncprivyAPIRequest) SetLevelExpireTime(_levelExpireTime int64) error {
	r._levelExpireTime = _levelExpireTime
	r.Set("level_expire_time", _levelExpireTime)
	return nil
}

// GetLevelExpireTime LevelExpireTime Getter
func (r TmallmeicrmmembersyncprivyAPIRequest) GetLevelExpireTime() int64 {
	return r._levelExpireTime
}

// SetLevelType is LevelType Setter
// 1(表示传统等级模式), 2(成长值等级模式)
func (r *TmallmeicrmmembersyncprivyAPIRequest) SetLevelType(_levelType int64) error {
	r._levelType = _levelType
	r.Set("level_type", _levelType)
	return nil
}

// GetLevelType LevelType Getter
func (r TmallmeicrmmembersyncprivyAPIRequest) GetLevelType() int64 {
	return r._levelType
}
