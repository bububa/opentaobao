package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowitemadgrouprptdailylistAPIRequest 推广单元分日数据查询 API请求
// taobao.feedflow.item.adgroup.rptdailylist
//
// 推广单元分日数据查询
type TaobaofeedflowitemadgrouprptdailylistAPIRequest struct {
	model.Params
	// 查询条件
	_rptQueryDTO *RptQueryDto
}

// NewTaobaofeedflowitemadgrouprptdailylistRequest 初始化TaobaofeedflowitemadgrouprptdailylistAPIRequest对象
func NewTaobaofeedflowitemadgrouprptdailylistRequest() *TaobaofeedflowitemadgrouprptdailylistAPIRequest {
	return &TaobaofeedflowitemadgrouprptdailylistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofeedflowitemadgrouprptdailylistAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.adgroup.rptdailylist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofeedflowitemadgrouprptdailylistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofeedflowitemadgrouprptdailylistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRptQueryDTO is RptQueryDTO Setter
// 查询条件
func (r *TaobaofeedflowitemadgrouprptdailylistAPIRequest) SetRptQueryDTO(_rptQueryDTO *RptQueryDto) error {
	r._rptQueryDTO = _rptQueryDTO
	r.Set("rpt_query_d_t_o", _rptQueryDTO)
	return nil
}

// GetRptQueryDTO RptQueryDTO Getter
func (r TaobaofeedflowitemadgrouprptdailylistAPIRequest) GetRptQueryDTO() *RptQueryDto {
	return r._rptQueryDTO
}
