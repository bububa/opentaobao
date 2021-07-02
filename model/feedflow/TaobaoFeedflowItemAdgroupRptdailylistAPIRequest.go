package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemAdgroupRptdailylistAPIRequest 推广单元分日数据查询 API请求
// taobao.feedflow.item.adgroup.rptdailylist
//
// 推广单元分日数据查询
type TaobaoFeedflowItemAdgroupRptdailylistAPIRequest struct {
	model.Params
	// 查询条件
	_rptQueryDTO *RptQueryDto
}

// NewTaobaoFeedflowItemAdgroupRptdailylistRequest 初始化TaobaoFeedflowItemAdgroupRptdailylistAPIRequest对象
func NewTaobaoFeedflowItemAdgroupRptdailylistRequest() *TaobaoFeedflowItemAdgroupRptdailylistAPIRequest {
	return &TaobaoFeedflowItemAdgroupRptdailylistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemAdgroupRptdailylistAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.adgroup.rptdailylist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemAdgroupRptdailylistAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRptQueryDTO is RptQueryDTO Setter
// 查询条件
func (r *TaobaoFeedflowItemAdgroupRptdailylistAPIRequest) SetRptQueryDTO(_rptQueryDTO *RptQueryDto) error {
	r._rptQueryDTO = _rptQueryDTO
	r.Set("rpt_query_d_t_o", _rptQueryDTO)
	return nil
}

// GetRptQueryDTO RptQueryDTO Getter
func (r TaobaoFeedflowItemAdgroupRptdailylistAPIRequest) GetRptQueryDTO() *RptQueryDto {
	return r._rptQueryDTO
}
