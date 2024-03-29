package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpReportQueryNotItemPromotionAPIRequest 其他主体报表查询 API请求
// taobao.universalbp.report.query.not.item.promotion
//
// 其他主体报表查询
type TaobaoUniversalbpReportQueryNotItemPromotionAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// topOtherPromotionReportQueryVO
	_topOtherPromotionReportQueryVO *TopOtherPromotionReportQueryVo
}

// NewTaobaoUniversalbpReportQueryNotItemPromotionRequest 初始化TaobaoUniversalbpReportQueryNotItemPromotionAPIRequest对象
func NewTaobaoUniversalbpReportQueryNotItemPromotionRequest() *TaobaoUniversalbpReportQueryNotItemPromotionAPIRequest {
	return &TaobaoUniversalbpReportQueryNotItemPromotionAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUniversalbpReportQueryNotItemPromotionAPIRequest) Reset() {
	r._topServiceContext = nil
	r._topOtherPromotionReportQueryVO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUniversalbpReportQueryNotItemPromotionAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.report.query.not.item.promotion"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUniversalbpReportQueryNotItemPromotionAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUniversalbpReportQueryNotItemPromotionAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaoUniversalbpReportQueryNotItemPromotionAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaoUniversalbpReportQueryNotItemPromotionAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetTopOtherPromotionReportQueryVO is TopOtherPromotionReportQueryVO Setter
// topOtherPromotionReportQueryVO
func (r *TaobaoUniversalbpReportQueryNotItemPromotionAPIRequest) SetTopOtherPromotionReportQueryVO(_topOtherPromotionReportQueryVO *TopOtherPromotionReportQueryVo) error {
	r._topOtherPromotionReportQueryVO = _topOtherPromotionReportQueryVO
	r.Set("top_other_promotion_report_query_v_o", _topOtherPromotionReportQueryVO)
	return nil
}

// GetTopOtherPromotionReportQueryVO TopOtherPromotionReportQueryVO Getter
func (r TaobaoUniversalbpReportQueryNotItemPromotionAPIRequest) GetTopOtherPromotionReportQueryVO() *TopOtherPromotionReportQueryVo {
	return r._topOtherPromotionReportQueryVO
}

var poolTaobaoUniversalbpReportQueryNotItemPromotionAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUniversalbpReportQueryNotItemPromotionRequest()
	},
}

// GetTaobaoUniversalbpReportQueryNotItemPromotionRequest 从 sync.Pool 获取 TaobaoUniversalbpReportQueryNotItemPromotionAPIRequest
func GetTaobaoUniversalbpReportQueryNotItemPromotionAPIRequest() *TaobaoUniversalbpReportQueryNotItemPromotionAPIRequest {
	return poolTaobaoUniversalbpReportQueryNotItemPromotionAPIRequest.Get().(*TaobaoUniversalbpReportQueryNotItemPromotionAPIRequest)
}

// ReleaseTaobaoUniversalbpReportQueryNotItemPromotionAPIRequest 将 TaobaoUniversalbpReportQueryNotItemPromotionAPIRequest 放入 sync.Pool
func ReleaseTaobaoUniversalbpReportQueryNotItemPromotionAPIRequest(v *TaobaoUniversalbpReportQueryNotItemPromotionAPIRequest) {
	v.Reset()
	poolTaobaoUniversalbpReportQueryNotItemPromotionAPIRequest.Put(v)
}
