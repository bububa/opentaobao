package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpReportQueryItemPromotionAPIRequest 宝贝主体报表查询 API请求
// taobao.universalbp.report.query.item.promotion
//
// 宝贝主体报表查询
type TaobaoUniversalbpReportQueryItemPromotionAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// topItemPromotionReportQueryVO
	_topItemPromotionReportQueryVO *TopItemPromotionReportQueryVo
}

// NewTaobaoUniversalbpReportQueryItemPromotionRequest 初始化TaobaoUniversalbpReportQueryItemPromotionAPIRequest对象
func NewTaobaoUniversalbpReportQueryItemPromotionRequest() *TaobaoUniversalbpReportQueryItemPromotionAPIRequest {
	return &TaobaoUniversalbpReportQueryItemPromotionAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUniversalbpReportQueryItemPromotionAPIRequest) Reset() {
	r._topServiceContext = nil
	r._topItemPromotionReportQueryVO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUniversalbpReportQueryItemPromotionAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.report.query.item.promotion"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUniversalbpReportQueryItemPromotionAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUniversalbpReportQueryItemPromotionAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaoUniversalbpReportQueryItemPromotionAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaoUniversalbpReportQueryItemPromotionAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetTopItemPromotionReportQueryVO is TopItemPromotionReportQueryVO Setter
// topItemPromotionReportQueryVO
func (r *TaobaoUniversalbpReportQueryItemPromotionAPIRequest) SetTopItemPromotionReportQueryVO(_topItemPromotionReportQueryVO *TopItemPromotionReportQueryVo) error {
	r._topItemPromotionReportQueryVO = _topItemPromotionReportQueryVO
	r.Set("top_item_promotion_report_query_v_o", _topItemPromotionReportQueryVO)
	return nil
}

// GetTopItemPromotionReportQueryVO TopItemPromotionReportQueryVO Getter
func (r TaobaoUniversalbpReportQueryItemPromotionAPIRequest) GetTopItemPromotionReportQueryVO() *TopItemPromotionReportQueryVo {
	return r._topItemPromotionReportQueryVO
}

var poolTaobaoUniversalbpReportQueryItemPromotionAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUniversalbpReportQueryItemPromotionRequest()
	},
}

// GetTaobaoUniversalbpReportQueryItemPromotionRequest 从 sync.Pool 获取 TaobaoUniversalbpReportQueryItemPromotionAPIRequest
func GetTaobaoUniversalbpReportQueryItemPromotionAPIRequest() *TaobaoUniversalbpReportQueryItemPromotionAPIRequest {
	return poolTaobaoUniversalbpReportQueryItemPromotionAPIRequest.Get().(*TaobaoUniversalbpReportQueryItemPromotionAPIRequest)
}

// ReleaseTaobaoUniversalbpReportQueryItemPromotionAPIRequest 将 TaobaoUniversalbpReportQueryItemPromotionAPIRequest 放入 sync.Pool
func ReleaseTaobaoUniversalbpReportQueryItemPromotionAPIRequest(v *TaobaoUniversalbpReportQueryItemPromotionAPIRequest) {
	v.Reset()
	poolTaobaoUniversalbpReportQueryItemPromotionAPIRequest.Put(v)
}
