package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowitemadzonerpthourlistAPIRequest 超级推荐【商品推广】资源位分时报表查询 API请求
// taobao.feedflow.item.adzone.rpthourlist
//
// 广告主资源包分时数据查询，支持广告主查询最近90天内某一天的资源包维度分时报表数据
type TaobaofeedflowitemadzonerpthourlistAPIRequest struct {
	model.Params
	// 查询参数
	_rptQuery *RptQueryDto
}

// NewTaobaofeedflowitemadzonerpthourlistRequest 初始化TaobaofeedflowitemadzonerpthourlistAPIRequest对象
func NewTaobaofeedflowitemadzonerpthourlistRequest() *TaobaofeedflowitemadzonerpthourlistAPIRequest {
	return &TaobaofeedflowitemadzonerpthourlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofeedflowitemadzonerpthourlistAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.adzone.rpthourlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofeedflowitemadzonerpthourlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofeedflowitemadzonerpthourlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRptQuery is RptQuery Setter
// 查询参数
func (r *TaobaofeedflowitemadzonerpthourlistAPIRequest) SetRptQuery(_rptQuery *RptQueryDto) error {
	r._rptQuery = _rptQuery
	r.Set("rpt_query", _rptQuery)
	return nil
}

// GetRptQuery RptQuery Getter
func (r TaobaofeedflowitemadzonerpthourlistAPIRequest) GetRptQuery() *RptQueryDto {
	return r._rptQuery
}
