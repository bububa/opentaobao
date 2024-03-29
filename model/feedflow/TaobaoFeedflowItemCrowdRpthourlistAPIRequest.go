package feedflow

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemCrowdRpthourlistAPIRequest 超级推荐【商品推广】定向分时报表查询 API请求
// taobao.feedflow.item.crowd.rpthourlist
//
// 广告主定向分时数据查询，支持广告主查询最近90天内某一天的定向维度分时报表数据
type TaobaoFeedflowItemCrowdRpthourlistAPIRequest struct {
	model.Params
	// 查询参数
	_rptQuery *RptQueryDto
}

// NewTaobaoFeedflowItemCrowdRpthourlistRequest 初始化TaobaoFeedflowItemCrowdRpthourlistAPIRequest对象
func NewTaobaoFeedflowItemCrowdRpthourlistRequest() *TaobaoFeedflowItemCrowdRpthourlistAPIRequest {
	return &TaobaoFeedflowItemCrowdRpthourlistAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFeedflowItemCrowdRpthourlistAPIRequest) Reset() {
	r._rptQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCrowdRpthourlistAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.crowd.rpthourlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCrowdRpthourlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFeedflowItemCrowdRpthourlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRptQuery is RptQuery Setter
// 查询参数
func (r *TaobaoFeedflowItemCrowdRpthourlistAPIRequest) SetRptQuery(_rptQuery *RptQueryDto) error {
	r._rptQuery = _rptQuery
	r.Set("rpt_query", _rptQuery)
	return nil
}

// GetRptQuery RptQuery Getter
func (r TaobaoFeedflowItemCrowdRpthourlistAPIRequest) GetRptQuery() *RptQueryDto {
	return r._rptQuery
}

var poolTaobaoFeedflowItemCrowdRpthourlistAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFeedflowItemCrowdRpthourlistRequest()
	},
}

// GetTaobaoFeedflowItemCrowdRpthourlistRequest 从 sync.Pool 获取 TaobaoFeedflowItemCrowdRpthourlistAPIRequest
func GetTaobaoFeedflowItemCrowdRpthourlistAPIRequest() *TaobaoFeedflowItemCrowdRpthourlistAPIRequest {
	return poolTaobaoFeedflowItemCrowdRpthourlistAPIRequest.Get().(*TaobaoFeedflowItemCrowdRpthourlistAPIRequest)
}

// ReleaseTaobaoFeedflowItemCrowdRpthourlistAPIRequest 将 TaobaoFeedflowItemCrowdRpthourlistAPIRequest 放入 sync.Pool
func ReleaseTaobaoFeedflowItemCrowdRpthourlistAPIRequest(v *TaobaoFeedflowItemCrowdRpthourlistAPIRequest) {
	v.Reset()
	poolTaobaoFeedflowItemCrowdRpthourlistAPIRequest.Put(v)
}
