package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家地址库修改 API请求
taobao.logistics.address.modify

卖家地址库修改
*/
type TaobaoLogisticsAddressModifyRequest struct {
    model.Params
    // 地址库ID
    _contactId   int64
    // 联系人姓名<br/><font color='red'>长度不可超过20个字节</font>
    _contactName   string
    // 所在省
    _province   string
    // 所在市
    _city   string
    // 区、县<br/><br><font color='red'>如果所在地区是海外的可以为空，否则为必参</font>
    _country   string
    // 详细街道地址，不需要重复填写省/市/区
    _addr   string
    // 地区邮政编码<br/><br><font color='red'>如果所在地区是海外的可以为空，否则为必参</font>
    _zipCode   string
    // 电话号码,手机与电话必需有一个
    _phone   string
    // 手机号码，手机与电话必需有一个 <br><font color='red'>手机号码不能超过20位</font>
    _mobilePhone   string
    // 公司名称,<br/><br><font color='red'>公司名称长能不能超过20字节</font>
    _sellerCompany   string
    // 备注,<br><font color='red'>备注不能超过256字节</font>
    _memo   string
    // 默认取货地址。<br><br/><font color='red'>选择此项(true)，将当前地址设为默认取货地址，撤消原默认取货地址</font>
    _getDef   bool
    // 默认退货地址。<br><br/><font color='red'>选择此项(true)，将当前地址设为默认退货地址，撤消原默认退货地址</font>
    _cancelDef   bool
}

// 初始化TaobaoLogisticsAddressModifyRequest对象
func NewTaobaoLogisticsAddressModifyRequest() *TaobaoLogisticsAddressModifyRequest{
    return &TaobaoLogisticsAddressModifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsAddressModifyRequest) GetApiMethodName() string {
    return "taobao.logistics.address.modify"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsAddressModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ContactId Setter
// 地址库ID
func (r *TaobaoLogisticsAddressModifyRequest) SetContactId(_contactId int64) error {
    r._contactId = _contactId
    r.Set("contact_id", _contactId)
    return nil
}

// ContactId Getter
func (r TaobaoLogisticsAddressModifyRequest) GetContactId() int64 {
    return r._contactId
}
// ContactName Setter
// 联系人姓名<br/><font color='red'>长度不可超过20个字节</font>
func (r *TaobaoLogisticsAddressModifyRequest) SetContactName(_contactName string) error {
    r._contactName = _contactName
    r.Set("contact_name", _contactName)
    return nil
}

// ContactName Getter
func (r TaobaoLogisticsAddressModifyRequest) GetContactName() string {
    return r._contactName
}
// Province Setter
// 所在省
func (r *TaobaoLogisticsAddressModifyRequest) SetProvince(_province string) error {
    r._province = _province
    r.Set("province", _province)
    return nil
}

// Province Getter
func (r TaobaoLogisticsAddressModifyRequest) GetProvince() string {
    return r._province
}
// City Setter
// 所在市
func (r *TaobaoLogisticsAddressModifyRequest) SetCity(_city string) error {
    r._city = _city
    r.Set("city", _city)
    return nil
}

// City Getter
func (r TaobaoLogisticsAddressModifyRequest) GetCity() string {
    return r._city
}
// Country Setter
// 区、县<br/><br><font color='red'>如果所在地区是海外的可以为空，否则为必参</font>
func (r *TaobaoLogisticsAddressModifyRequest) SetCountry(_country string) error {
    r._country = _country
    r.Set("country", _country)
    return nil
}

// Country Getter
func (r TaobaoLogisticsAddressModifyRequest) GetCountry() string {
    return r._country
}
// Addr Setter
// 详细街道地址，不需要重复填写省/市/区
func (r *TaobaoLogisticsAddressModifyRequest) SetAddr(_addr string) error {
    r._addr = _addr
    r.Set("addr", _addr)
    return nil
}

// Addr Getter
func (r TaobaoLogisticsAddressModifyRequest) GetAddr() string {
    return r._addr
}
// ZipCode Setter
// 地区邮政编码<br/><br><font color='red'>如果所在地区是海外的可以为空，否则为必参</font>
func (r *TaobaoLogisticsAddressModifyRequest) SetZipCode(_zipCode string) error {
    r._zipCode = _zipCode
    r.Set("zip_code", _zipCode)
    return nil
}

// ZipCode Getter
func (r TaobaoLogisticsAddressModifyRequest) GetZipCode() string {
    return r._zipCode
}
// Phone Setter
// 电话号码,手机与电话必需有一个
func (r *TaobaoLogisticsAddressModifyRequest) SetPhone(_phone string) error {
    r._phone = _phone
    r.Set("phone", _phone)
    return nil
}

// Phone Getter
func (r TaobaoLogisticsAddressModifyRequest) GetPhone() string {
    return r._phone
}
// MobilePhone Setter
// 手机号码，手机与电话必需有一个 <br><font color='red'>手机号码不能超过20位</font>
func (r *TaobaoLogisticsAddressModifyRequest) SetMobilePhone(_mobilePhone string) error {
    r._mobilePhone = _mobilePhone
    r.Set("mobile_phone", _mobilePhone)
    return nil
}

// MobilePhone Getter
func (r TaobaoLogisticsAddressModifyRequest) GetMobilePhone() string {
    return r._mobilePhone
}
// SellerCompany Setter
// 公司名称,<br/><br><font color='red'>公司名称长能不能超过20字节</font>
func (r *TaobaoLogisticsAddressModifyRequest) SetSellerCompany(_sellerCompany string) error {
    r._sellerCompany = _sellerCompany
    r.Set("seller_company", _sellerCompany)
    return nil
}

// SellerCompany Getter
func (r TaobaoLogisticsAddressModifyRequest) GetSellerCompany() string {
    return r._sellerCompany
}
// Memo Setter
// 备注,<br><font color='red'>备注不能超过256字节</font>
func (r *TaobaoLogisticsAddressModifyRequest) SetMemo(_memo string) error {
    r._memo = _memo
    r.Set("memo", _memo)
    return nil
}

// Memo Getter
func (r TaobaoLogisticsAddressModifyRequest) GetMemo() string {
    return r._memo
}
// GetDef Setter
// 默认取货地址。<br><br/><font color='red'>选择此项(true)，将当前地址设为默认取货地址，撤消原默认取货地址</font>
func (r *TaobaoLogisticsAddressModifyRequest) SetGetDef(_getDef bool) error {
    r._getDef = _getDef
    r.Set("get_def", _getDef)
    return nil
}

// GetDef Getter
func (r TaobaoLogisticsAddressModifyRequest) GetGetDef() bool {
    return r._getDef
}
// CancelDef Setter
// 默认退货地址。<br><br/><font color='red'>选择此项(true)，将当前地址设为默认退货地址，撤消原默认退货地址</font>
func (r *TaobaoLogisticsAddressModifyRequest) SetCancelDef(_cancelDef bool) error {
    r._cancelDef = _cancelDef
    r.Set("cancel_def", _cancelDef)
    return nil
}

// CancelDef Getter
func (r TaobaoLogisticsAddressModifyRequest) GetCancelDef() bool {
    return r._cancelDef
}
