package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkDgVegasTljCreateAPIRequest 淘宝客-推广者-淘礼金创建 API请求
// taobao.tbk.dg.vegas.tlj.create
//
// 创建淘礼金
type TaobaoTbkDgVegasTljCreateAPIRequest struct {
	model.Params
	// 使用开始日期。相对时间，无需填写，以用户领取时间作为使用开始时间。绝对时间，格式 yyyy-MM-dd，例如，2019-01-29，表示从2019-01-29 00:00:00 开始
	_useStartTime string
	// 使用结束日期。如果是结束时间模式为相对时间，时间格式为1-7直接的整数, 例如，1（相对领取时间1天）； 如果是绝对时间，格式为yyyy-MM-dd，例如，2019-01-29，表示到2019-01-29 23:59:59结束
	_useEndTime string
	// 发放截止时间
	_sendEndTime string
	// 发放开始时间
	_sendStartTime string
	// 单个淘礼金面额，支持两位小数，单位元
	_perFace string
	// 淘礼金名称，最大10个字符
	_name string
	// 宝贝ID或营销ID
	_itemId string
	// 已下线，后续不需要填写
	_campaignType string
	// 必须传入0
	_securityLevel int64
	// 结束日期的模式,1:相对时间，2:绝对时间
	_useEndTimeMode int64
	// 单用户累计中奖次数上限
	_userTotalWinNumLimit int64
	// 淘礼金总个数
	_totalNum int64
	// 妈妈广告位Id
	_adzoneId int64
	// 必须设置为true，默认开启安全
	_securitySwitch bool
}

// NewTaobaoTbkDgVegasTljCreateRequest 初始化TaobaoTbkDgVegasTljCreateAPIRequest对象
func NewTaobaoTbkDgVegasTljCreateRequest() *TaobaoTbkDgVegasTljCreateAPIRequest {
	return &TaobaoTbkDgVegasTljCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkDgVegasTljCreateAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.dg.vegas.tlj.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkDgVegasTljCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTbkDgVegasTljCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUseStartTime is UseStartTime Setter
// 使用开始日期。相对时间，无需填写，以用户领取时间作为使用开始时间。绝对时间，格式 yyyy-MM-dd，例如，2019-01-29，表示从2019-01-29 00:00:00 开始
func (r *TaobaoTbkDgVegasTljCreateAPIRequest) SetUseStartTime(_useStartTime string) error {
	r._useStartTime = _useStartTime
	r.Set("use_start_time", _useStartTime)
	return nil
}

// GetUseStartTime UseStartTime Getter
func (r TaobaoTbkDgVegasTljCreateAPIRequest) GetUseStartTime() string {
	return r._useStartTime
}

// SetUseEndTime is UseEndTime Setter
// 使用结束日期。如果是结束时间模式为相对时间，时间格式为1-7直接的整数, 例如，1（相对领取时间1天）； 如果是绝对时间，格式为yyyy-MM-dd，例如，2019-01-29，表示到2019-01-29 23:59:59结束
func (r *TaobaoTbkDgVegasTljCreateAPIRequest) SetUseEndTime(_useEndTime string) error {
	r._useEndTime = _useEndTime
	r.Set("use_end_time", _useEndTime)
	return nil
}

// GetUseEndTime UseEndTime Getter
func (r TaobaoTbkDgVegasTljCreateAPIRequest) GetUseEndTime() string {
	return r._useEndTime
}

// SetSendEndTime is SendEndTime Setter
// 发放截止时间
func (r *TaobaoTbkDgVegasTljCreateAPIRequest) SetSendEndTime(_sendEndTime string) error {
	r._sendEndTime = _sendEndTime
	r.Set("send_end_time", _sendEndTime)
	return nil
}

// GetSendEndTime SendEndTime Getter
func (r TaobaoTbkDgVegasTljCreateAPIRequest) GetSendEndTime() string {
	return r._sendEndTime
}

// SetSendStartTime is SendStartTime Setter
// 发放开始时间
func (r *TaobaoTbkDgVegasTljCreateAPIRequest) SetSendStartTime(_sendStartTime string) error {
	r._sendStartTime = _sendStartTime
	r.Set("send_start_time", _sendStartTime)
	return nil
}

// GetSendStartTime SendStartTime Getter
func (r TaobaoTbkDgVegasTljCreateAPIRequest) GetSendStartTime() string {
	return r._sendStartTime
}

// SetPerFace is PerFace Setter
// 单个淘礼金面额，支持两位小数，单位元
func (r *TaobaoTbkDgVegasTljCreateAPIRequest) SetPerFace(_perFace string) error {
	r._perFace = _perFace
	r.Set("per_face", _perFace)
	return nil
}

// GetPerFace PerFace Getter
func (r TaobaoTbkDgVegasTljCreateAPIRequest) GetPerFace() string {
	return r._perFace
}

// SetName is Name Setter
// 淘礼金名称，最大10个字符
func (r *TaobaoTbkDgVegasTljCreateAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoTbkDgVegasTljCreateAPIRequest) GetName() string {
	return r._name
}

// SetItemId is ItemId Setter
// 宝贝ID或营销ID
func (r *TaobaoTbkDgVegasTljCreateAPIRequest) SetItemId(_itemId string) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoTbkDgVegasTljCreateAPIRequest) GetItemId() string {
	return r._itemId
}

