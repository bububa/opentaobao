package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建订单并发货 API请求
taobao.logistics.consign.order.createandsend

创建物流订单，并发货。
*/
type TaobaoLogisticsConsignOrderCreateandsendRequest struct {
    model.Params
    // 用户ID
    userId   int64
    // 订单来源，值选择：30
    orderSource   int64
    // 订单类型，值固定选择：30
    orderType   int64
    // 物流订单物流类型，值固定选择：2
    logisType   int64
    // 物流公司ID
    companyId   int64
    // 交易流水号，淘外订单号或者商家内部交易流水号
    tradeId   int64
    // 运单号
    mailNo   string
    // 费用承担方式 1买家承担运费 2卖家承担运费
    shipping   string
    // 发件人名称
    sName   string
    // 发件人区域ID
    sAreaId   int64
    // 发件人街道地址
    sAddress   string
    // 发件人出编
    sZipCode   string
    // 手机号码
    sMobilePhone   string
    // 电话号码
    sTelephone   string
    // 省
    sProvName   string
    // 市
    sCityName   string
    // 区
    sDistName   string
    // 收件人名称
    rName   string
    // 收件人区域ID
    rAreaId   int64
    // 收件人街道地址
    rAddress   string
    // 收件人邮编
    rZipCode   string
    // 手机号码
    rMobilePhone   string
    // 电话号码
    rTelephone   string
    // 省
    rProvName   string
    // 市
    rCityName   string
    // 区
    rDistName   string
    // 物品的json数据。
    itemJsonString   string
}

