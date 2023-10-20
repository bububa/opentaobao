package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemAdzoneRptdailylistAPIRequest 资源包分日数据查询 API请求
// taobao.feedflow.item.adzone.rptdailylist
//
// 资源包分日数据查询
type TaobaoFeedflowItemAdzoneRptdailylistAPIRequest struct {
	model.Params
	// 查询参数
	_rptQueryDTO *RptQueryDto
}

// NewTaobaoFeedflowItemAdzoneRptdailylistRequest 初始化TaobaoFeedflowItemAdzoneRptdailylistAPIRequest对象
func NewTaobaoFeedflowItemAdzoneRptdailylistRequest() *TaobaoFeedflowItemAdzoneRptdailylistAPIRequest {
	return &TaobaoFeedflowItemAdzoneRptdailylistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemAdzoneRptdailylistAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.adzone.rptdailylist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemAdzoneRptdailylistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFeedflowItemAdzoneRptdailylistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRptQueryDTO is RptQueryDTO Setter
// 查询参数
func (r *TaobaoFeedflowItemAdzoneRptdailylistAPIRequest) SetRptQueryDTO(_rptQueryDTO *RptQueryDto) error {
	r._rptQueryDTO = _rptQueryDTO
	r.Set("rpt_query_d_t_o", _rptQueryDTO)
	return nil
}

// GetRptQueryDTO RptQueryDTO Getter
func (r TaobaoFeedflowItemAdzoneRptdailylistAPIRequest) GetRptQueryDTO() *RptQueryDto {
	return r._rptQueryDTO
}
