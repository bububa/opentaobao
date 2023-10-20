package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowitemcampaignrptdailylistAPIRequest 推广计划分日数据查询 API请求
// taobao.feedflow.item.campaign.rptdailylist
//
// 推广计划分日数据查询
type TaobaofeedflowitemcampaignrptdailylistAPIRequest struct {
	model.Params
	// 查询条件
	_rptQueryDTO *RptQueryDto
}

// NewTaobaofeedflowitemcampaignrptdailylistRequest 初始化TaobaofeedflowitemcampaignrptdailylistAPIRequest对象
func NewTaobaofeedflowitemcampaignrptdailylistRequest() *TaobaofeedflowitemcampaignrptdailylistAPIRequest {
	return &TaobaofeedflowitemcampaignrptdailylistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofeedflowitemcampaignrptdailylistAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.campaign.rptdailylist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofeedflowitemcampaignrptdailylistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofeedflowitemcampaignrptdailylistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRptQueryDTO is RptQueryDTO Setter
// 查询条件
func (r *TaobaofeedflowitemcampaignrptdailylistAPIRequest) SetRptQueryDTO(_rptQueryDTO *RptQueryDto) error {
	r._rptQueryDTO = _rptQueryDTO
	r.Set("rpt_query_d_t_o", _rptQueryDTO)
	return nil
}

// GetRptQueryDTO RptQueryDTO Getter
func (r TaobaofeedflowitemcampaignrptdailylistAPIRequest) GetRptQueryDTO() *RptQueryDto {
	return r._rptQueryDTO
}
