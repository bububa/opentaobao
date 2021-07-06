package fundplatform

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFundplatformCardordersStatusSendedAPIRequest 制卡商通知实体卡发货完成 API请求
// alibaba.fundplatform.cardorders.status.sended
//
// 当制卡商将实体卡发货完成后，需要调用该接口，通知我们已发货。
type AlibabaFundplatformCardordersStatusSendedAPIRequest struct {
	model.Params
	// 物流单号
	_logisticsOrderId string
	// 物流商名称
	_logisticsCompany string
	// 子制卡单ID
	_cardOrderId int64
}

// NewAlibabaFundplatformCardordersStatusSendedRequest 初始化AlibabaFundplatformCardordersStatusSendedAPIRequest对象
func NewAlibabaFundplatformCardordersStatusSendedRequest() *AlibabaFundplatformCardordersStatusSendedAPIRequest {
	return &AlibabaFundplatformCardordersStatusSendedAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaFundplatformCardordersStatusSendedAPIRequest) GetApiMethodName() string {
	return "alibaba.fundplatform.cardorders.status.sended"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaFundplatformCardordersStatusSendedAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetLogisticsOrderId is LogisticsOrderId Setter
// 物流单号
func (r *AlibabaFundplatformCardordersStatusSendedAPIRequest) SetLogisticsOrderId(_logisticsOrderId string) error {
	r._logisticsOrderId = _logisticsOrderId
	r.Set("logistics_order_id", _logisticsOrderId)
	return nil
}

// GetLogisticsOrderId LogisticsOrderId Getter
func (r AlibabaFundplatformCardordersStatusSendedAPIRequest) GetLogisticsOrderId() string {
	return r._logisticsOrderId
}

// SetLogisticsCompany is LogisticsCompany Setter
// 物流商名称
func (r *AlibabaFundplatformCardordersStatusSendedAPIRequest) SetLogisticsCompany(_logisticsCompany string) error {
	r._logisticsCompany = _logisticsCompany
	r.Set("logistics_company", _logisticsCompany)
	return nil
}

// GetLogisticsCompany LogisticsCompany Getter
func (r AlibabaFundplatformCardordersStatusSendedAPIRequest) GetLogisticsCompany() string {
	return r._logisticsCompany
}

// SetCardOrderId is CardOrderId Setter
// 子制卡单ID
func (r *AlibabaFundplatformCardordersStatusSendedAPIRequest) SetCardOrderId(_cardOrderId int64) error {
	r._cardOrderId = _cardOrderId
	r.Set("card_order_id", _cardOrderId)
	return nil
}

// GetCardOrderId CardOrderId Getter
func (r AlibabaFundplatformCardordersStatusSendedAPIRequest) GetCardOrderId() int64 {
	return r._cardOrderId
}
