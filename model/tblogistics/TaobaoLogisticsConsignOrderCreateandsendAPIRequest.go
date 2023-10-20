package tblogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsconsignordercreateandsendAPIRequest 创建订单并发货 API请求
// taobao.logistics.consign.order.createandsend
//
// 创建物流订单，并发货。
type TaobaologisticsconsignordercreateandsendAPIRequest struct {
	model.Params
	// 发件人街道地址
	_sAddress string
	// 市
	_sCityName string
	// 发件人名称
	_sName string
	// 省
	_sProvName string
	// 区
	_sDistName string
	// 电话号码
	_sTelephone string
	// 手机号码
	_sMobilePhone string
	// 发件人出编
	_sZipCode string
	// 电话号码
	_rTelephone string
	// 省
	_rProvName string
	// 收件人街道地址
	_rAddress string
	// 市
	_rCityName string
	// 手机号码
	_rMobilePhone string
	// 区
	_rDistName string
	// 收件人邮编
	_rZipCode string
	// 收件人名称
	_rName string
	// 运单号
	_mailNo string
	// 费用承担方式 1买家承担运费 2卖家承担运费
	_shipping string
	// 物品的json数据。
	_itemJsonString string
	// 发件人区域ID
	_sAreaId int64
	// 物流公司ID
	_companyId int64
	// 收件人区域ID
	_rAreaId int64
	// 交易流水号，淘外订单号或者商家内部交易流水号
	_tradeId int64
	// 订单来源，值选择：30
	_orderSource int64
	// 用户ID
	_userId int64
	// 物流订单物流类型，值固定选择：2
	_logisType int64
	// 订单类型，值固定选择：30
	_orderType int64
}

