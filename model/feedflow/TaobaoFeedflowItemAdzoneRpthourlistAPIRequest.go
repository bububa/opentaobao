package feedflow

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemAdzoneRpthourlistAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.adzone.rpthourlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemAdzoneRpthourlistAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
