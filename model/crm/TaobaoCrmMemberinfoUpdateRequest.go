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
    buyerNick   string
    // 用于描述会员的状态，normal表示正常，blacklist表示黑名单(不享受会员折扣).
    status   string
    // 会员等级，1：普通客户，2：高级会员，3：高级会员 ，4：至尊vip只有正常会员才给予升级，对于status blacklist的会员升级无效
    grade   int64
    // 北京=1,天津=2,河北省=3,山西省=4,内蒙古自治区=5,辽宁省=6,吉林省=7,黑龙江省=8,上海=9,江苏省=10,浙江省=11,安徽省=12,福建省=13,江西省=14,山东省=15,河南省=16,湖北省=17,湖南省=18, 广东省=19,广西壮族自治区=20,海南省=21,重庆=22,四川省=23,贵州省=24,云南省=25,西藏自治区=26,陕西省=27,甘肃省=28,青海省=29,宁夏回族自治区=30,新疆维吾尔自治区=31,台湾省=32,香港特别行政区=33,澳门特别行政区=34,海外=35，约定36为清除Province设置.请注意:从2014.4.15之后,省市将采用地区标准码,请通过物流API taobao.areas.get接口获取,参考:http://api.taobao.com/apidoc/api.htm?spm=0.0.0.0.nOOF9g&path=cid:7-apiId:59.API对于老的省市代码兼容会逐步下线.
    province   string
    // 城市.请注意:从2014.4.15之后,省市将采用地区标准码,请通过物流API taobao.areas.get接口获取,参考:http://api.taobao.com/apidoc/api.htm?spm=0.0.0.0.nOOF9g&path=cid:7-apiId:59.API对于老的省市代码兼容会逐步下线
    city   string
    // 交易笔数
    tradeCount   int64
    // 交易金额，单位：分
    tradeAmount   int64
    // 交易关闭次数
    closeTradeCount   int64
    // 交易关闭金额，单位：分
    closeTradeAmount   int64
    // 分组的id集合字符串
    groupIds   string
    // 宝贝件数
    itemNum   int64
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
func (r *TaobaoCrmMemberinfoUpdateRequest) SetBuyerNick(buyerNick string) error {
    r.buyerNick = buyerNick
    r.Set("buyer_nick", buyerNick)
    return nil
}

// BuyerNick Getter
func (r TaobaoCrmMemberinfoUpdateRequest) GetBuyerNick() string {
    return r.buyerNick
}
// Status Setter
// 用于描述会员的状态，normal表示正常，blacklist表示黑名单(不享受会员折扣).
func (r *TaobaoCrmMemberinfoUpdateRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r TaobaoCrmMemberinfoUpdateRequest) GetStatus() string {
    return r.status
}
// Grade Setter
// 会员等级，1：普通客户，2：高级会员，3：高级会员 ，4：至尊vip只有正常会员才给予升级，对于status blacklist的会员升级无效
func (r *TaobaoCrmMemberinfoUpdateRequest) SetGrade(grade int64) error {
    r.grade = grade
    r.Set("grade", grade)
    return nil
}

// Grade Getter
func (r TaobaoCrmMemberinfoUpdateRequest) GetGrade() int64 {
    return r.grade
}
// Province Setter
// 北京=1,天津=2,河北省=3,山西省=4,内蒙古自治区=5,辽宁省=6,吉林省=7,黑龙江省=8,上海=9,江苏省=10,浙江省=11,安徽省=12,福建省=13,江西省=14,山东省=15,河南省=16,湖北省=17,湖南省=18, 广东省=19,广西壮族自治区=20,海南省=21,重庆=22,四川省=23,贵州省=24,云南省=25,西藏自治区=26,陕西省=27,甘肃省=28,青海省=29,宁夏回族自治区=30,新疆维吾尔自治区=31,台湾省=32,香港特别行政区=33,澳门特别行政区=34,海外=35，约定36为清除Province设置.请注意:从2014.4.15之后,省市将采用地区标准码,请通过物流API taobao.areas.get接口获取,参考:http://api.taobao.com/apidoc/api.htm?spm=0.0.0.0.nOOF9g&path=cid:7-apiId:59.API对于老的省市代码兼容会逐步下线.
func (r *TaobaoCrmMemberinfoUpdateRequest) SetProvince(province string) error {
    r.province = province
    r.Set("province", province)
    return nil
}

// Province Getter
func (r TaobaoCrmMemberinfoUpdateRequest) GetProvince() string {
    return r.province
}
// City Setter
// 城市.请注意:从2014.4.15之后,省市将采用地区标准码,请通过物流API taobao.areas.get接口获取,参考:http://api.taobao.com/apidoc/api.htm?spm=0.0.0.0.nOOF9g&path=cid:7-apiId:59.API对于老的省市代码兼容会逐步下线
func (r *TaobaoCrmMemberinfoUpdateRequest) SetCity(city string) error {
    r.city = city
    r.Set("city", city)
    return nil
}

// City Getter
func (r TaobaoCrmMemberinfoUpdateRequest) GetCity() string {
    return r.city
}
// TradeCount Setter
// 交易笔数
func (r *TaobaoCrmMemberinfoUpdateRequest) SetTradeCount(tradeCount int64) error {
    r.tradeCount = tradeCount
    r.Set("trade_count", tradeCount)
    return nil
}

// TradeCount Getter
func (r TaobaoCrmMemberinfoUpdateRequest) GetTradeCount() int64 {
    return r.tradeCount
}
// TradeAmount Setter
// 交易金额，单位：分
func (r *TaobaoCrmMemberinfoUpdateRequest) SetTradeAmount(tradeAmount int64) error {
    r.tradeAmount = tradeAmount
    r.Set("trade_amount", tradeAmount)
    return nil
}

// TradeAmount Getter
func (r TaobaoCrmMemberinfoUpdateRequest) GetTradeAmount() int64 {
    return r.tradeAmount
}
// CloseTradeCount Setter
// 交易关闭次数
func (r *TaobaoCrmMemberinfoUpdateRequest) SetCloseTradeCount(closeTradeCount int64) error {
    r.closeTradeCount = closeTradeCount
    r.Set("close_trade_count", closeTradeCount)
    return nil
}

// CloseTradeCount Getter
func (r TaobaoCrmMemberinfoUpdateRequest) GetCloseTradeCount() int64 {
    return r.closeTradeCount
}
// CloseTradeAmount Setter
// 交易关闭金额，单位：分
func (r *TaobaoCrmMemberinfoUpdateRequest) SetCloseTradeAmount(closeTradeAmount int64) error {
    r.closeTradeAmount = closeTradeAmount
    r.Set("close_trade_amount", closeTradeAmount)
    return nil
}

// CloseTradeAmount Getter
func (r TaobaoCrmMemberinfoUpdateRequest) GetCloseTradeAmount() int64 {
    return r.closeTradeAmount
}
// GroupIds Setter
// 分组的id集合字符串
func (r *TaobaoCrmMemberinfoUpdateRequest) SetGroupIds(groupIds string) error {
    r.groupIds = groupIds
    r.Set("group_ids", groupIds)
    return nil
}

// GroupIds Getter
func (r TaobaoCrmMemberinfoUpdateRequest) GetGroupIds() string {
    return r.groupIds
}
// ItemNum Setter
// 宝贝件数
func (r *TaobaoCrmMemberinfoUpdateRequest) SetItemNum(itemNum int64) error {
    r.itemNum = itemNum
    r.Set("item_num", itemNum)
    return nil
}

// ItemNum Getter
func (r TaobaoCrmMemberinfoUpdateRequest) GetItemNum() int64 {
    return r.itemNum
}
