package idleisv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleisvordershipAPIRequest 闲鱼订单服务商物流发货 API请求
// alibaba.idle.isv.order.ship
//
// 闲鱼开放平台服务商订单发货接口
type AlibabaidleisvordershipAPIRequest struct {
	model.Params
	// 订单号
	_bizOrderId string
	// 物流公司(底层最新发货已不使用)
	_logisticsCompany string
	// 运单号
	_shipMailNo string
	// 发货人手机号码，用于菜鸟同步物流信息[重要]
	_senderPhone string
	// 发货人地址，用于菜鸟同步物流信息[重要]
	_senderAddress string
	// 发货人姓名，用于菜鸟同步物流信息[重要]
	_senderName string
	// 快递公司编码,从alibaba.idle.logistics.companies.query查询
	_lcCode string
	// 行政区划Id，最小行政单位code，若是地区级别，则为地区级别的id；否则为城市级别的id(long型，6位)
	_senderDivisionid int64
}

// NewAlibabaidleisvordershipRequest 初始化AlibabaidleisvordershipAPIRequest对象
func NewAlibabaidleisvordershipRequest() *AlibabaidleisvordershipAPIRequest {
	return &AlibabaidleisvordershipAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidleisvordershipAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.isv.order.ship"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidleisvordershipAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidleisvordershipAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizOrderId is BizOrderId Setter
// 订单号
func (r *AlibabaidleisvordershipAPIRequest) SetBizOrderId(_bizOrderId string) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r AlibabaidleisvordershipAPIRequest) GetBizOrderId() string {
	return r._bizOrderId
}

// SetLogisticsCompany is LogisticsCompany Setter
// 物流公司(底层最新发货已不使用)
func (r *AlibabaidleisvordershipAPIRequest) SetLogisticsCompany(_logisticsCompany string) error {
	r._logisticsCompany = _logisticsCompany
	r.Set("logistics_company", _logisticsCompany)
	return nil
}

// GetLogisticsCompany LogisticsCompany Getter
func (r AlibabaidleisvordershipAPIRequest) GetLogisticsCompany() string {
	return r._logisticsCompany
}

// SetShipMailNo is ShipMailNo Setter
// 运单号
func (r *AlibabaidleisvordershipAPIRequest) SetShipMailNo(_shipMailNo string) error {
	r._shipMailNo = _shipMailNo
	r.Set("ship_mail_no", _shipMailNo)
	return nil
}

// GetShipMailNo ShipMailNo Getter
func (r AlibabaidleisvordershipAPIRequest) GetShipMailNo() string {
	return r._shipMailNo
}

// SetSenderPhone is SenderPhone Setter
// 发货人手机号码，用于菜鸟同步物流信息[重要]
func (r *AlibabaidleisvordershipAPIRequest) SetSenderPhone(_senderPhone string) error {
	r._senderPhone = _senderPhone
	r.Set("sender_phone", _senderPhone)
	return nil
}

// GetSenderPhone SenderPhone Getter
func (r AlibabaidleisvordershipAPIRequest) GetSenderPhone() string {
	return r._senderPhone
}

// SetSenderAddress is SenderAddress Setter
// 发货人地址，用于菜鸟同步物流信息[重要]
func (r *AlibabaidleisvordershipAPIRequest) SetSenderAddress(_senderAddress string) error {
	r._senderAddress = _senderAddress
	r.Set("sender_address", _senderAddress)
	return nil
}

// GetSenderAddress SenderAddress Getter
func (r AlibabaidleisvordershipAPIRequest) GetSenderAddress() string {
	return r._senderAddress
}

// SetSenderName is SenderName Setter
// 发货人姓名，用于菜鸟同步物流信息[重要]
func (r *AlibabaidleisvordershipAPIRequest) SetSenderName(_senderName string) error {
	r._senderName = _senderName
	r.Set("sender_name", _senderName)
	return nil
}

// GetSenderName SenderName Getter
func (r AlibabaidleisvordershipAPIRequest) GetSenderName() string {
	return r._senderName
}

// SetLcCode is LcCode Setter
// 快递公司编码,从alibaba.idle.logistics.companies.query查询
func (r *AlibabaidleisvordershipAPIRequest) SetLcCode(_lcCode string) error {
	r._lcCode = _lcCode
	r.Set("lc_code", _lcCode)
	return nil
}

// GetLcCode LcCode Getter
func (r AlibabaidleisvordershipAPIRequest) GetLcCode() string {
	return r._lcCode
}

// SetSenderDivisionid is SenderDivisionid Setter
// 行政区划Id，最小行政单位code，若是地区级别，则为地区级别的id；否则为城市级别的id(long型，6位)
func (r *AlibabaidleisvordershipAPIRequest) SetSenderDivisionid(_senderDivisionid int64) error {
	r._senderDivisionid = _senderDivisionid
	r.Set("sender_divisionid", _senderDivisionid)
	return nil
}

// GetSenderDivisionid SenderDivisionid Getter
func (r AlibabaidleisvordershipAPIRequest) GetSenderDivisionid() int64 {
	return r._senderDivisionid
}
