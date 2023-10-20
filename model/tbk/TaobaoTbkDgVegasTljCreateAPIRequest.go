package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkdgvegastljcreateAPIRequest 淘宝客-推广者-淘礼金创建 API请求
// taobao.tbk.dg.vegas.tlj.create
//
// 创建淘礼金
type TaobaotbkdgvegastljcreateAPIRequest struct {
	model.Params
	// 使用开始日期。相对时间，无需填写，以用户领取时间作为使用开始时间。绝对时间，格式 yyyy-MM-dd，例如，2019-01-29，表示从2019-01-29 00:00:00 开始
	_usestarttime string
	// 使用结束日期。如果是结束时间模式为相对时间，时间格式为1-7直接的整数, 例如，1（相对领取时间1天）； 如果是绝对时间，格式为yyyy-MM-dd，例如，2019-01-29，表示到2019-01-29 23:59:59结束
	_useendtime string
	// 发放截止时间
	_sendendtime string
	// 发放开始时间
	_sendstarttime string
	// 单个淘礼金面额，支持两位小数，单位元
	_perface string
	// 淘礼金名称，最大10个字符
	_name string
	// 宝贝ID或营销ID
	_itemid string
	// 已下线，后续不需要填写
	_campaigntype string
	// 必须传入0
	_securitylevel int64
	// 结束日期的模式,1:相对时间，2:绝对时间
	_useendtimemode int64
	// 单用户累计中奖次数上限
	_usertotalwinnumlimit int64
	// 淘礼金总个数
	_totalnum int64
	// 妈妈广告位Id
	_adzoneid int64
	// 必须设置为true，默认开启安全
	_securityswitch bool
}

