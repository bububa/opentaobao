package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmMemberinfoUpdatePrivyAPIRequest 编辑会员资料 API请求
// taobao.crm.memberinfo.update.privy
//
// 编辑会员的基本资料，接口返回会员信息修改是否成功
type TaobaoCrmMemberinfoUpdatePrivyAPIRequest struct {
	model.Params
	// 用于描述会员的状态，normal表示正常，blacklist表示黑名单(不享受会员折扣).
	_status string
	// 北京=1,天津=2,河北省=3,山西省=4,内蒙古自治区=5,辽宁省=6,吉林省=7,黑龙江省=8,上海=9,江苏省=10,浙江省=11,安徽省=12,福建省=13,江西省=14,山东省=15,河南省=16,湖北省=17,湖南省=18, 广东省=19,广西壮族自治区=20,海南省=21,重庆=22,四川省=23,贵州省=24,云南省=25,西藏自治区=26,陕西省=27,甘肃省=28,青海省=29,宁夏回族自治区=30,新疆维吾尔自治区=31,台湾省=32,香港特别行政区=33,澳门特别行政区=34,海外=35，约定36为清除Province设置.请注意:从2014.4.15之后,省市将采用地区标准码,请通过物流API taobao.areas.get接口获取,参考:http://api.taobao.com/apidoc/api.htm?spm=0.0.0.0.nOOF9g&path=cid:7-apiId:59.API对于老的省市代码兼容会逐步下线.
	_province string
	// 城市.请注意:从2014.4.15之后,省市将采用地区标准码,请通过物流API taobao.areas.get接口获取,参考:http://api.taobao.com/apidoc/api.htm?spm=0.0.0.0.nOOF9g&path=cid:7-apiId:59.API对于老的省市代码兼容会逐步下线
	_city string
	// 分组的id集合字符串
	_groupIds string
	// ouid
	_ouid string
	// 会员等级，1：普通客户，2：高级会员，3：高级会员 ，4：至尊vip只有正常会员才给予升级，对于status blacklist的会员升级无效
	_grade int64
	// 交易笔数
	_tradeCount int64
	// 交易金额，单位：分
	_tradeAmount int64
	// 交易关闭次数
	_closeTradeCount int64
	// 交易关闭金额，单位：分
	_closeTradeAmount int64
	// 宝贝件数
	_itemNum int64
}