// NewTaobaologisticsconsignordercreateandsendRequest 初始化TaobaologisticsconsignordercreateandsendAPIRequest对象
func NewTaobaologisticsconsignordercreateandsendRequest() *TaobaologisticsconsignordercreateandsendAPIRequest {
	return &TaobaologisticsconsignordercreateandsendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticsconsignordercreateandsendAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.consign.order.createandsend"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticsconsignordercreateandsendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticsconsignordercreateandsendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSAddress is SAddress Setter
// 发件人街道地址
func (r *TaobaologisticsconsignordercreateandsendAPIRequest) SetSAddress(_sAddress string) error {
	r._sAddress = _sAddress
	r.Set("s_address", _sAddress)
	return nil
}

// GetSAddress SAddress Getter
func (r TaobaologisticsconsignordercreateandsendAPIRequest) GetSAddress() string {
	return r._sAddress
}

// SetSCityName is SCityName Setter
// 市
func (r *TaobaologisticsconsignordercreateandsendAPIRequest) SetSCityName(_sCityName string) error {
	r._sCityName = _sCityName
	r.Set("s_city_name", _sCityName)
	return nil
}

// GetSCityName SCityName Getter
func (r TaobaologisticsconsignordercreateandsendAPIRequest) GetSCityName() string {
	return r._sCityName
}

// SetSName is SName Setter
// 发件人名称
func (r *TaobaologisticsconsignordercreateandsendAPIRequest) SetSName(_sName string) error {
	r._sName = _sName
	r.Set("s_name", _sName)
	return nil
}

// GetSName SName Getter
func (r TaobaologisticsconsignordercreateandsendAPIRequest) GetSName() string {
	return r._sName
}

// SetSProvName is SProvName Setter
// 省
func (r *TaobaologisticsconsignordercreateandsendAPIRequest) SetSProvName(_sProvName string) error {
	r._sProvName = _sProvName
	r.Set("s_prov_name", _sProvName)
	return nil
}

// GetSProvName SProvName Getter
func (r TaobaologisticsconsignordercreateandsendAPIRequest) GetSProvName() string {
	return r._sProvName
}

// SetSDistName is SDistName Setter
// 区
func (r *TaobaologisticsconsignordercreateandsendAPIRequest) SetSDistName(_sDistName string) error {
	r._sDistName = _sDistName
	r.Set("s_dist_name", _sDistName)
	return nil
}

// GetSDistName SDistName Getter
func (r TaobaologisticsconsignordercreateandsendAPIRequest) GetSDistName() string {
	return r._sDistName
}

// SetSTelephone is STelephone Setter
// 电话号码
func (r *TaobaologisticsconsignordercreateandsendAPIRequest) SetSTelephone(_sTelephone string) error {
	r._sTelephone = _sTelephone
	r.Set("s_telephone", _sTelephone)
	return nil
}

// GetSTelephone STelephone Getter
func (r TaobaologisticsconsignordercreateandsendAPIRequest) GetSTelephone() string {
	return r._sTelephone
}

// SetSMobilePhone is SMobilePhone Setter
// 手机号码
func (r *TaobaologisticsconsignordercreateandsendAPIRequest) SetSMobilePhone(_sMobilePhone string) error {
	r._sMobilePhone = _sMobilePhone
	r.Set("s_mobile_phone", _sMobilePhone)
	return nil
}

// GetSMobilePhone SMobilePhone Getter
func (r TaobaologisticsconsignordercreateandsendAPIRequest) GetSMobilePhone() string {
	return r._sMobilePhone
}

// SetSZipCode is SZipCode Setter
// 发件人出编
func (r *TaobaologisticsconsignordercreateandsendAPIRequest) SetSZipCode(_sZipCode string) error {
	r._sZipCode = _sZipCode
	r.Set("s_zip_code", _sZipCode)
	return nil
}

// GetSZipCode SZipCode Getter
func (r TaobaologisticsconsignordercreateandsendAPIRequest) GetSZipCode() string {
	return r._sZipCode
}

// SetRTelephone is RTelephone Setter
// 电话号码
func (r *TaobaologisticsconsignordercreateandsendAPIRequest) SetRTelephone(_rTelephone string) error {
	r._rTelephone = _rTelephone
	r.Set("r_telephone", _rTelephone)
	return nil
}

// GetRTelephone RTelephone Getter
func (r TaobaologisticsconsignordercreateandsendAPIRequest) GetRTelephone() string {
	return r._rTelephone
}

// SetRProvName is RProvName Setter
// 省
func (r *TaobaologisticsconsignordercreateandsendAPIRequest) SetRProvName(_rProvName string) error {
	r._rProvName = _rProvName
	r.Set("r_prov_name", _rProvName)
	return nil
}

// GetRProvName RProvName Getter
func (r TaobaologisticsconsignordercreateandsendAPIRequest) GetRProvName() string {
	return r._rProvName
}

// SetRAddress is RAddress Setter
// 收件人街道地址
func (r *TaobaologisticsconsignordercreateandsendAPIRequest) SetRAddress(_rAddress string) error {
	r._rAddress = _rAddress
	r.Set("r_address", _rAddress)
	return nil
}

// GetRAddress RAddress Getter
func (r TaobaologisticsconsignordercreateandsendAPIRequest) GetRAddress() string {
	return r._rAddress
}

// SetRCityName is RCityName Setter
// 市
func (r *TaobaologisticsconsignordercreateandsendAPIRequest) SetRCityName(_rCityName string) error {
	r._rCityName = _rCityName
	r.Set("r_city_name", _rCityName)
	return nil
}

// GetRCityName RCityName Getter
func (r TaobaologisticsconsignordercreateandsendAPIRequest) GetRCityName() string {
	return r._rCityName
}

// SetRMobilePhone is RMobilePhone Setter
// 手机号码
func (r *TaobaologisticsconsignordercreateandsendAPIRequest) SetRMobilePhone(_rMobilePhone string) error {
	r._rMobilePhone = _rMobilePhone
	r.Set("r_mobile_phone", _rMobilePhone)
	return nil
}

// GetRMobilePhone RMobilePhone Getter
func (r TaobaologisticsconsignordercreateandsendAPIRequest) GetRMobilePhone() string {
	return r._rMobilePhone
}

// SetRDistName is RDistName Setter
// 区
func (r *TaobaologisticsconsignordercreateandsendAPIRequest) SetRDistName(_rDistName string) error {
	r._rDistName = _rDistName
	r.Set("r_dist_name", _rDistName)
	return nil
}

// GetRDistName RDistName Getter
func (r TaobaologisticsconsignordercreateandsendAPIRequest) GetRDistName() string {
	return r._rDistName
}

// SetRZipCode is RZipCode Setter
// 收件人邮编
func (r *TaobaologisticsconsignordercreateandsendAPIRequest) SetRZipCode(_rZipCode string) error {
	r._rZipCode = _rZipCode
	r.Set("r_zip_code", _rZipCode)
	return nil
}

// GetRZipCode RZipCode Getter
func (r TaobaologisticsconsignordercreateandsendAPIRequest) GetRZipCode() string {
	return r._rZipCode
}

// SetRName is RName Setter
// 收件人名称
func (r *TaobaologisticsconsignordercreateandsendAPIRequest) SetRName(_rName string) error {
	r._rName = _rName
	r.Set("r_name", _rName)
	return nil
}

// GetRName RName Getter
func (r TaobaologisticsconsignordercreateandsendAPIRequest) GetRName() string {
	return r._rName
}

// SetMailNo is MailNo Setter
// 运单号
func (r *TaobaologisticsconsignordercreateandsendAPIRequest) SetMailNo(_mailNo string) error {
	r._mailNo = _mailNo
	r.Set("mail_no", _mailNo)
	return nil
}

// GetMailNo MailNo Getter
func (r TaobaologisticsconsignordercreateandsendAPIRequest) GetMailNo() string {
	return r._mailNo
}

// SetShipping is Shipping Setter
// 费用承担方式 1买家承担运费 2卖家承担运费
func (r *TaobaologisticsconsignordercreateandsendAPIRequest) SetShipping(_shipping string) error {
	r._shipping = _shipping
	r.Set("shipping", _shipping)
	return nil
}

// GetShipping Shipping Getter
func (r TaobaologisticsconsignordercreateandsendAPIRequest) GetShipping() string {
	return r._shipping
}

// SetItemJsonString is ItemJsonString Setter
// 物品的json数据。
func (r *TaobaologisticsconsignordercreateandsendAPIRequest) SetItemJsonString(_itemJsonString string) error {
	r._itemJsonString = _itemJsonString
	r.Set("item_json_string", _itemJsonString)
	return nil
}

// GetItemJsonString ItemJsonString Getter
func (r TaobaologisticsconsignordercreateandsendAPIRequest) GetItemJsonString() string {
	return r._itemJsonString
}

// SetSAreaId is SAreaId Setter
// 发件人区域ID
func (r *TaobaologisticsconsignordercreateandsendAPIRequest) SetSAreaId(_sAreaId int64) error {
	r._sAreaId = _sAreaId
	r.Set("s_area_id", _sAreaId)
	return nil
}

// GetSAreaId SAreaId Getter
func (r TaobaologisticsconsignordercreateandsendAPIRequest) GetSAreaId() int64 {
	return r._sAreaId
}

// SetCompanyId is CompanyId Setter
// 物流公司ID
func (r *TaobaologisticsconsignordercreateandsendAPIRequest) SetCompanyId(_companyId int64) error {
	r._companyId = _companyId
	r.Set("company_id", _companyId)
	return nil
}

// GetCompanyId CompanyId Getter
func (r TaobaologisticsconsignordercreateandsendAPIRequest) GetCompanyId() int64 {
	return r._companyId
}

// SetRAreaId is RAreaId Setter
// 收件人区域ID
func (r *TaobaologisticsconsignordercreateandsendAPIRequest) SetRAreaId(_rAreaId int64) error {
	r._rAreaId = _rAreaId
	r.Set("r_area_id", _rAreaId)
	return nil
}

// GetRAreaId RAreaId Getter
func (r TaobaologisticsconsignordercreateandsendAPIRequest) GetRAreaId() int64 {
	return r._rAreaId
}

// SetTradeId is TradeId Setter
// 交易流水号，淘外订单号或者商家内部交易流水号
func (r *TaobaologisticsconsignordercreateandsendAPIRequest) SetTradeId(_tradeId int64) error {
	r._tradeId = _tradeId
	r.Set("trade_id", _tradeId)
	return nil
}

// GetTradeId TradeId Getter
func (r TaobaologisticsconsignordercreateandsendAPIRequest) GetTradeId() int64 {
	return r._tradeId
}

// SetOrderSource is OrderSource Setter
// 订单来源，值选择：30
func (r *TaobaologisticsconsignordercreateandsendAPIRequest) SetOrderSource(_orderSource int64) error {
	r._orderSource = _orderSource
	r.Set("order_source", _orderSource)
	return nil
}

// GetOrderSource OrderSource Getter
func (r TaobaologisticsconsignordercreateandsendAPIRequest) GetOrderSource() int64 {
	return r._orderSource
}

// SetUserId is UserId Setter
// 用户ID
func (r *TaobaologisticsconsignordercreateandsendAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaologisticsconsignordercreateandsendAPIRequest) GetUserId() int64 {
	return r._userId
}

// SetLogisType is LogisType Setter
// 物流订单物流类型，值固定选择：2
func (r *TaobaologisticsconsignordercreateandsendAPIRequest) SetLogisType(_logisType int64) error {
	r._logisType = _logisType
	r.Set("logis_type", _logisType)
	return nil
}

// GetLogisType LogisType Getter
func (r TaobaologisticsconsignordercreateandsendAPIRequest) GetLogisType() int64 {
	return r._logisType
}

// SetOrderType is OrderType Setter
// 订单类型，值固定选择：30
func (r *TaobaologisticsconsignordercreateandsendAPIRequest) SetOrderType(_orderType int64) error {
	r._orderType = _orderType
	r.Set("order_type", _orderType)
	return nil
}

// GetOrderType OrderType Getter
func (r TaobaologisticsconsignordercreateandsendAPIRequest) GetOrderType() int64 {
	return r._orderType
}
