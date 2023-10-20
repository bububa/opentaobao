package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalisportsdatasportssyncsleepdataAPIRequest 阿里体育数据中心用户睡眠数据同步接口 API请求
// alibaba.alisports.data.sports.syncsleepdata
//
// 阿里体育数据中心用户睡眠数据同步接口
type AlibabaalisportsdatasportssyncsleepdataAPIRequest struct {
	model.Params
	// 应用appkey
	_alispAppKey string
	// 入睡时间，格式：y-m-d h:i:s
	_stime string
	// 清醒时间，单位：小时
	_soberTime string
	// 浅度睡眠时间，单位：小时
	_shallowTime string
	// 深度睡眠时间，单位：小时
	_deepTime string
	// 睡眠总时间，单位：小时
	_allTime string
	// 阿里体育用户id
	_aliuid string
	// 接口签名
	_alispSign string
	// 时间戳精确到秒
	_alispTime string
	// 醒来时间，格式：y-m-d h:i:s
	_etime string
}

// NewAlibabaalisportsdatasportssyncsleepdataRequest 初始化AlibabaalisportsdatasportssyncsleepdataAPIRequest对象
func NewAlibabaalisportsdatasportssyncsleepdataRequest() *AlibabaalisportsdatasportssyncsleepdataAPIRequest {
	return &AlibabaalisportsdatasportssyncsleepdataAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalisportsdatasportssyncsleepdataAPIRequest) GetApiMethodName() string {
	return "alibaba.alisports.data.sports.syncsleepdata"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalisportsdatasportssyncsleepdataAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalisportsdatasportssyncsleepdataAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlispAppKey is AlispAppKey Setter
// 应用appkey
func (r *AlibabaalisportsdatasportssyncsleepdataAPIRequest) SetAlispAppKey(_alispAppKey string) error {
	r._alispAppKey = _alispAppKey
	r.Set("alisp_app_key", _alispAppKey)
	return nil
}

// GetAlispAppKey AlispAppKey Getter
func (r AlibabaalisportsdatasportssyncsleepdataAPIRequest) GetAlispAppKey() string {
	return r._alispAppKey
}

// SetStime is Stime Setter
// 入睡时间，格式：y-m-d h:i:s
func (r *AlibabaalisportsdatasportssyncsleepdataAPIRequest) SetStime(_stime string) error {
	r._stime = _stime
	r.Set("stime", _stime)
	return nil
}

// GetStime Stime Getter
func (r AlibabaalisportsdatasportssyncsleepdataAPIRequest) GetStime() string {
	return r._stime
}

// SetSoberTime is SoberTime Setter
// 清醒时间，单位：小时
func (r *AlibabaalisportsdatasportssyncsleepdataAPIRequest) SetSoberTime(_soberTime string) error {
	r._soberTime = _soberTime
	r.Set("sober_time", _soberTime)
	return nil
}

// GetSoberTime SoberTime Getter
func (r AlibabaalisportsdatasportssyncsleepdataAPIRequest) GetSoberTime() string {
	return r._soberTime
}

// SetShallowTime is ShallowTime Setter
// 浅度睡眠时间，单位：小时
func (r *AlibabaalisportsdatasportssyncsleepdataAPIRequest) SetShallowTime(_shallowTime string) error {
	r._shallowTime = _shallowTime
	r.Set("shallow_time", _shallowTime)
	return nil
}

// GetShallowTime ShallowTime Getter
func (r AlibabaalisportsdatasportssyncsleepdataAPIRequest) GetShallowTime() string {
	return r._shallowTime
}

// SetDeepTime is DeepTime Setter
// 深度睡眠时间，单位：小时
func (r *AlibabaalisportsdatasportssyncsleepdataAPIRequest) SetDeepTime(_deepTime string) error {
	r._deepTime = _deepTime
	r.Set("deep_time", _deepTime)
	return nil
}

// GetDeepTime DeepTime Getter
func (r AlibabaalisportsdatasportssyncsleepdataAPIRequest) GetDeepTime() string {
	return r._deepTime
}

// SetAllTime is AllTime Setter
// 睡眠总时间，单位：小时
func (r *AlibabaalisportsdatasportssyncsleepdataAPIRequest) SetAllTime(_allTime string) error {
	r._allTime = _allTime
	r.Set("all_time", _allTime)
	return nil
}

// GetAllTime AllTime Getter
func (r AlibabaalisportsdatasportssyncsleepdataAPIRequest) GetAllTime() string {
	return r._allTime
}

// SetAliuid is Aliuid Setter
// 阿里体育用户id
func (r *AlibabaalisportsdatasportssyncsleepdataAPIRequest) SetAliuid(_aliuid string) error {
	r._aliuid = _aliuid
	r.Set("aliuid", _aliuid)
	return nil
}

// GetAliuid Aliuid Getter
func (r AlibabaalisportsdatasportssyncsleepdataAPIRequest) GetAliuid() string {
	return r._aliuid
}

// SetAlispSign is AlispSign Setter
// 接口签名
func (r *AlibabaalisportsdatasportssyncsleepdataAPIRequest) SetAlispSign(_alispSign string) error {
	r._alispSign = _alispSign
	r.Set("alisp_sign", _alispSign)
	return nil
}

// GetAlispSign AlispSign Getter
func (r AlibabaalisportsdatasportssyncsleepdataAPIRequest) GetAlispSign() string {
	return r._alispSign
}

// SetAlispTime is AlispTime Setter
// 时间戳精确到秒
func (r *AlibabaalisportsdatasportssyncsleepdataAPIRequest) SetAlispTime(_alispTime string) error {
	r._alispTime = _alispTime
	r.Set("alisp_time", _alispTime)
	return nil
}

// GetAlispTime AlispTime Getter
func (r AlibabaalisportsdatasportssyncsleepdataAPIRequest) GetAlispTime() string {
	return r._alispTime
}

// SetEtime is Etime Setter
// 醒来时间，格式：y-m-d h:i:s
func (r *AlibabaalisportsdatasportssyncsleepdataAPIRequest) SetEtime(_etime string) error {
	r._etime = _etime
	r.Set("etime", _etime)
	return nil
}

// GetEtime Etime Getter
func (r AlibabaalisportsdatasportssyncsleepdataAPIRequest) GetEtime() string {
	return r._etime
}
