package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpreportquerynotitempromotionAPIRequest 其他主体报表查询 API请求
// taobao.universalbp.report.query.not.item.promotion
//
// 其他主体报表查询
type TaobaouniversalbpreportquerynotitempromotionAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// topOtherPromotionReportQueryVO
	_topOtherPromotionReportQueryVO *TopOtherPromotionReportQueryVo
}

// NewTaobaouniversalbpreportquerynotitempromotionRequest 初始化TaobaouniversalbpreportquerynotitempromotionAPIRequest对象
func NewTaobaouniversalbpreportquerynotitempromotionRequest() *TaobaouniversalbpreportquerynotitempromotionAPIRequest {
	return &TaobaouniversalbpreportquerynotitempromotionAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaouniversalbpreportquerynotitempromotionAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.report.query.not.item.promotion"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaouniversalbpreportquerynotitempromotionAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaouniversalbpreportquerynotitempromotionAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaouniversalbpreportquerynotitempromotionAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaouniversalbpreportquerynotitempromotionAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetTopOtherPromotionReportQueryVO is TopOtherPromotionReportQueryVO Setter
// topOtherPromotionReportQueryVO
func (r *TaobaouniversalbpreportquerynotitempromotionAPIRequest) SetTopOtherPromotionReportQueryVO(_topOtherPromotionReportQueryVO *TopOtherPromotionReportQueryVo) error {
	r._topOtherPromotionReportQueryVO = _topOtherPromotionReportQueryVO
	r.Set("top_other_promotion_report_query_v_o", _topOtherPromotionReportQueryVO)
	return nil
}

// GetTopOtherPromotionReportQueryVO TopOtherPromotionReportQueryVO Getter
func (r TaobaouniversalbpreportquerynotitempromotionAPIRequest) GetTopOtherPromotionReportQueryVO() *TopOtherPromotionReportQueryVo {
	return r._topOtherPromotionReportQueryVO
}
