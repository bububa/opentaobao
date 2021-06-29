package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家地址库新增接口 API请求
taobao.logistics.address.add

通过此接口新增卖家地址库,卖家最多可添加5条地址库,新增第一条卖家地址，将会自动设为默认地址库
*/
type TaobaoLogisticsAddressAddRequest struct {
    model.Params
    // 联系人姓名 <font color='red'>长度不可超过20个字节</font>
    contactName   string
    // 所在省
    province   string
    // 所在市
    city   string
    // 区、县<br/><br><font color='red'>如果所在地区是海外的可以为空，否则为必参</font>
    country   string
    // 详细街道地址，不需要重复填写省/市/区
    addr   string
    // 地区邮政编码<br/><br><font color='red'>如果所在地区是海外的可以为空，否则为必参</font>
    zipCode   string
    // 电话号码,手机与电话必需有一个
    phone   string
    // 手机号码，手机与电话必需有一个<br/><br><font color='red'>手机号码不能超过20位</font>
    mobilePhone   string
    // 公司名称,<br><font color="red">公司名称长能不能超过20字节</font>
    sellerCompany   string
    // 备注,<br><font color='red'>备注不能超过256字节</font>
    memo   string
    // 默认取货地址。<br><br/><font color='red'>选择此项(true)，将当前地址设为默认取货地址，撤消原默认取货地址</font>
    getDef   bool
    // 默认退货地址。<br><br/><font color='red'>选择此项(true)，将当前地址设为默认退货地址，撤消原默认退货地址</font>
    cancelDef   bool
}

// 初始化TaobaoLogisticsAddressAddRequest对象
func NewTaobaoLogisticsAddressAddRequest() *TaobaoLogisticsAddressAddRequest{
    return &TaobaoLogisticsAddressAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsAddressAddRequest) GetApiMethodName() string {
    return "taobao.logistics.address.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsAddressAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ContactName Setter
// 联系人姓名 <font color='red'>长度不可超过20个字节</font>
func (r *TaobaoLogisticsAddressAddRequest) SetContactName(contactName string) error {
    r.contactName = contactName
    r.Set("contact_name", contactName)
    return nil
}

// ContactName Getter
func (r TaobaoLogisticsAddressAddRequest) GetContactName() string {
    return r.contactName
}
// Province Setter
// 所在省
func (r *TaobaoLogisticsAddressAddRequest) SetProvince(province string) error {
    r.province = province
    r.Set("province", province)
    return nil
}

// Province Getter
func (r TaobaoLogisticsAddressAddRequest) GetProvince() string {
    return r.province
}
// City Setter
// 所在市
func (r *TaobaoLogisticsAddressAddRequest) SetCity(city string) error {
    r.city = city
    r.Set("city", city)
    return nil
}

// City Getter
func (r TaobaoLogisticsAddressAddRequest) GetCity() string {
    return r.city
}
// Country Setter
// 区、县<br/><br><font color='red'>如果所在地区是海外的可以为空，否则为必参</font>
func (r *TaobaoLogisticsAddressAddRequest) SetCountry(country string) error {
    r.country = country
    r.Set("country", country)
    return nil
}

// Country Getter
func (r TaobaoLogisticsAddressAddRequest) GetCountry() string {
    return r.country
}
// Addr Setter
// 详细街道地址，不需要重复填写省/市/区
func (r *TaobaoLogisticsAddressAddRequest) SetAddr(addr string) error {
    r.addr = addr
    r.Set("addr", addr)
    return nil
}

// Addr Getter
func (r TaobaoLogisticsAddressAddRequest) GetAddr() string {
    return r.addr
}
// ZipCode Setter
// 地区邮政编码<br/><br><font color='red'>如果所在地区是海外的可以为空，否则为必参</font>
func (r *TaobaoLogisticsAddressAddRequest) SetZipCode(zipCode string) error {
    r.zipCode = zipCode
    r.Set("zip_code", zipCode)
    return nil
}

// ZipCode Getter
func (r TaobaoLogisticsAddressAddRequest) GetZipCode() string {
    return r.zipCode
}
// Phone Setter
// 电话号码,手机与电话必需有一个
func (r *TaobaoLogisticsAddressAddRequest) SetPhone(phone string) error {
    r.phone = phone
    r.Set("phone", phone)
    return nil
}

// Phone Getter
func (r TaobaoLogisticsAddressAddRequest) GetPhone() string {
    return r.phone
}
// MobilePhone Setter
// 手机号码，手机与电话必需有一个<br/><br><font color='red'>手机号码不能超过20位</font>
func (r *TaobaoLogisticsAddressAddRequest) SetMobilePhone(mobilePhone string) error {
    r.mobilePhone = mobilePhone
    r.Set("mobile_phone", mobilePhone)
    return nil
}

// MobilePhone Getter
func (r TaobaoLogisticsAddressAddRequest) GetMobilePhone() string {
    return r.mobilePhone
}
// SellerCompany Setter
// 公司名称,<br><font color="red">公司名称长能不能超过20字节</font>
func (r *TaobaoLogisticsAddressAddRequest) SetSellerCompany(sellerCompany string) error {
    r.sellerCompany = sellerCompany
    r.Set("seller_company", sellerCompany)
    return nil
}

// SellerCompany Getter
func (r TaobaoLogisticsAddressAddRequest) GetSellerCompany() string {
    return r.sellerCompany
}
// Memo Setter
// 备注,<br><font color='red'>备注不能超过256字节</font>
func (r *TaobaoLogisticsAddressAddRequest) SetMemo(memo string) error {
    r.memo = memo
    r.Set("memo", memo)
    return nil
}

// Memo Getter
func (r TaobaoLogisticsAddressAddRequest) GetMemo() string {
    return r.memo
}
// GetDef Setter
// 默认取货地址。<br><br/><font color='red'>选择此项(true)，将当前地址设为默认取货地址，撤消原默认取货地址</font>
func (r *TaobaoLogisticsAddressAddRequest) SetGetDef(getDef bool) error {
    r.getDef = getDef
    r.Set("get_def", getDef)
    return nil
}

// GetDef Getter
func (r TaobaoLogisticsAddressAddRequest) GetGetDef() bool {
    return r.getDef
}
// CancelDef Setter
// 默认退货地址。<br><br/><font color='red'>选择此项(true)，将当前地址设为默认退货地址，撤消原默认退货地址</font>
func (r *TaobaoLogisticsAddressAddRequest) SetCancelDef(cancelDef bool) error {
    r.cancelDef = cancelDef
    r.Set("cancel_def", cancelDef)
    return nil
}

// CancelDef Getter
func (r TaobaoLogisticsAddressAddRequest) GetCancelDef() bool {
    return r.cancelDef
}
