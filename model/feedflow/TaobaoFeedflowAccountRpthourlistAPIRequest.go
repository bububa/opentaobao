package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowaccountrpthourlistAPIRequest 超级推荐广告主分时报表查询 API请求
// taobao.feedflow.account.rpthourlist
//
// 广告主分时报表查询，支持广告主查询最近90天内某一天的账户维度分时报表数据
type TaobaofeedflowaccountrpthourlistAPIRequest struct {
	model.Params
	// 查询参数
	_rptQuery *RptQueryDto
}

// NewTaobaofeedflowaccountrpthourlistRequest 初始化TaobaofeedflowaccountrpthourlistAPIRequest对象
func NewTaobaofeedflowaccountrpthourlistRequest() *TaobaofeedflowaccountrpthourlistAPIRequest {
	return &TaobaofeedflowaccountrpthourlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofeedflowaccountrpthourlistAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.account.rpthourlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofeedflowaccountrpthourlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofeedflowaccountrpthourlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRptQuery is RptQuery Setter
// 查询参数
func (r *TaobaofeedflowaccountrpthourlistAPIRequest) SetRptQuery(_rptQuery *RptQueryDto) error {
	r._rptQuery = _rptQuery
	r.Set("rpt_query", _rptQuery)
	return nil
}

// GetRptQuery RptQuery Getter
func (r TaobaofeedflowaccountrpthourlistAPIRequest) GetRptQuery() *RptQueryDto {
	return r._rptQuery
}
