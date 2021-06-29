package crm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
编辑会员资料 API请求
taobao.crm.memberinfo.update

编辑会员的基本资料，接口返回会员信息修改是否成功
*/
type TaobaoCrmMemberinfoUpdateRequest struct {
    model.Params
    // 买家昵称
    _buyerNick   string
    // 用于描述会员的状态，normal表示正常，blacklist表示黑名单(不享受会员折扣).
    _status   string
    // 会员等级，1：普通客户，2：高级会员，3：高级会员 ，4：至尊vip只有正常会员才给予升级，对于status blacklist的会员升级无效
    _grade   int64
    // 北京=1,天津=2,河北省=3,山西省=4,内蒙古自治区=5,辽宁省=6,吉林省=7,黑龙江省=8,上海=9,江苏省=10,浙江省=11,安徽省=12,福建省=13,江西省=14,山东省=15,河南省=16,湖北省=17,湖南省=18, 广东省=19,广西壮族自治区=20,海南省=21,重庆=22,四川省=23,贵州省=24,云南省=25,西藏自治区=26,陕西省=27,甘肃省=28,青海省=29,宁夏回族自治区=30,新疆维吾尔自治区=31,台湾省=32,香港特别行政区=33,澳门特别行政区=34,海外=35，约定36为清除Province设置.请注意:从2014.4.15之后,省市将采用地区标准码,请通过物流API taobao.areas.get接口获取,参考:http://api.taobao.com/apidoc/api.htm?spm=0.0.0.0.nOOF9g&path=cid:7-apiId:59.API对于老的省市代码兼容会逐步下线.
    _province   string
    // 城市.请注意:从2014.4.15之后,省市将采用地区标准码,请通过物流API taobao.areas.get接口获取,参考:http://api.taobao.com/apidoc/api.htm?spm=0.0.0.0.nOOF9g&path=cid:7-apiId:59.API对于老的省市代码兼容会逐步下线
    _city   string
    // 交易笔数
    _tradeCount   int64
    // 交易金额，单位：分
    _tradeAmount   int64
    // 交易关闭次数
    _closeTradeCount   int64
    // 交易关闭金额，单位：分
    _closeTradeAmount   int64
    // 分组的id集合字符串
    _groupIds   string
    // 宝贝件数
    _itemNum   int64
}

// 初始化TaobaoCrmMemberinfoUpdateRequest对象
func NewTaobaoCrmMemberinfoUpdateRequest() *TaobaoCrmMemberinfoUpdateRequest{
    return &TaobaoCrmMemberinfoUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoCrmMemberinfoUpdateRequest) GetApiMethodName() string {
    return "taobao.crm.memberinfo.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoCrmMemberinfoUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BuyerNick Setter
// 买家昵称
func (r *TaobaoCrmMemberinfoUpdateRequest) SetBuyerNick(_buyerNick string) error {
    r._buyerNick = _buyerNick
    r.Set("buyer_nick", _buyerNick)
    return nil
}

// BuyerNick Getter
func (r TaobaoCrmMemberinfoUpdateRequest) GetBuyerNick() string {
    return r._buyerNick
}
// Status Setter
// 用于描述会员的状态，normal表示正常，blacklist表示黑名单(不享受会员折扣).
func (r *TaobaoCrmMemberinfoUpdateRequest) SetStatus(_status string) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r TaobaoCrmMemberinfoUpdateRequest) GetStatus() string {
    return r._status
}
// Grade Setter
// 会员等级，1：普通客户，2：高级会员，3：高级会员 ，4：至尊vip只有正常会员才给予升级，对于status blacklist的会员升级无效
func (r *TaobaoCrmMemberinfoUpdateRequest) SetGrade(_grade int64) error {
    r._grade = _grade
    r.Set("grade", _grade)
    return nil
}

