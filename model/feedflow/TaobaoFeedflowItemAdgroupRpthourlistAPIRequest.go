package feedflow

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemAdgroupRpthourlistAPIRequest 超级推荐【商品推广】单元分时报表查询 API请求
// taobao.feedflow.item.adgroup.rpthourlist
//
// 广告主推广组分时数据查询，支持广告主查询最近90天内某一天的单元维度分时报表数据
type TaobaoFeedflowItemAdgroupRpthourlistAPIRequest struct {
	model.Params
	// 查询参数
	_rptQuery *RptQueryDto
}

// NewTaobaoFeedflowItemAdgroupRpthourlistRequest 初始化TaobaoFeedflowItemAdgroupRpthourlistAPIRequest对象
func NewTaobaoFeedflowItemAdgroupRpthourlistRequest() *TaobaoFeedflowItemAdgroupRpthourlistAPIRequest {
	return &TaobaoFeedflowItemAdgroupRpthourlistAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFeedflowItemAdgroupRpthourlistAPIRequest) Reset() {
	r._rptQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemAdgroupRpthourlistAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.adgroup.rpthourlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemAdgroupRpthourlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFeedflowItemAdgroupRpthourlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRptQuery is RptQuery Setter
// 查询参数
func (r *TaobaoFeedflowItemAdgroupRpthourlistAPIRequest) SetRptQuery(_rptQuery *RptQueryDto) error {
	r._rptQuery = _rptQuery
	r.Set("rpt_query", _rptQuery)
	return nil
}

// GetRptQuery RptQuery Getter
func (r TaobaoFeedflowItemAdgroupRpthourlistAPIRequest) GetRptQuery() *RptQueryDto {
	return r._rptQuery
}

var poolTaobaoFeedflowItemAdgroupRpthourlistAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFeedflowItemAdgroupRpthourlistRequest()
	},
}

// GetTaobaoFeedflowItemAdgroupRpthourlistRequest 从 sync.Pool 获取 TaobaoFeedflowItemAdgroupRpthourlistAPIRequest
func GetTaobaoFeedflowItemAdgroupRpthourlistAPIRequest() *TaobaoFeedflowItemAdgroupRpthourlistAPIRequest {
	return poolTaobaoFeedflowItemAdgroupRpthourlistAPIRequest.Get().(*TaobaoFeedflowItemAdgroupRpthourlistAPIRequest)
}

// ReleaseTaobaoFeedflowItemAdgroupRpthourlistAPIRequest 将 TaobaoFeedflowItemAdgroupRpthourlistAPIRequest 放入 sync.Pool
func ReleaseTaobaoFeedflowItemAdgroupRpthourlistAPIRequest(v *TaobaoFeedflowItemAdgroupRpthourlistAPIRequest) {
	v.Reset()
	poolTaobaoFeedflowItemAdgroupRpthourlistAPIRequest.Put(v)
}
