package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJstMiniappCrowdCreateAPIRequest
小程序活动创建 API请求
taobao.jst.miniapp.crowd.create

小程序活动创建 */
type TaobaoJstMiniappCrowdCreateAPIRequest struct {
	model.Params
	// 活动开始时间，开始时间和结束时间不能超过1个月
	_endDate string
	// 活动描述
	_description string
	// 活动开始时间
	_startDate string
	// 活动名称
	_crowdName string
}

// NewTaobaoJstMiniappCrowdCreateRequest 初始化TaobaoJstMiniappCrowdCreateAPIRequest对象
func NewTaobaoJstMiniappCrowdCreateRequest() *TaobaoJstMiniappCrowdCreateAPIRequest {
	return &TaobaoJstMiniappCrowdCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstMiniappCrowdCreateAPIRequest) GetApiMethodName() string {
	return "taobao.jst.miniapp.crowd.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstMiniappCrowdCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is EndDate Setter
// 活动开始时间，开始时间和结束时间不能超过1个月
func (r *TaobaoJstMiniappCrowdCreateAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// Get EndDate Getter
func (r TaobaoJstMiniappCrowdCreateAPIRequest) GetEndDate() string {
	return r._endDate
}

// Set is Description Setter
// 活动描述
func (r *TaobaoJstMiniappCrowdCreateAPIRequest) SetDescription(_description string) error {
	r._description = _description
	r.Set("description", _description)
	return nil
}

// Get Description Getter
func (r TaobaoJstMiniappCrowdCreateAPIRequest) GetDescription() string {
	return r._description
}

// Set is StartDate Setter
// 活动开始时间
func (r *TaobaoJstMiniappCrowdCreateAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// Get StartDate Getter
func (r TaobaoJstMiniappCrowdCreateAPIRequest) GetStartDate() string {
	return r._startDate
}

// Set is CrowdName Setter
// 活动名称
func (r *TaobaoJstMiniappCrowdCreateAPIRequest) SetCrowdName(_crowdName string) error {
	r._crowdName = _crowdName
	r.Set("crowd_name", _crowdName)
	return nil
}

// Get CrowdName Getter
func (r TaobaoJstMiniappCrowdCreateAPIRequest) GetCrowdName() string {
	return r._crowdName
}