// NewTaobaoCrmMemberinfoUpdatePrivyRequest 初始化TaobaoCrmMemberinfoUpdatePrivyAPIRequest对象
func NewTaobaoCrmMemberinfoUpdatePrivyRequest() *TaobaoCrmMemberinfoUpdatePrivyAPIRequest {
	return &TaobaoCrmMemberinfoUpdatePrivyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCrmMemberinfoUpdatePrivyAPIRequest) GetApiMethodName() string {
	return "taobao.crm.memberinfo.update.privy"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCrmMemberinfoUpdatePrivyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoCrmMemberinfoUpdatePrivyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStatus is Status Setter
// 用于描述会员的状态，normal表示正常，blacklist表示黑名单(不享受会员折扣).
func (r *TaobaoCrmMemberinfoUpdatePrivyAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoCrmMemberinfoUpdatePrivyAPIRequest) GetStatus() string {
	return r._status
}

// SetProvince is Province Setter
// 北京=1,天津=2,河北省=3,山西省=4,内蒙古自治区=5,辽宁省=6,吉林省=7,黑龙江省=8,上海=9,江苏省=10,浙江省=11,安徽省=12,福建省=13,江西省=14,山东省=15,河南省=16,湖北省=17,湖南省=18, 广东省=19,广西壮族自治区=20,海南省=21,重庆=22,四川省=23,贵州省=24,云南省=25,西藏自治区=26,陕西省=27,甘肃省=28,青海省=29,宁夏回族自治区=30,新疆维吾尔自治区=31,台湾省=32,香港特别行政区=33,澳门特别行政区=34,海外=35，约定36为清除Province设置.请注意:从2014.4.15之后,省市将采用地区标准码,请通过物流API taobao.areas.get接口获取,参考:http://api.taobao.com/apidoc/api.htm?spm=0.0.0.0.nOOF9g&amp;path=cid:7-apiId:59.API对于老的省市代码兼容会逐步下线.
func (r *TaobaoCrmMemberinfoUpdatePrivyAPIRequest) SetProvince(_province string) error {
	r._province = _province
	r.Set("province", _province)
	return nil
}

// GetProvince Province Getter
func (r TaobaoCrmMemberinfoUpdatePrivyAPIRequest) GetProvince() string {
	return r._province
}

// SetCity is City Setter
// 城市.请注意:从2014.4.15之后,省市将采用地区标准码,请通过物流API taobao.areas.get接口获取,参考:http://api.taobao.com/apidoc/api.htm?spm=0.0.0.0.nOOF9g&amp;path=cid:7-apiId:59.API对于老的省市代码兼容会逐步下线
func (r *TaobaoCrmMemberinfoUpdatePrivyAPIRequest) SetCity(_city string) error {
	r._city = _city
	r.Set("city", _city)
	return nil
}

// GetCity City Getter
func (r TaobaoCrmMemberinfoUpdatePrivyAPIRequest) GetCity() string {
	return r._city
}

// SetGroupIds is GroupIds Setter
// 分组的id集合字符串
func (r *TaobaoCrmMemberinfoUpdatePrivyAPIRequest) SetGroupIds(_groupIds string) error {
	r._groupIds = _groupIds
	r.Set("group_ids", _groupIds)
	return nil
}

// GetGroupIds GroupIds Getter
func (r TaobaoCrmMemberinfoUpdatePrivyAPIRequest) GetGroupIds() string {
	return r._groupIds
}

// SetOuid is Ouid Setter
// ouid
func (r *TaobaoCrmMemberinfoUpdatePrivyAPIRequest) SetOuid(_ouid string) error {
	r._ouid = _ouid
	r.Set("ouid", _ouid)
	return nil
}

// GetOuid Ouid Getter
func (r TaobaoCrmMemberinfoUpdatePrivyAPIRequest) GetOuid() string {
	return r._ouid
}

// SetGrade is Grade Setter
// 会员等级，1：普通客户，2：高级会员，3：高级会员 ，4：至尊vip只有正常会员才给予升级，对于status blacklist的会员升级无效
func (r *TaobaoCrmMemberinfoUpdatePrivyAPIRequest) SetGrade(_grade int64) error {
	r._grade = _grade
	r.Set("grade", _grade)
	return nil
}

// GetGrade Grade Getter
func (r TaobaoCrmMemberinfoUpdatePrivyAPIRequest) GetGrade() int64 {
	return r._grade
}

// SetTradeCount is TradeCount Setter
// 交易笔数
func (r *TaobaoCrmMemberinfoUpdatePrivyAPIRequest) SetTradeCount(_tradeCount int64) error {
	r._tradeCount = _tradeCount
	r.Set("trade_count", _tradeCount)
	return nil
}

// GetTradeCount TradeCount Getter
func (r TaobaoCrmMemberinfoUpdatePrivyAPIRequest) GetTradeCount() int64 {
	return r._tradeCount
}

// SetTradeAmount is TradeAmount Setter
// 交易金额，单位：分
func (r *TaobaoCrmMemberinfoUpdatePrivyAPIRequest) SetTradeAmount(_tradeAmount int64) error {
	r._tradeAmount = _tradeAmount
	r.Set("trade_amount", _tradeAmount)
	return nil
}

// GetTradeAmount TradeAmount Getter
func (r TaobaoCrmMemberinfoUpdatePrivyAPIRequest) GetTradeAmount() int64 {
	return r._tradeAmount
}

// SetCloseTradeCount is CloseTradeCount Setter
// 交易关闭次数
func (r *TaobaoCrmMemberinfoUpdatePrivyAPIRequest) SetCloseTradeCount(_closeTradeCount int64) error {
	r._closeTradeCount = _closeTradeCount
	r.Set("close_trade_count", _closeTradeCount)
	return nil
}

// GetCloseTradeCount CloseTradeCount Getter
func (r TaobaoCrmMemberinfoUpdatePrivyAPIRequest) GetCloseTradeCount() int64 {
	return r._closeTradeCount
}

// SetCloseTradeAmount is CloseTradeAmount Setter
// 交易关闭金额，单位：分
func (r *TaobaoCrmMemberinfoUpdatePrivyAPIRequest) SetCloseTradeAmount(_closeTradeAmount int64) error {
	r._closeTradeAmount = _closeTradeAmount
	r.Set("close_trade_amount", _closeTradeAmount)
	return nil
}

// GetCloseTradeAmount CloseTradeAmount Getter
func (r TaobaoCrmMemberinfoUpdatePrivyAPIRequest) GetCloseTradeAmount() int64 {
	return r._closeTradeAmount
}

// SetItemNum is ItemNum Setter
// 宝贝件数
func (r *TaobaoCrmMemberinfoUpdatePrivyAPIRequest) SetItemNum(_itemNum int64) error {
	r._itemNum = _itemNum
	r.Set("item_num", _itemNum)
	return nil
}

// GetItemNum ItemNum Getter
func (r TaobaoCrmMemberinfoUpdatePrivyAPIRequest) GetItemNum() int64 {
	return r._itemNum
}