// Grade Getter
func (r TaobaoCrmMemberinfoUpdateRequest) GetGrade() int64 {
    return r._grade
}
// Province Setter
// 北京=1,天津=2,河北省=3,山西省=4,内蒙古自治区=5,辽宁省=6,吉林省=7,黑龙江省=8,上海=9,江苏省=10,浙江省=11,安徽省=12,福建省=13,江西省=14,山东省=15,河南省=16,湖北省=17,湖南省=18, 广东省=19,广西壮族自治区=20,海南省=21,重庆=22,四川省=23,贵州省=24,云南省=25,西藏自治区=26,陕西省=27,甘肃省=28,青海省=29,宁夏回族自治区=30,新疆维吾尔自治区=31,台湾省=32,香港特别行政区=33,澳门特别行政区=34,海外=35，约定36为清除Province设置.请注意:从2014.4.15之后,省市将采用地区标准码,请通过物流API taobao.areas.get接口获取,参考:http://api.taobao.com/apidoc/api.htm?spm=0.0.0.0.nOOF9g&path=cid:7-apiId:59.API对于老的省市代码兼容会逐步下线.
func (r *TaobaoCrmMemberinfoUpdateRequest) SetProvince(_province string) error {
    r._province = _province
    r.Set("province", _province)
    return nil
}

// Province Getter
func (r TaobaoCrmMemberinfoUpdateRequest) GetProvince() string {
    return r._province
}
// City Setter
// 城市.请注意:从2014.4.15之后,省市将采用地区标准码,请通过物流API taobao.areas.get接口获取,参考:http://api.taobao.com/apidoc/api.htm?spm=0.0.0.0.nOOF9g&path=cid:7-apiId:59.API对于老的省市代码兼容会逐步下线
func (r *TaobaoCrmMemberinfoUpdateRequest) SetCity(_city string) error {
    r._city = _city
    r.Set("city", _city)
    return nil
}

// City Getter
func (r TaobaoCrmMemberinfoUpdateRequest) GetCity() string {
    return r._city
}
// TradeCount Setter
// 交易笔数
func (r *TaobaoCrmMemberinfoUpdateRequest) SetTradeCount(_tradeCount int64) error {
    r._tradeCount = _tradeCount
    r.Set("trade_count", _tradeCount)
    return nil
}

// TradeCount Getter
func (r TaobaoCrmMemberinfoUpdateRequest) GetTradeCount() int64 {
    return r._tradeCount
}
// TradeAmount Setter
// 交易金额，单位：分
func (r *TaobaoCrmMemberinfoUpdateRequest) SetTradeAmount(_tradeAmount int64) error {
    r._tradeAmount = _tradeAmount
    r.Set("trade_amount", _tradeAmount)
    return nil
}

// TradeAmount Getter
func (r TaobaoCrmMemberinfoUpdateRequest) GetTradeAmount() int64 {
    return r._tradeAmount
}
// CloseTradeCount Setter
// 交易关闭次数
func (r *TaobaoCrmMemberinfoUpdateRequest) SetCloseTradeCount(_closeTradeCount int64) error {
    r._closeTradeCount = _closeTradeCount
    r.Set("close_trade_count", _closeTradeCount)
    return nil
}

// CloseTradeCount Getter
func (r TaobaoCrmMemberinfoUpdateRequest) GetCloseTradeCount() int64 {
    return r._closeTradeCount
}
// CloseTradeAmount Setter
// 交易关闭金额，单位：分
func (r *TaobaoCrmMemberinfoUpdateRequest) SetCloseTradeAmount(_closeTradeAmount int64) error {
    r._closeTradeAmount = _closeTradeAmount
    r.Set("close_trade_amount", _closeTradeAmount)
    return nil
}

// CloseTradeAmount Getter
func (r TaobaoCrmMemberinfoUpdateRequest) GetCloseTradeAmount() int64 {
    return r._closeTradeAmount
}
// GroupIds Setter
// 分组的id集合字符串
func (r *TaobaoCrmMemberinfoUpdateRequest) SetGroupIds(_groupIds string) error {
    r._groupIds = _groupIds
    r.Set("group_ids", _groupIds)
    return nil
}

// GroupIds Getter
func (r TaobaoCrmMemberinfoUpdateRequest) GetGroupIds() string {
    return r._groupIds
}
// ItemNum Setter
// 宝贝件数
func (r *TaobaoCrmMemberinfoUpdateRequest) SetItemNum(_itemNum int64) error {
    r._itemNum = _itemNum
    r.Set("item_num", _itemNum)
    return nil
}

// ItemNum Getter
func (r TaobaoCrmMemberinfoUpdateRequest) GetItemNum() int64 {
    return r._itemNum
}
