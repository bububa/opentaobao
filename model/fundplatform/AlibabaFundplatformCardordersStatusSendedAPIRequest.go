package fundplatform

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabafundplatformcardordersstatussendedAPIRequest 制卡商通知实体卡发货完成 API请求
// alibaba.fundplatform.cardorders.status.sended
//
// 当制卡商将实体卡发货完成后，需要调用该接口，通知我们已发货。
type AlibabafundplatformcardordersstatussendedAPIRequest struct {
	model.Params
	// 物流单号
	_logisticsOrderId string
	// 物流商名称
	_logisticsCompany string
	// 子制卡单ID
	_cardOrderId int64
}

// NewAlibabafundplatformcardordersstatussendedRequest 初始化AlibabafundplatformcardordersstatussendedAPIRequest对象
func NewAlibabafundplatformcardordersstatussendedRequest() *AlibabafundplatformcardordersstatussendedAPIRequest {
	return &AlibabafundplatformcardordersstatussendedAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabafundplatformcardordersstatussendedAPIRequest) GetApiMethodName() string {
	return "alibaba.fundplatform.cardorders.status.sended"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabafundplatformcardordersstatussendedAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabafundplatformcardordersstatussendedAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLogisticsOrderId is LogisticsOrderId Setter
// 物流单号
func (r *AlibabafundplatformcardordersstatussendedAPIRequest) SetLogisticsOrderId(_logisticsOrderId string) error {
	r._logisticsOrderId = _logisticsOrderId
	r.Set("logistics_order_id", _logisticsOrderId)
	return nil
}

// GetLogisticsOrderId LogisticsOrderId Getter
func (r AlibabafundplatformcardordersstatussendedAPIRequest) GetLogisticsOrderId() string {
	return r._logisticsOrderId
}

// SetLogisticsCompany is LogisticsCompany Setter
// 物流商名称
func (r *AlibabafundplatformcardordersstatussendedAPIRequest) SetLogisticsCompany(_logisticsCompany string) error {
	r._logisticsCompany = _logisticsCompany
	r.Set("logistics_company", _logisticsCompany)
	return nil
}

// GetLogisticsCompany LogisticsCompany Getter
func (r AlibabafundplatformcardordersstatussendedAPIRequest) GetLogisticsCompany() string {
	return r._logisticsCompany
}

// SetCardOrderId is CardOrderId Setter
// 子制卡单ID
func (r *AlibabafundplatformcardordersstatussendedAPIRequest) SetCardOrderId(_cardOrderId int64) error {
	r._cardOrderId = _cardOrderId
	r.Set("card_order_id", _cardOrderId)
	return nil
}

// GetCardOrderId CardOrderId Getter
func (r AlibabafundplatformcardordersstatussendedAPIRequest) GetCardOrderId() int64 {
	return r._cardOrderId
}
