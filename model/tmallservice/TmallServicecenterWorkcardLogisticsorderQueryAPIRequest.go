package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterWorkcardLogisticsorderQueryAPIRequest
物流单信息查询 API请求
tmall.servicecenter.workcard.logisticsorder.query

物流订单信息查询API */
type TmallServicecenterWorkcardLogisticsorderQueryAPIRequest struct {
	model.Params
	// 物流单号
	_logisticsOrderId int64
}

// NewTmallServicecenterWorkcardLogisticsorderQueryRequest 初始化TmallServicecenterWorkcardLogisticsorderQueryAPIRequest对象
func NewTmallServicecenterWorkcardLogisticsorderQueryRequest() *TmallServicecenterWorkcardLogisticsorderQueryAPIRequest {
	return &TmallServicecenterWorkcardLogisticsorderQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardLogisticsorderQueryAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.logisticsorder.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardLogisticsorderQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is LogisticsOrderId Setter
// 物流单号
func (r *TmallServicecenterWorkcardLogisticsorderQueryAPIRequest) SetLogisticsOrderId(_logisticsOrderId int64) error {
	r._logisticsOrderId = _logisticsOrderId
	r.Set("logistics_order_id", _logisticsOrderId)
	return nil
}

// Get LogisticsOrderId Getter
func (r TmallServicecenterWorkcardLogisticsorderQueryAPIRequest) GetLogisticsOrderId() int64 {
	return r._logisticsOrderId
}
