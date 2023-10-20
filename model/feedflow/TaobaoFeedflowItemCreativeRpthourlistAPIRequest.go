package feedflow

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemCreativeRpthourlistAPIRequest 超级推荐【商品推广】创意分时报表查询 API请求
// taobao.feedflow.item.creative.rpthourlist
//
// 创意分时数据查询，支持广告主查询最近90天内某一天的创意维度分时报表数据
type TaobaoFeedflowItemCreativeRpthourlistAPIRequest struct {
	model.Params
	// 查询参数
	_rptQuery *RptQueryDto
}

// NewTaobaoFeedflowItemCreativeRpthourlistRequest 初始化TaobaoFeedflowItemCreativeRpthourlistAPIRequest对象
func NewTaobaoFeedflowItemCreativeRpthourlistRequest() *TaobaoFeedflowItemCreativeRpthourlistAPIRequest {
	return &TaobaoFeedflowItemCreativeRpthourlistAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFeedflowItemCreativeRpthourlistAPIRequest) Reset() {
	r._rptQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCreativeRpthourlistAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.creative.rpthourlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCreativeRpthourlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFeedflowItemCreativeRpthourlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRptQuery is RptQuery Setter
// 查询参数
func (r *TaobaoFeedflowItemCreativeRpthourlistAPIRequest) SetRptQuery(_rptQuery *RptQueryDto) error {
	r._rptQuery = _rptQuery
	r.Set("rpt_query", _rptQuery)
	return nil
}

// GetRptQuery RptQuery Getter
func (r TaobaoFeedflowItemCreativeRpthourlistAPIRequest) GetRptQuery() *RptQueryDto {
	return r._rptQuery
}

var poolTaobaoFeedflowItemCreativeRpthourlistAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFeedflowItemCreativeRpthourlistRequest()
	},
}

// GetTaobaoFeedflowItemCreativeRpthourlistRequest 从 sync.Pool 获取 TaobaoFeedflowItemCreativeRpthourlistAPIRequest
func GetTaobaoFeedflowItemCreativeRpthourlistAPIRequest() *TaobaoFeedflowItemCreativeRpthourlistAPIRequest {
	return poolTaobaoFeedflowItemCreativeRpthourlistAPIRequest.Get().(*TaobaoFeedflowItemCreativeRpthourlistAPIRequest)
}

// ReleaseTaobaoFeedflowItemCreativeRpthourlistAPIRequest 将 TaobaoFeedflowItemCreativeRpthourlistAPIRequest 放入 sync.Pool
func ReleaseTaobaoFeedflowItemCreativeRpthourlistAPIRequest(v *TaobaoFeedflowItemCreativeRpthourlistAPIRequest) {
	v.Reset()
	poolTaobaoFeedflowItemCreativeRpthourlistAPIRequest.Put(v)
}
