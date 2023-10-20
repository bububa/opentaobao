package cainiaolocker

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoendpointlockertoporderwithholdAPIRequest 代扣支付 API请求
// cainiao.endpoint.locker.top.order.withhold
//
// 提供代扣，允许有一笔欠款。
type CainiaoendpointlockertoporderwithholdAPIRequest struct {
	model.Params
	// 柜子公司编码
	_companyCode string
	// 柜子id
	_guiId string
	// 柜子订单编码
	_orderCode string
	// 运单号
	_mailNo string
	// 开放用户id
	_openUserId string
	// 扩展字段
	_extra string
	// 订单类型(0-取件业务，1-寄件业务，2-派样业务)
	_orderType int64
	// 代扣金额（全额），单位：分
	_totalFee int64
}

// NewCainiaoendpointlockertoporderwithholdRequest 初始化CainiaoendpointlockertoporderwithholdAPIRequest对象
func NewCainiaoendpointlockertoporderwithholdRequest() *CainiaoendpointlockertoporderwithholdAPIRequest {
	return &CainiaoendpointlockertoporderwithholdAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoendpointlockertoporderwithholdAPIRequest) GetApiMethodName() string {
	return "cainiao.endpoint.locker.top.order.withhold"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoendpointlockertoporderwithholdAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoendpointlockertoporderwithholdAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCompanyCode is CompanyCode Setter
// 柜子公司编码
func (r *CainiaoendpointlockertoporderwithholdAPIRequest) SetCompanyCode(_companyCode string) error {
	r._companyCode = _companyCode
	r.Set("company_code", _companyCode)
	return nil
}

// GetCompanyCode CompanyCode Getter
func (r CainiaoendpointlockertoporderwithholdAPIRequest) GetCompanyCode() string {
	return r._companyCode
}

// SetGuiId is GuiId Setter
// 柜子id
func (r *CainiaoendpointlockertoporderwithholdAPIRequest) SetGuiId(_guiId string) error {
	r._guiId = _guiId
	r.Set("gui_id", _guiId)
	return nil
}

// GetGuiId GuiId Getter
func (r CainiaoendpointlockertoporderwithholdAPIRequest) GetGuiId() string {
	return r._guiId
}

// SetOrderCode is OrderCode Setter
// 柜子订单编码
func (r *CainiaoendpointlockertoporderwithholdAPIRequest) SetOrderCode(_orderCode string) error {
	r._orderCode = _orderCode
	r.Set("order_code", _orderCode)
	return nil
}

// GetOrderCode OrderCode Getter
func (r CainiaoendpointlockertoporderwithholdAPIRequest) GetOrderCode() string {
	return r._orderCode
}

// SetMailNo is MailNo Setter
// 运单号
func (r *CainiaoendpointlockertoporderwithholdAPIRequest) SetMailNo(_mailNo string) error {
	r._mailNo = _mailNo
	r.Set("mail_no", _mailNo)
	return nil
}

// GetMailNo MailNo Getter
func (r CainiaoendpointlockertoporderwithholdAPIRequest) GetMailNo() string {
	return r._mailNo
}

// SetOpenUserId is OpenUserId Setter
// 开放用户id
func (r *CainiaoendpointlockertoporderwithholdAPIRequest) SetOpenUserId(_openUserId string) error {
	r._openUserId = _openUserId
	r.Set("open_user_id", _openUserId)
	return nil
}

// GetOpenUserId OpenUserId Getter
func (r CainiaoendpointlockertoporderwithholdAPIRequest) GetOpenUserId() string {
	return r._openUserId
}

// SetExtra is Extra Setter
// 扩展字段
func (r *CainiaoendpointlockertoporderwithholdAPIRequest) SetExtra(_extra string) error {
	r._extra = _extra
	r.Set("extra", _extra)
	return nil
}

// GetExtra Extra Getter
func (r CainiaoendpointlockertoporderwithholdAPIRequest) GetExtra() string {
	return r._extra
}

// SetOrderType is OrderType Setter
// 订单类型(0-取件业务，1-寄件业务，2-派样业务)
func (r *CainiaoendpointlockertoporderwithholdAPIRequest) SetOrderType(_orderType int64) error {
	r._orderType = _orderType
	r.Set("order_type", _orderType)
	return nil
}

// GetOrderType OrderType Getter
func (r CainiaoendpointlockertoporderwithholdAPIRequest) GetOrderType() int64 {
	return r._orderType
}

// SetTotalFee is TotalFee Setter
// 代扣金额（全额），单位：分
func (r *CainiaoendpointlockertoporderwithholdAPIRequest) SetTotalFee(_totalFee int64) error {
	r._totalFee = _totalFee
	r.Set("total_fee", _totalFee)
	return nil
}

// GetTotalFee TotalFee Getter
func (r CainiaoendpointlockertoporderwithholdAPIRequest) GetTotalFee() int64 {
	return r._totalFee
}
