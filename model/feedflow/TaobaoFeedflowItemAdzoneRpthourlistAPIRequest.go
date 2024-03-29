package feedflow

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemAdzoneRpthourlistAPIRequest 超级推荐【商品推广】资源位分时报表查询 API请求
// taobao.feedflow.item.adzone.rpthourlist
//
// 广告主资源包分时数据查询，支持广告主查询最近90天内某一天的资源包维度分时报表数据
type TaobaoFeedflowItemAdzoneRpthourlistAPIRequest struct {
	model.Params
	// 查询参数
	_rptQuery *RptQueryDto
}

// NewTaobaoFeedflowItemAdzoneRpthourlistRequest 初始化TaobaoFeedflowItemAdzoneRpthourlistAPIRequest对象
func NewTaobaoFeedflowItemAdzoneRpthourlistRequest() *TaobaoFeedflowItemAdzoneRpthourlistAPIRequest {
	return &TaobaoFeedflowItemAdzoneRpthourlistAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFeedflowItemAdzoneRpthourlistAPIRequest) Reset() {
	r._rptQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemAdzoneRpthourlistAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.adzone.rpthourlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemAdzoneRpthourlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFeedflowItemAdzoneRpthourlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRptQuery is RptQuery Setter
// 查询参数
func (r *TaobaoFeedflowItemAdzoneRpthourlistAPIRequest) SetRptQuery(_rptQuery *RptQueryDto) error {
	r._rptQuery = _rptQuery
	r.Set("rpt_query", _rptQuery)
	return nil
}

// GetRptQuery RptQuery Getter
func (r TaobaoFeedflowItemAdzoneRpthourlistAPIRequest) GetRptQuery() *RptQueryDto {
	return r._rptQuery
}

var poolTaobaoFeedflowItemAdzoneRpthourlistAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFeedflowItemAdzoneRpthourlistRequest()
	},
}

// GetTaobaoFeedflowItemAdzoneRpthourlistRequest 从 sync.Pool 获取 TaobaoFeedflowItemAdzoneRpthourlistAPIRequest
func GetTaobaoFeedflowItemAdzoneRpthourlistAPIRequest() *TaobaoFeedflowItemAdzoneRpthourlistAPIRequest {
	return poolTaobaoFeedflowItemAdzoneRpthourlistAPIRequest.Get().(*TaobaoFeedflowItemAdzoneRpthourlistAPIRequest)
}

// ReleaseTaobaoFeedflowItemAdzoneRpthourlistAPIRequest 将 TaobaoFeedflowItemAdzoneRpthourlistAPIRequest 放入 sync.Pool
func ReleaseTaobaoFeedflowItemAdzoneRpthourlistAPIRequest(v *TaobaoFeedflowItemAdzoneRpthourlistAPIRequest) {
	v.Reset()
	poolTaobaoFeedflowItemAdzoneRpthourlistAPIRequest.Put(v)
}
