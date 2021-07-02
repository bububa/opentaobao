package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenmallRefundSubmitAPIRequest 提交OpenMall退款单物流 API请求
// taobao.openmall.refund.submit
//
// 提交OpenMall退款单物流
type TaobaoOpenmallRefundSubmitAPIRequest struct {
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

// NewTaobaoOpenmallRefundSubmitRequest 初始化TaobaoOpenmallRefundSubmitAPIRequest对象
func NewTaobaoOpenmallRefundSubmitRequest() *TaobaoOpenmallRefundSubmitAPIRequest {
	return &TaobaoOpenmallRefundSubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallRefundSubmitAPIRequest) GetApiMethodName() string {
	return "taobao.openmall.refund.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallRefundSubmitAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Distributor Setter
// 渠道
func (r *TaobaoOpenmallRefundSubmitAPIRequest) SetDistributor(_distributor string) error {
	r._distributor = _distributor
	r.Set("distributor", _distributor)
	return nil
}

// Get Distributor Getter
func (r TaobaoOpenmallRefundSubmitAPIRequest) GetDistributor() string {
	return r._distributor
}

// Set is LogisticsCompanyCode Setter
// 物流公司编码
func (r *TaobaoOpenmallRefundSubmitAPIRequest) SetLogisticsCompanyCode(_logisticsCompanyCode string) error {
	r._logisticsCompanyCode = _logisticsCompanyCode
	r.Set("logistics_company_code", _logisticsCompanyCode)
	return nil
}

// Get LogisticsCompanyCode Getter
func (r TaobaoOpenmallRefundSubmitAPIRequest) GetLogisticsCompanyCode() string {
	return r._logisticsCompanyCode
}

// Set is LogisticsCompanyName Setter
// 物流公司名称
func (r *TaobaoOpenmallRefundSubmitAPIRequest) SetLogisticsCompanyName(_logisticsCompanyName string) error {
	r._logisticsCompanyName = _logisticsCompanyName
	r.Set("logistics_company_name", _logisticsCompanyName)
	return nil
}

// Get LogisticsCompanyName Getter
func (r TaobaoOpenmallRefundSubmitAPIRequest) GetLogisticsCompanyName() string {
	return r._logisticsCompanyName
}

// Set is LogisticsNo Setter
// 快递单号
func (r *TaobaoOpenmallRefundSubmitAPIRequest) SetLogisticsNo(_logisticsNo string) error {
	r._logisticsNo = _logisticsNo
	r.Set("logistics_no", _logisticsNo)
	return nil
}

// Get LogisticsNo Getter
func (r TaobaoOpenmallRefundSubmitAPIRequest) GetLogisticsNo() string {
	return r._logisticsNo
}

// Set is RefundId Setter
// 退款单ID
func (r *TaobaoOpenmallRefundSubmitAPIRequest) SetRefundId(_refundId int64) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// Get RefundId Getter
func (r TaobaoOpenmallRefundSubmitAPIRequest) GetRefundId() int64 {
	return r._refundId
}
