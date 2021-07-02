package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsAddressAddAPIRequest 卖家地址库新增接口 API请求
// taobao.logistics.address.add
//
// 通过此接口新增卖家地址库,卖家最多可添加5条地址库,新增第一条卖家地址，将会自动设为默认地址库
type TaobaoLogisticsAddressAddAPIRequest struct {
	model.Params
	// 联系人姓名 <font color='red'>长度不可超过20个字节</font>
	_contactName string
	// 所在省
	_province string
	// 所在市
	_city string
	// 区、县<br/><br><font color='red'>如果所在地区是海外的可以为空，否则为必参</font>
	_country string
	// 详细街道地址，不需要重复填写省/市/区
	_addr string
	// 地区邮政编码<br/><br><font color='red'>如果所在地区是海外的可以为空，否则为必参</font>
	_zipCode string
	// 电话号码,手机与电话必需有一个
	_phone string
	// 手机号码，手机与电话必需有一个<br/><br><font color='red'>手机号码不能超过20位</font>
	_mobilePhone string
	// 公司名称,<br><font color="red">公司名称长能不能超过20字节</font>
	_sellerCompany string
	// 备注,<br><font color='red'>备注不能超过256字节</font>
	_memo string
	// 默认取货地址。<br><br/><font color='red'>选择此项(true)，将当前地址设为默认取货地址，撤消原默认取货地址</font>
	_getDef bool
	// 默认退货地址。<br><br/><font color='red'>选择此项(true)，将当前地址设为默认退货地址，撤消原默认退货地址</font>
	_cancelDef bool
}

// NewTaobaoLogisticsAddressAddRequest 初始化TaobaoLogisticsAddressAddAPIRequest对象
func NewTaobaoLogisticsAddressAddRequest() *TaobaoLogisticsAddressAddAPIRequest {
	return &TaobaoLogisticsAddressAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsAddressAddAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.address.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsAddressAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetContactName is ContactName Setter
// 联系人姓名 <font color='red'>长度不可超过20个字节</font>
func (r *TaobaoLogisticsAddressAddAPIRequest) SetContactName(_contactName string) error {
	r._contactName = _contactName
	r.Set("contact_name", _contactName)
	return nil
}

// GetContactName ContactName Getter
func (r TaobaoLogisticsAddressAddAPIRequest) GetContactName() string {
	return r._contactName
}

// SetProvince is Province Setter
// 所在省
func (r *TaobaoLogisticsAddressAddAPIRequest) SetProvince(_province string) error {
	r._province = _province
	r.Set("province", _province)
	return nil
}

// GetProvince Province Getter
func (r TaobaoLogisticsAddressAddAPIRequest) GetProvince() string {
	return r._province
}

// SetCity is City Setter
// 所在市
func (r *TaobaoLogisticsAddressAddAPIRequest) SetCity(_city string) error {
	r._city = _city
	r.Set("city", _city)
	return nil
}

// GetCity City Getter
func (r TaobaoLogisticsAddressAddAPIRequest) GetCity() string {
	return r._city
}

// SetCountry is Country Setter
// 区、县<br/><br><font color='red'>如果所在地区是海外的可以为空，否则为必参</font>
func (r *TaobaoLogisticsAddressAddAPIRequest) SetCountry(_country string) error {
	r._country = _country
	r.Set("country", _country)
	return nil
}

// GetCountry Country Getter
func (r TaobaoLogisticsAddressAddAPIRequest) GetCountry() string {
	return r._country
}

// SetAddr is Addr Setter
// 详细街道地址，不需要重复填写省/市/区
func (r *TaobaoLogisticsAddressAddAPIRequest) SetAddr(_addr string) error {
	r._addr = _addr
	r.Set("addr", _addr)
	return nil
}

// GetAddr Addr Getter
func (r TaobaoLogisticsAddressAddAPIRequest) GetAddr() string {
	return r._addr
}

// SetZipCode is ZipCode Setter
// 地区邮政编码<br/><br><font color='red'>如果所在地区是海外的可以为空，否则为必参</font>
func (r *TaobaoLogisticsAddressAddAPIRequest) SetZipCode(_zipCode string) error {
	r._zipCode = _zipCode
	r.Set("zip_code", _zipCode)
	return nil
}

// GetZipCode ZipCode Getter
func (r TaobaoLogisticsAddressAddAPIRequest) GetZipCode() string {
	return r._zipCode
}

// SetPhone is Phone Setter
// 电话号码,手机与电话必需有一个
func (r *TaobaoLogisticsAddressAddAPIRequest) SetPhone(_phone string) error {
	r._phone = _phone
	r.Set("phone", _phone)
	return nil
}

// GetPhone Phone Getter
func (r TaobaoLogisticsAddressAddAPIRequest) GetPhone() string {
	return r._phone
}

// SetMobilePhone is MobilePhone Setter
// 手机号码，手机与电话必需有一个<br/><br><font color='red'>手机号码不能超过20位</font>
func (r *TaobaoLogisticsAddressAddAPIRequest) SetMobilePhone(_mobilePhone string) error {
	r._mobilePhone = _mobilePhone
	r.Set("mobile_phone", _mobilePhone)
	return nil
}

// GetMobilePhone MobilePhone Getter
func (r TaobaoLogisticsAddressAddAPIRequest) GetMobilePhone() string {
	return r._mobilePhone
}

// SetSellerCompany is SellerCompany Setter
// 公司名称,<br><font color="red">公司名称长能不能超过20字节</font>
func (r *TaobaoLogisticsAddressAddAPIRequest) SetSellerCompany(_sellerCompany string) error {
	r._sellerCompany = _sellerCompany
	r.Set("seller_company", _sellerCompany)
	return nil
}

// GetSellerCompany SellerCompany Getter
func (r TaobaoLogisticsAddressAddAPIRequest) GetSellerCompany() string {
	return r._sellerCompany
}

// SetMemo is Memo Setter
// 备注,<br><font color='red'>备注不能超过256字节</font>
func (r *TaobaoLogisticsAddressAddAPIRequest) SetMemo(_memo string) error {
	r._memo = _memo
	r.Set("memo", _memo)
	return nil
}

// GetMemo Memo Getter
func (r TaobaoLogisticsAddressAddAPIRequest) GetMemo() string {
	return r._memo
}

// SetGetDef is GetDef Setter
// 默认取货地址。<br><br/><font color='red'>选择此项(true)，将当前地址设为默认取货地址，撤消原默认取货地址</font>
func (r *TaobaoLogisticsAddressAddAPIRequest) SetGetDef(_getDef bool) error {
	r._getDef = _getDef
	r.Set("get_def", _getDef)
	return nil
}

// GetGetDef GetDef Getter
func (r TaobaoLogisticsAddressAddAPIRequest) GetGetDef() bool {
	return r._getDef
}

// SetCancelDef is CancelDef Setter
// 默认退货地址。<br><br/><font color='red'>选择此项(true)，将当前地址设为默认退货地址，撤消原默认退货地址</font>
func (r *TaobaoLogisticsAddressAddAPIRequest) SetCancelDef(_cancelDef bool) error {
	r._cancelDef = _cancelDef
	r.Set("cancel_def", _cancelDef)
	return nil
}

// GetCancelDef CancelDef Getter
func (r TaobaoLogisticsAddressAddAPIRequest) GetCancelDef() bool {
	return r._cancelDef
}