// 初始化TaobaoLogisticsConsignOrderCreateandsendRequest对象
func NewTaobaoLogisticsConsignOrderCreateandsendRequest() *TaobaoLogisticsConsignOrderCreateandsendRequest{
    return &TaobaoLogisticsConsignOrderCreateandsendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsConsignOrderCreateandsendRequest) GetApiMethodName() string {
    return "taobao.logistics.consign.order.createandsend"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsConsignOrderCreateandsendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserId Setter
// 用户ID
func (r *TaobaoLogisticsConsignOrderCreateandsendRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r TaobaoLogisticsConsignOrderCreateandsendRequest) GetUserId() int64 {
    return r.userId
}
// OrderSource Setter
// 订单来源，值选择：30
func (r *TaobaoLogisticsConsignOrderCreateandsendRequest) SetOrderSource(orderSource int64) error {
    r.orderSource = orderSource
    r.Set("order_source", orderSource)
    return nil
}

// OrderSource Getter
func (r TaobaoLogisticsConsignOrderCreateandsendRequest) GetOrderSource() int64 {
    return r.orderSource
}
// OrderType Setter
// 订单类型，值固定选择：30
func (r *TaobaoLogisticsConsignOrderCreateandsendRequest) SetOrderType(orderType int64) error {
    r.orderType = orderType
    r.Set("order_type", orderType)
    return nil
}

// OrderType Getter
func (r TaobaoLogisticsConsignOrderCreateandsendRequest) GetOrderType() int64 {
    return r.orderType
}
// LogisType Setter
// 物流订单物流类型，值固定选择：2
func (r *TaobaoLogisticsConsignOrderCreateandsendRequest) SetLogisType(logisType int64) error {
    r.logisType = logisType
    r.Set("logis_type", logisType)
    return nil
}

// LogisType Getter
func (r TaobaoLogisticsConsignOrderCreateandsendRequest) GetLogisType() int64 {
    return r.logisType
}
// CompanyId Setter
// 物流公司ID
func (r *TaobaoLogisticsConsignOrderCreateandsendRequest) SetCompanyId(companyId int64) error {
    r.companyId = companyId
    r.Set("company_id", companyId)
    return nil
}

// CompanyId Getter
func (r TaobaoLogisticsConsignOrderCreateandsendRequest) GetCompanyId() int64 {
    return r.companyId
}
// TradeId Setter
// 交易流水号，淘外订单号或者商家内部交易流水号
func (r *TaobaoLogisticsConsignOrderCreateandsendRequest) SetTradeId(tradeId int64) error {
    r.tradeId = tradeId
    r.Set("trade_id", tradeId)
    return nil
}

// TradeId Getter
func (r TaobaoLogisticsConsignOrderCreateandsendRequest) GetTradeId() int64 {
    return r.tradeId
}
// MailNo Setter
// 运单号
func (r *TaobaoLogisticsConsignOrderCreateandsendRequest) SetMailNo(mailNo string) error {
    r.mailNo = mailNo
    r.Set("mail_no", mailNo)
    return nil
}

// MailNo Getter
func (r TaobaoLogisticsConsignOrderCreateandsendRequest) GetMailNo() string {
    return r.mailNo
}
// Shipping Setter
// 费用承担方式 1买家承担运费 2卖家承担运费
func (r *TaobaoLogisticsConsignOrderCreateandsendRequest) SetShipping(shipping string) error {
    r.shipping = shipping
    r.Set("shipping", shipping)
    return nil
}

// Shipping Getter
func (r TaobaoLogisticsConsignOrderCreateandsendRequest) GetShipping() string {
    return r.shipping
}
// SName Setter
// 发件人名称
func (r *TaobaoLogisticsConsignOrderCreateandsendRequest) SetSName(sName string) error {
    r.sName = sName
    r.Set("s_name", sName)
    return nil
}

// SName Getter
func (r TaobaoLogisticsConsignOrderCreateandsendRequest) GetSName() string {
    return r.sName
}
// SAreaId Setter
// 发件人区域ID
func (r *TaobaoLogisticsConsignOrderCreateandsendRequest) SetSAreaId(sAreaId int64) error {
    r.sAreaId = sAreaId
    r.Set("s_area_id", sAreaId)
    return nil
}

// SAreaId Getter
func (r TaobaoLogisticsConsignOrderCreateandsendRequest) GetSAreaId() int64 {
    return r.sAreaId
}
// SAddress Setter
// 发件人街道地址
func (r *TaobaoLogisticsConsignOrderCreateandsendRequest) SetSAddress(sAddress string) error {
    r.sAddress = sAddress
    r.Set("s_address", sAddress)
    return nil
}

// SAddress Getter
func (r TaobaoLogisticsConsignOrderCreateandsendRequest) GetSAddress() string {
    return r.sAddress
}
// SZipCode Setter
// 发件人出编
func (r *TaobaoLogisticsConsignOrderCreateandsendRequest) SetSZipCode(sZipCode string) error {
    r.sZipCode = sZipCode
    r.Set("s_zip_code", sZipCode)
    return nil
}

// SZipCode Getter
func (r TaobaoLogisticsConsignOrderCreateandsendRequest) GetSZipCode() string {
    return r.sZipCode
}
// SMobilePhone Setter
// 手机号码
func (r *TaobaoLogisticsConsignOrderCreateandsendRequest) SetSMobilePhone(sMobilePhone string) error {
    r.sMobilePhone = sMobilePhone
    r.Set("s_mobile_phone", sMobilePhone)
    return nil
}

// SMobilePhone Getter
func (r TaobaoLogisticsConsignOrderCreateandsendRequest) GetSMobilePhone() string {
    return r.sMobilePhone
}
// STelephone Setter
// 电话号码
func (r *TaobaoLogisticsConsignOrderCreateandsendRequest) SetSTelephone(sTelephone string) error {
    r.sTelephone = sTelephone
    r.Set("s_telephone", sTelephone)
    return nil
}

// STelephone Getter
func (r TaobaoLogisticsConsignOrderCreateandsendRequest) GetSTelephone() string {
    return r.sTelephone
}
// SProvName Setter
// 省
func (r *TaobaoLogisticsConsignOrderCreateandsendRequest) SetSProvName(sProvName string) error {
    r.sProvName = sProvName
    r.Set("s_prov_name", sProvName)
    return nil
}

// SProvName Getter
func (r TaobaoLogisticsConsignOrderCreateandsendRequest) GetSProvName() string {
    return r.sProvName
}
// SCityName Setter
// 市
func (r *TaobaoLogisticsConsignOrderCreateandsendRequest) SetSCityName(sCityName string) error {
    r.sCityName = sCityName
    r.Set("s_city_name", sCityName)
    return nil
}

// SCityName Getter
func (r TaobaoLogisticsConsignOrderCreateandsendRequest) GetSCityName() string {
    return r.sCityName
}
// SDistName Setter
// 区
func (r *TaobaoLogisticsConsignOrderCreateandsendRequest) SetSDistName(sDistName string) error {
    r.sDistName = sDistName
    r.Set("s_dist_name", sDistName)
    return nil
}

// SDistName Getter
func (r TaobaoLogisticsConsignOrderCreateandsendRequest) GetSDistName() string {
    return r.sDistName
}
// RName Setter
// 收件人名称
func (r *TaobaoLogisticsConsignOrderCreateandsendRequest) SetRName(rName string) error {
    r.rName = rName
    r.Set("r_name", rName)
    return nil
}

// RName Getter
func (r TaobaoLogisticsConsignOrderCreateandsendRequest) GetRName() string {
    return r.rName
}
// RAreaId Setter
// 收件人区域ID
func (r *TaobaoLogisticsConsignOrderCreateandsendRequest) SetRAreaId(rAreaId int64) error {
    r.rAreaId = rAreaId
    r.Set("r_area_id", rAreaId)
    return nil
}

// RAreaId Getter
func (r TaobaoLogisticsConsignOrderCreateandsendRequest) GetRAreaId() int64 {
    return r.rAreaId
}
// RAddress Setter
// 收件人街道地址
func (r *TaobaoLogisticsConsignOrderCreateandsendRequest) SetRAddress(rAddress string) error {
    r.rAddress = rAddress
    r.Set("r_address", rAddress)
    return nil
}

// RAddress Getter
func (r TaobaoLogisticsConsignOrderCreateandsendRequest) GetRAddress() string {
    return r.rAddress
}
// RZipCode Setter
// 收件人邮编
func (r *TaobaoLogisticsConsignOrderCreateandsendRequest) SetRZipCode(rZipCode string) error {
    r.rZipCode = rZipCode
    r.Set("r_zip_code", rZipCode)
    return nil
}

// RZipCode Getter
func (r TaobaoLogisticsConsignOrderCreateandsendRequest) GetRZipCode() string {
    return r.rZipCode
}
// RMobilePhone Setter
// 手机号码
func (r *TaobaoLogisticsConsignOrderCreateandsendRequest) SetRMobilePhone(rMobilePhone string) error {
    r.rMobilePhone = rMobilePhone
    r.Set("r_mobile_phone", rMobilePhone)
    return nil
}

// RMobilePhone Getter
func (r TaobaoLogisticsConsignOrderCreateandsendRequest) GetRMobilePhone() string {
    return r.rMobilePhone
}
// RTelephone Setter
// 电话号码
func (r *TaobaoLogisticsConsignOrderCreateandsendRequest) SetRTelephone(rTelephone string) error {
    r.rTelephone = rTelephone
    r.Set("r_telephone", rTelephone)
    return nil
}

// RTelephone Getter
func (r TaobaoLogisticsConsignOrderCreateandsendRequest) GetRTelephone() string {
    return r.rTelephone
}
// RProvName Setter
// 省
func (r *TaobaoLogisticsConsignOrderCreateandsendRequest) SetRProvName(rProvName string) error {
    r.rProvName = rProvName
    r.Set("r_prov_name", rProvName)
    return nil
}

// RProvName Getter
func (r TaobaoLogisticsConsignOrderCreateandsendRequest) GetRProvName() string {
    return r.rProvName
}
// RCityName Setter
// 市
func (r *TaobaoLogisticsConsignOrderCreateandsendRequest) SetRCityName(rCityName string) error {
    r.rCityName = rCityName
    r.Set("r_city_name", rCityName)
    return nil
}

// RCityName Getter
func (r TaobaoLogisticsConsignOrderCreateandsendRequest) GetRCityName() string {
    return r.rCityName
}
// RDistName Setter
// 区
func (r *TaobaoLogisticsConsignOrderCreateandsendRequest) SetRDistName(rDistName string) error {
    r.rDistName = rDistName
    r.Set("r_dist_name", rDistName)
    return nil
}

// RDistName Getter
func (r TaobaoLogisticsConsignOrderCreateandsendRequest) GetRDistName() string {
    return r.rDistName
}
// ItemJsonString Setter
// 物品的json数据。
func (r *TaobaoLogisticsConsignOrderCreateandsendRequest) SetItemJsonString(itemJsonString string) error {
    r.itemJsonString = itemJsonString
    r.Set("item_json_string", itemJsonString)
    return nil
}

// ItemJsonString Getter
func (r TaobaoLogisticsConsignOrderCreateandsendRequest) GetItemJsonString() string {
    return r.itemJsonString
}
