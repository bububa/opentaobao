package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpreportqueryitempromotionAPIRequest 宝贝主体报表查询 API请求
// taobao.universalbp.report.query.item.promotion
//
// 宝贝主体报表查询
type TaobaouniversalbpreportqueryitempromotionAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// topItemPromotionReportQueryVO
	_topItemPromotionReportQueryVO *TopItemPromotionReportQueryVo
}

// NewTaobaouniversalbpreportqueryitempromotionRequest 初始化TaobaouniversalbpreportqueryitempromotionAPIRequest对象
func NewTaobaouniversalbpreportqueryitempromotionRequest() *TaobaouniversalbpreportqueryitempromotionAPIRequest {
	return &TaobaouniversalbpreportqueryitempromotionAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaouniversalbpreportqueryitempromotionAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.report.query.item.promotion"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaouniversalbpreportqueryitempromotionAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaouniversalbpreportqueryitempromotionAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaouniversalbpreportqueryitempromotionAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaouniversalbpreportqueryitempromotionAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetTopItemPromotionReportQueryVO is TopItemPromotionReportQueryVO Setter
// topItemPromotionReportQueryVO
func (r *TaobaouniversalbpreportqueryitempromotionAPIRequest) SetTopItemPromotionReportQueryVO(_topItemPromotionReportQueryVO *TopItemPromotionReportQueryVo) error {
	r._topItemPromotionReportQueryVO = _topItemPromotionReportQueryVO
	r.Set("top_item_promotion_report_query_v_o", _topItemPromotionReportQueryVO)
	return nil
}

// GetTopItemPromotionReportQueryVO TopItemPromotionReportQueryVO Getter
func (r TaobaouniversalbpreportqueryitempromotionAPIRequest) GetTopItemPromotionReportQueryVO() *TopItemPromotionReportQueryVo {
	return r._topItemPromotionReportQueryVO
}
