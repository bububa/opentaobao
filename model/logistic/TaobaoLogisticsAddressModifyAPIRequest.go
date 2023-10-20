package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsaddressmodifyAPIRequest 卖家地址库修改 API请求
// taobao.logistics.address.modify
//
// 卖家地址库修改
type TaobaologisticsaddressmodifyAPIRequest struct {
	model.Params
	// 联系人姓名<br/><font color='red'>长度不可超过20个字节</font>
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
	// 手机号码，手机与电话必需有一个 <br><font color='red'>手机号码不能超过20位</font>
	_mobilePhone string
	// 公司名称,<br/><br><font color='red'>公司名称长能不能超过20字节</font>
	_sellerCompany string
	// 备注,<br><font color='red'>备注不能超过256字节</font>
	_memo string
	// 地址库ID
	_contactId int64
	// 默认取货地址。<br><br/><font color='red'>选择此项(true)，将当前地址设为默认取货地址，撤消原默认取货地址</font>
	_getDef bool
	// 默认退货地址。<br><br/><font color='red'>选择此项(true)，将当前地址设为默认退货地址，撤消原默认退货地址</font>
	_cancelDef bool
}

// NewTaobaologisticsaddressmodifyRequest 初始化TaobaologisticsaddressmodifyAPIRequest对象
func NewTaobaologisticsaddressmodifyRequest() *TaobaologisticsaddressmodifyAPIRequest {
	return &TaobaologisticsaddressmodifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticsaddressmodifyAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.address.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticsaddressmodifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticsaddressmodifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContactName is ContactName Setter
// 联系人姓名&lt;br/&gt;&lt;font color=&#39;red&#39;&gt;长度不可超过20个字节&lt;/font&gt;
func (r *TaobaologisticsaddressmodifyAPIRequest) SetContactName(_contactName string) error {
	r._contactName = _contactName
	r.Set("contact_name", _contactName)
	return nil
}

// GetContactName ContactName Getter
func (r TaobaologisticsaddressmodifyAPIRequest) GetContactName() string {
	return r._contactName
}

// SetProvince is Province Setter
// 所在省
func (r *TaobaologisticsaddressmodifyAPIRequest) SetProvince(_province string) error {
	r._province = _province
	r.Set("province", _province)
	return nil
}

// GetProvince Province Getter
func (r TaobaologisticsaddressmodifyAPIRequest) GetProvince() string {
	return r._province
}

// SetCity is City Setter
// 所在市
func (r *TaobaologisticsaddressmodifyAPIRequest) SetCity(_city string) error {
	r._city = _city
	r.Set("city", _city)
	return nil
}

// GetCity City Getter
func (r TaobaologisticsaddressmodifyAPIRequest) GetCity() string {
	return r._city
}

// SetCountry is Country Setter
// 区、县&lt;br/&gt;&lt;br&gt;&lt;font color=&#39;red&#39;&gt;如果所在地区是海外的可以为空，否则为必参&lt;/font&gt;
func (r *TaobaologisticsaddressmodifyAPIRequest) SetCountry(_country string) error {
	r._country = _country
	r.Set("country", _country)
	return nil
}

// GetCountry Country Getter
func (r TaobaologisticsaddressmodifyAPIRequest) GetCountry() string {
	return r._country
}

// SetAddr is Addr Setter
// 详细街道地址，不需要重复填写省/市/区
func (r *TaobaologisticsaddressmodifyAPIRequest) SetAddr(_addr string) error {
	r._addr = _addr
	r.Set("addr", _addr)
	return nil
}

// GetAddr Addr Getter
func (r TaobaologisticsaddressmodifyAPIRequest) GetAddr() string {
	return r._addr
}

// SetZipCode is ZipCode Setter
// 地区邮政编码&lt;br/&gt;&lt;br&gt;&lt;font color=&#39;red&#39;&gt;如果所在地区是海外的可以为空，否则为必参&lt;/font&gt;
func (r *TaobaologisticsaddressmodifyAPIRequest) SetZipCode(_zipCode string) error {
	r._zipCode = _zipCode
	r.Set("zip_code", _zipCode)
	return nil
}

// GetZipCode ZipCode Getter
func (r TaobaologisticsaddressmodifyAPIRequest) GetZipCode() string {
	return r._zipCode
}

// SetPhone is Phone Setter
// 电话号码,手机与电话必需有一个
func (r *TaobaologisticsaddressmodifyAPIRequest) SetPhone(_phone string) error {
	r._phone = _phone
	r.Set("phone", _phone)
	return nil
}

// GetPhone Phone Getter
func (r TaobaologisticsaddressmodifyAPIRequest) GetPhone() string {
	return r._phone
}

// SetMobilePhone is MobilePhone Setter
// 手机号码，手机与电话必需有一个 &lt;br&gt;&lt;font color=&#39;red&#39;&gt;手机号码不能超过20位&lt;/font&gt;
func (r *TaobaologisticsaddressmodifyAPIRequest) SetMobilePhone(_mobilePhone string) error {
	r._mobilePhone = _mobilePhone
	r.Set("mobile_phone", _mobilePhone)
	return nil
}

// GetMobilePhone MobilePhone Getter
func (r TaobaologisticsaddressmodifyAPIRequest) GetMobilePhone() string {
	return r._mobilePhone
}

// SetSellerCompany is SellerCompany Setter
// 公司名称,&lt;br/&gt;&lt;br&gt;&lt;font color=&#39;red&#39;&gt;公司名称长能不能超过20字节&lt;/font&gt;
func (r *TaobaologisticsaddressmodifyAPIRequest) SetSellerCompany(_sellerCompany string) error {
	r._sellerCompany = _sellerCompany
	r.Set("seller_company", _sellerCompany)
	return nil
}

// GetSellerCompany SellerCompany Getter
func (r TaobaologisticsaddressmodifyAPIRequest) GetSellerCompany() string {
	return r._sellerCompany
}

// SetMemo is Memo Setter
// 备注,&lt;br&gt;&lt;font color=&#39;red&#39;&gt;备注不能超过256字节&lt;/font&gt;
func (r *TaobaologisticsaddressmodifyAPIRequest) SetMemo(_memo string) error {
	r._memo = _memo
	r.Set("memo", _memo)
	return nil
}

// GetMemo Memo Getter
func (r TaobaologisticsaddressmodifyAPIRequest) GetMemo() string {
	return r._memo
}

// SetContactId is ContactId Setter
// 地址库ID
func (r *TaobaologisticsaddressmodifyAPIRequest) SetContactId(_contactId int64) error {
	r._contactId = _contactId
	r.Set("contact_id", _contactId)
	return nil
}

// GetContactId ContactId Getter
func (r TaobaologisticsaddressmodifyAPIRequest) GetContactId() int64 {
	return r._contactId
}

// SetGetDef is GetDef Setter
// 默认取货地址。&lt;br&gt;&lt;br/&gt;&lt;font color=&#39;red&#39;&gt;选择此项(true)，将当前地址设为默认取货地址，撤消原默认取货地址&lt;/font&gt;
func (r *TaobaologisticsaddressmodifyAPIRequest) SetGetDef(_getDef bool) error {
	r._getDef = _getDef
	r.Set("get_def", _getDef)
	return nil
}

// GetGetDef GetDef Getter
func (r TaobaologisticsaddressmodifyAPIRequest) GetGetDef() bool {
	return r._getDef
}

// SetCancelDef is CancelDef Setter
// 默认退货地址。&lt;br&gt;&lt;br/&gt;&lt;font color=&#39;red&#39;&gt;选择此项(true)，将当前地址设为默认退货地址，撤消原默认退货地址&lt;/font&gt;
func (r *TaobaologisticsaddressmodifyAPIRequest) SetCancelDef(_cancelDef bool) error {
	r._cancelDef = _cancelDef
	r.Set("cancel_def", _cancelDef)
	return nil
}

// GetCancelDef CancelDef Getter
func (r TaobaologisticsaddressmodifyAPIRequest) GetCancelDef() bool {
	return r._cancelDef
}