// NewTaobaotbkdgvegastljcreateRequest 初始化TaobaotbkdgvegastljcreateAPIRequest对象
func NewTaobaotbkdgvegastljcreateRequest() *TaobaotbkdgvegastljcreateAPIRequest {
	return &TaobaotbkdgvegastljcreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkdgvegastljcreateAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.dg.vegas.tlj.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkdgvegastljcreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkdgvegastljcreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUsestarttime is Usestarttime Setter
// 使用开始日期。相对时间，无需填写，以用户领取时间作为使用开始时间。绝对时间，格式 yyyy-MM-dd，例如，2019-01-29，表示从2019-01-29 00:00:00 开始
func (r *TaobaotbkdgvegastljcreateAPIRequest) SetUsestarttime(_usestarttime string) error {
	r._usestarttime = _usestarttime
	r.Set("use_start_time", _usestarttime)
	return nil
}

// GetUsestarttime Usestarttime Getter
func (r TaobaotbkdgvegastljcreateAPIRequest) GetUsestarttime() string {
	return r._usestarttime
}

// SetUseendtime is Useendtime Setter
// 使用结束日期。如果是结束时间模式为相对时间，时间格式为1-7直接的整数, 例如，1（相对领取时间1天）； 如果是绝对时间，格式为yyyy-MM-dd，例如，2019-01-29，表示到2019-01-29 23:59:59结束
func (r *TaobaotbkdgvegastljcreateAPIRequest) SetUseendtime(_useendtime string) error {
	r._useendtime = _useendtime
	r.Set("use_end_time", _useendtime)
	return nil
}

// GetUseendtime Useendtime Getter
func (r TaobaotbkdgvegastljcreateAPIRequest) GetUseendtime() string {
	return r._useendtime
}

// SetSendendtime is Sendendtime Setter
// 发放截止时间
func (r *TaobaotbkdgvegastljcreateAPIRequest) SetSendendtime(_sendendtime string) error {
	r._sendendtime = _sendendtime
	r.Set("send_end_time", _sendendtime)
	return nil
}

// GetSendendtime Sendendtime Getter
func (r TaobaotbkdgvegastljcreateAPIRequest) GetSendendtime() string {
	return r._sendendtime
}

// SetSendstarttime is Sendstarttime Setter
// 发放开始时间
func (r *TaobaotbkdgvegastljcreateAPIRequest) SetSendstarttime(_sendstarttime string) error {
	r._sendstarttime = _sendstarttime
	r.Set("send_start_time", _sendstarttime)
	return nil
}

// GetSendstarttime Sendstarttime Getter
func (r TaobaotbkdgvegastljcreateAPIRequest) GetSendstarttime() string {
	return r._sendstarttime
}

// SetPerface is Perface Setter
// 单个淘礼金面额，支持两位小数，单位元
func (r *TaobaotbkdgvegastljcreateAPIRequest) SetPerface(_perface string) error {
	r._perface = _perface
	r.Set("per_face", _perface)
	return nil
}

// GetPerface Perface Getter
func (r TaobaotbkdgvegastljcreateAPIRequest) GetPerface() string {
	return r._perface
}

// SetName is Name Setter
// 淘礼金名称，最大10个字符
func (r *TaobaotbkdgvegastljcreateAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaotbkdgvegastljcreateAPIRequest) GetName() string {
	return r._name
}

// SetItemid is Itemid Setter
// 宝贝ID或营销ID
func (r *TaobaotbkdgvegastljcreateAPIRequest) SetItemid(_itemid string) error {
	r._itemid = _itemid
	r.Set("item_id", _itemid)
	return nil
}

// GetItemid Itemid Getter
func (r TaobaotbkdgvegastljcreateAPIRequest) GetItemid() string {
	return r._itemid
}

// SetCampaigntype is Campaigntype Setter
// 已下线，后续不需要填写
func (r *TaobaotbkdgvegastljcreateAPIRequest) SetCampaigntype(_campaigntype string) error {
	r._campaigntype = _campaigntype
	r.Set("campaign_type", _campaigntype)
	return nil
}

// GetCampaigntype Campaigntype Getter
func (r TaobaotbkdgvegastljcreateAPIRequest) GetCampaigntype() string {
	return r._campaigntype
}

// SetSecuritylevel is Securitylevel Setter
// 必须传入0
func (r *TaobaotbkdgvegastljcreateAPIRequest) SetSecuritylevel(_securitylevel int64) error {
	r._securitylevel = _securitylevel
	r.Set("security_level", _securitylevel)
	return nil
}

// GetSecuritylevel Securitylevel Getter
func (r TaobaotbkdgvegastljcreateAPIRequest) GetSecuritylevel() int64 {
	return r._securitylevel
}

// SetUseendtimemode is Useendtimemode Setter
// 结束日期的模式,1:相对时间，2:绝对时间
func (r *TaobaotbkdgvegastljcreateAPIRequest) SetUseendtimemode(_useendtimemode int64) error {
	r._useendtimemode = _useendtimemode
	r.Set("use_end_time_mode", _useendtimemode)
	return nil
}

// GetUseendtimemode Useendtimemode Getter
func (r TaobaotbkdgvegastljcreateAPIRequest) GetUseendtimemode() int64 {
	return r._useendtimemode
}

// SetUsertotalwinnumlimit is Usertotalwinnumlimit Setter
// 单用户累计中奖次数上限
func (r *TaobaotbkdgvegastljcreateAPIRequest) SetUsertotalwinnumlimit(_usertotalwinnumlimit int64) error {
	r._usertotalwinnumlimit = _usertotalwinnumlimit
	r.Set("user_total_win_num_limit", _usertotalwinnumlimit)
	return nil
}

// GetUsertotalwinnumlimit Usertotalwinnumlimit Getter
func (r TaobaotbkdgvegastljcreateAPIRequest) GetUsertotalwinnumlimit() int64 {
	return r._usertotalwinnumlimit
}

// SetTotalnum is Totalnum Setter
// 淘礼金总个数
func (r *TaobaotbkdgvegastljcreateAPIRequest) SetTotalnum(_totalnum int64) error {
	r._totalnum = _totalnum
	r.Set("total_num", _totalnum)
	return nil
}

// GetTotalnum Totalnum Getter
func (r TaobaotbkdgvegastljcreateAPIRequest) GetTotalnum() int64 {
	return r._totalnum
}

// SetAdzoneid is Adzoneid Setter
// 妈妈广告位Id
func (r *TaobaotbkdgvegastljcreateAPIRequest) SetAdzoneid(_adzoneid int64) error {
	r._adzoneid = _adzoneid
	r.Set("adzone_id", _adzoneid)
	return nil
}

// GetAdzoneid Adzoneid Getter
func (r TaobaotbkdgvegastljcreateAPIRequest) GetAdzoneid() int64 {
	return r._adzoneid
}

// SetSecurityswitch is Securityswitch Setter
// 必须设置为true，默认开启安全
func (r *TaobaotbkdgvegastljcreateAPIRequest) SetSecurityswitch(_securityswitch bool) error {
	r._securityswitch = _securityswitch
	r.Set("security_switch", _securityswitch)
	return nil
}

// GetSecurityswitch Securityswitch Getter
func (r TaobaotbkdgvegastljcreateAPIRequest) GetSecurityswitch() bool {
	return r._securityswitch
}
