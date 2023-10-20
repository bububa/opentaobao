package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowitemcreativerptdailylistAPIRequest 创意分日数据查询 API请求
// taobao.feedflow.item.creative.rptdailylist
//
// 创意分日数据查询
type TaobaofeedflowitemcreativerptdailylistAPIRequest struct {
	model.Params
	// 查询条件
	_rptQueryDTO *RptQueryDto
}

// NewTaobaofeedflowitemcreativerptdailylistRequest 初始化TaobaofeedflowitemcreativerptdailylistAPIRequest对象
func NewTaobaofeedflowitemcreativerptdailylistRequest() *TaobaofeedflowitemcreativerptdailylistAPIRequest {
	return &TaobaofeedflowitemcreativerptdailylistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofeedflowitemcreativerptdailylistAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.creative.rptdailylist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofeedflowitemcreativerptdailylistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofeedflowitemcreativerptdailylistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRptQueryDTO is RptQueryDTO Setter
// 查询条件
func (r *TaobaofeedflowitemcreativerptdailylistAPIRequest) SetRptQueryDTO(_rptQueryDTO *RptQueryDto) error {
	r._rptQueryDTO = _rptQueryDTO
	r.Set("rpt_query_d_t_o", _rptQueryDTO)
	return nil
}

// GetRptQueryDTO RptQueryDTO Getter
func (r TaobaofeedflowitemcreativerptdailylistAPIRequest) GetRptQueryDTO() *RptQueryDto {
	return r._rptQueryDTO
}
