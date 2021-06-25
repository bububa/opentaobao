package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家地址库新增接口 APIRequest
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

func NewTaobaoLogisticsAddressAddRequest() *TaobaoLogisticsAddressAddRequest{
    return &TaobaoLogisticsAddressAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoLogisticsAddressAddRequest) GetApiMethodName() string {
    return "taobao.logistics.address.add"
}

func (r TaobaoLogisticsAddressAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoLogisticsAddressAddRequest) SetContactName(contactName string) error {
    r.contactName = contactName
    r.Set("contact_name", contactName)
    return nil
}

func (r TaobaoLogisticsAddressAddRequest) GetContactName() string {
    return r.contactName
}

func (r *TaobaoLogisticsAddressAddRequest) SetProvince(province string) error {
    r.province = province
    r.Set("province", province)
    return nil
}

func (r TaobaoLogisticsAddressAddRequest) GetProvince() string {
    return r.province
}

func (r *TaobaoLogisticsAddressAddRequest) SetCity(city string) error {
    r.city = city
    r.Set("city", city)
    return nil
}

func (r TaobaoLogisticsAddressAddRequest) GetCity() string {
    return r.city
}

func (r *TaobaoLogisticsAddressAddRequest) SetCountry(country string) error {
    r.country = country
    r.Set("country", country)
    return nil
}

func (r TaobaoLogisticsAddressAddRequest) GetCountry() string {
    return r.country
}

func (r *TaobaoLogisticsAddressAddRequest) SetAddr(addr string) error {
    r.addr = addr
    r.Set("addr", addr)
    return nil
}

func (r TaobaoLogisticsAddressAddRequest) GetAddr() string {
    return r.addr
}

func (r *TaobaoLogisticsAddressAddRequest) SetZipCode(zipCode string) error {
    r.zipCode = zipCode
    r.Set("zip_code", zipCode)
    return nil
}

func (r TaobaoLogisticsAddressAddRequest) GetZipCode() string {
    return r.zipCode
}

func (r *TaobaoLogisticsAddressAddRequest) SetPhone(phone string) error {
    r.phone = phone
    r.Set("phone", phone)
    return nil
}

func (r TaobaoLogisticsAddressAddRequest) GetPhone() string {
    return r.phone
}

func (r *TaobaoLogisticsAddressAddRequest) SetMobilePhone(mobilePhone string) error {
    r.mobilePhone = mobilePhone
    r.Set("mobile_phone", mobilePhone)
    return nil
}

func (r TaobaoLogisticsAddressAddRequest) GetMobilePhone() string {
    return r.mobilePhone
}

func (r *TaobaoLogisticsAddressAddRequest) SetSellerCompany(sellerCompany string) error {
    r.sellerCompany = sellerCompany
    r.Set("seller_company", sellerCompany)
    return nil
}

func (r TaobaoLogisticsAddressAddRequest) GetSellerCompany() string {
    return r.sellerCompany
}

func (r *TaobaoLogisticsAddressAddRequest) SetMemo(memo string) error {
    r.memo = memo
    r.Set("memo", memo)
    return nil
}

func (r TaobaoLogisticsAddressAddRequest) GetMemo() string {
    return r.memo
}

func (r *TaobaoLogisticsAddressAddRequest) SetGetDef(getDef bool) error {
    r.getDef = getDef
    r.Set("get_def", getDef)
    return nil
}

func (r TaobaoLogisticsAddressAddRequest) GetGetDef() bool {
    return r.getDef
}

func (r *TaobaoLogisticsAddressAddRequest) SetCancelDef(cancelDef bool) error {
    r.cancelDef = cancelDef
    r.Set("cancel_def", cancelDef)
    return nil
}

func (r TaobaoLogisticsAddressAddRequest) GetCancelDef() bool {
    return r.cancelDef
}

