package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowitemcrowdrpthourlistAPIRequest 超级推荐【商品推广】定向分时报表查询 API请求
// taobao.feedflow.item.crowd.rpthourlist
//
// 广告主定向分时数据查询，支持广告主查询最近90天内某一天的定向维度分时报表数据
type TaobaofeedflowitemcrowdrpthourlistAPIRequest struct {
	model.Params
	// 查询参数
	_rptQuery *RptQueryDto
}

// NewTaobaofeedflowitemcrowdrpthourlistRequest 初始化TaobaofeedflowitemcrowdrpthourlistAPIRequest对象
func NewTaobaofeedflowitemcrowdrpthourlistRequest() *TaobaofeedflowitemcrowdrpthourlistAPIRequest {
	return &TaobaofeedflowitemcrowdrpthourlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofeedflowitemcrowdrpthourlistAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.crowd.rpthourlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofeedflowitemcrowdrpthourlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofeedflowitemcrowdrpthourlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRptQuery is RptQuery Setter
// 查询参数
func (r *TaobaofeedflowitemcrowdrpthourlistAPIRequest) SetRptQuery(_rptQuery *RptQueryDto) error {
	r._rptQuery = _rptQuery
	r.Set("rpt_query", _rptQuery)
	return nil
}

// GetRptQuery RptQuery Getter
func (r TaobaofeedflowitemcrowdrpthourlistAPIRequest) GetRptQuery() *RptQueryDto {
	return r._rptQuery
}