// SetCampaignType is CampaignType Setter
// 已下线，后续不需要填写
func (r *TaobaoTbkDgVegasTljCreateAPIRequest) SetCampaignType(_campaignType string) error {
	r._campaignType = _campaignType
	r.Set("campaign_type", _campaignType)
	return nil
}

// GetCampaignType CampaignType Getter
func (r TaobaoTbkDgVegasTljCreateAPIRequest) GetCampaignType() string {
	return r._campaignType
}

// SetSecurityLevel is SecurityLevel Setter
// 必须传入0
func (r *TaobaoTbkDgVegasTljCreateAPIRequest) SetSecurityLevel(_securityLevel int64) error {
	r._securityLevel = _securityLevel
	r.Set("security_level", _securityLevel)
	return nil
}

// GetSecurityLevel SecurityLevel Getter
func (r TaobaoTbkDgVegasTljCreateAPIRequest) GetSecurityLevel() int64 {
	return r._securityLevel
}

// SetUseEndTimeMode is UseEndTimeMode Setter
// 结束日期的模式,1:相对时间，2:绝对时间
func (r *TaobaoTbkDgVegasTljCreateAPIRequest) SetUseEndTimeMode(_useEndTimeMode int64) error {
	r._useEndTimeMode = _useEndTimeMode
	r.Set("use_end_time_mode", _useEndTimeMode)
	return nil
}

// GetUseEndTimeMode UseEndTimeMode Getter
func (r TaobaoTbkDgVegasTljCreateAPIRequest) GetUseEndTimeMode() int64 {
	return r._useEndTimeMode
}

// SetUserTotalWinNumLimit is UserTotalWinNumLimit Setter
// 单用户累计中奖次数上限
func (r *TaobaoTbkDgVegasTljCreateAPIRequest) SetUserTotalWinNumLimit(_userTotalWinNumLimit int64) error {
	r._userTotalWinNumLimit = _userTotalWinNumLimit
	r.Set("user_total_win_num_limit", _userTotalWinNumLimit)
	return nil
}

// GetUserTotalWinNumLimit UserTotalWinNumLimit Getter
func (r TaobaoTbkDgVegasTljCreateAPIRequest) GetUserTotalWinNumLimit() int64 {
	return r._userTotalWinNumLimit
}

// SetTotalNum is TotalNum Setter
// 淘礼金总个数
func (r *TaobaoTbkDgVegasTljCreateAPIRequest) SetTotalNum(_totalNum int64) error {
	r._totalNum = _totalNum
	r.Set("total_num", _totalNum)
	return nil
}

// GetTotalNum TotalNum Getter
func (r TaobaoTbkDgVegasTljCreateAPIRequest) GetTotalNum() int64 {
	return r._totalNum
}

// SetAdzoneId is AdzoneId Setter
// 妈妈广告位Id
func (r *TaobaoTbkDgVegasTljCreateAPIRequest) SetAdzoneId(_adzoneId int64) error {
	r._adzoneId = _adzoneId
	r.Set("adzone_id", _adzoneId)
	return nil
}

// GetAdzoneId AdzoneId Getter
func (r TaobaoTbkDgVegasTljCreateAPIRequest) GetAdzoneId() int64 {
	return r._adzoneId
}

// SetSecuritySwitch is SecuritySwitch Setter
// 必须设置为true，默认开启安全
func (r *TaobaoTbkDgVegasTljCreateAPIRequest) SetSecuritySwitch(_securitySwitch bool) error {
	r._securitySwitch = _securitySwitch
	r.Set("security_switch", _securitySwitch)
	return nil
}

// GetSecuritySwitch SecuritySwitch Getter
func (r TaobaoTbkDgVegasTljCreateAPIRequest) GetSecuritySwitch() bool {
	return r._securitySwitch
}
