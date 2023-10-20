package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenmallrefundsubmitAPIRequest 提交OpenMall退款单物流 API请求
// taobao.openmall.refund.submit
//
// 提交OpenMall退款单物流
type TaobaoopenmallrefundsubmitAPIRequest struct {
	model.Params
	// 渠道
	_distributor string
	// 物流公司编码
	_logisticsCompanyCode string
	// 物流公司名称
	_logisticsCompanyName string
	// 快递单号
	_logisticsNo string
	// 退款单ID
	_refundId int64
}

// NewTaobaoopenmallrefundsubmitRequest 初始化TaobaoopenmallrefundsubmitAPIRequest对象
func NewTaobaoopenmallrefundsubmitRequest() *TaobaoopenmallrefundsubmitAPIRequest {
	return &TaobaoopenmallrefundsubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopenmallrefundsubmitAPIRequest) GetApiMethodName() string {
	return "taobao.openmall.refund.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopenmallrefundsubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopenmallrefundsubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDistributor is Distributor Setter
// 渠道
func (r *TaobaoopenmallrefundsubmitAPIRequest) SetDistributor(_distributor string) error {
	r._distributor = _distributor
	r.Set("distributor", _distributor)
	return nil
}

// GetDistributor Distributor Getter
func (r TaobaoopenmallrefundsubmitAPIRequest) GetDistributor() string {
	return r._distributor
}

// SetLogisticsCompanyCode is LogisticsCompanyCode Setter
// 物流公司编码
func (r *TaobaoopenmallrefundsubmitAPIRequest) SetLogisticsCompanyCode(_logisticsCompanyCode string) error {
	r._logisticsCompanyCode = _logisticsCompanyCode
	r.Set("logistics_company_code", _logisticsCompanyCode)
	return nil
}

// GetLogisticsCompanyCode LogisticsCompanyCode Getter
func (r TaobaoopenmallrefundsubmitAPIRequest) GetLogisticsCompanyCode() string {
	return r._logisticsCompanyCode
}

// SetLogisticsCompanyName is LogisticsCompanyName Setter
// 物流公司名称
func (r *TaobaoopenmallrefundsubmitAPIRequest) SetLogisticsCompanyName(_logisticsCompanyName string) error {
	r._logisticsCompanyName = _logisticsCompanyName
	r.Set("logistics_company_name", _logisticsCompanyName)
	return nil
}

// GetLogisticsCompanyName LogisticsCompanyName Getter
func (r TaobaoopenmallrefundsubmitAPIRequest) GetLogisticsCompanyName() string {
	return r._logisticsCompanyName
}

// SetLogisticsNo is LogisticsNo Setter
// 快递单号
func (r *TaobaoopenmallrefundsubmitAPIRequest) SetLogisticsNo(_logisticsNo string) error {
	r._logisticsNo = _logisticsNo
	r.Set("logistics_no", _logisticsNo)
	return nil
}

// GetLogisticsNo LogisticsNo Getter
func (r TaobaoopenmallrefundsubmitAPIRequest) GetLogisticsNo() string {
	return r._logisticsNo
}

// SetRefundId is RefundId Setter
// 退款单ID
func (r *TaobaoopenmallrefundsubmitAPIRequest) SetRefundId(_refundId int64) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r TaobaoopenmallrefundsubmitAPIRequest) GetRefundId() int64 {
	return r._refundId
}
