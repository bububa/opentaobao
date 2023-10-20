package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowitemcreativerpthourlistAPIRequest 超级推荐【商品推广】创意分时报表查询 API请求
// taobao.feedflow.item.creative.rpthourlist
//
// 创意分时数据查询，支持广告主查询最近90天内某一天的创意维度分时报表数据
type TaobaofeedflowitemcreativerpthourlistAPIRequest struct {
	model.Params
	// 查询参数
	_rptQuery *RptQueryDto
}

// NewTaobaofeedflowitemcreativerpthourlistRequest 初始化TaobaofeedflowitemcreativerpthourlistAPIRequest对象
func NewTaobaofeedflowitemcreativerpthourlistRequest() *TaobaofeedflowitemcreativerpthourlistAPIRequest {
	return &TaobaofeedflowitemcreativerpthourlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofeedflowitemcreativerpthourlistAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.creative.rpthourlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofeedflowitemcreativerpthourlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofeedflowitemcreativerpthourlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRptQuery is RptQuery Setter
// 查询参数
func (r *TaobaofeedflowitemcreativerpthourlistAPIRequest) SetRptQuery(_rptQuery *RptQueryDto) error {
	r._rptQuery = _rptQuery
	r.Set("rpt_query", _rptQuery)
	return nil
}

// GetRptQuery RptQuery Getter
func (r TaobaofeedflowitemcreativerpthourlistAPIRequest) GetRptQuery() *RptQueryDto {
	return r._rptQuery
}
