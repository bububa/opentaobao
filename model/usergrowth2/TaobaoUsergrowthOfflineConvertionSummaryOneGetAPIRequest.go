package usergrowth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUsergrowthOfflineConvertionSummaryOneGetAPIRequest
手淘用增 线下业务 转化t+1汇总数据 API请求
taobao.usergrowth.offline.convertion.summary.one.get

淘系用户增长团队-线下拉新业务，对线下渠道提供mapi，目的是为了给有开发能力的渠道提供返数功能，方便渠道对手下的推广者结算（t+1汇总） */
type TaobaoUsergrowthOfflineConvertionSummaryOneGetAPIRequest struct {
	model.Params
	// 入参
	_query *OfflineMapiQuery
}

// NewTaobaoUsergrowthOfflineConvertionSummaryOneGetRequest 初始化TaobaoUsergrowthOfflineConvertionSummaryOneGetAPIRequest对象
func NewTaobaoUsergrowthOfflineConvertionSummaryOneGetRequest() *TaobaoUsergrowthOfflineConvertionSummaryOneGetAPIRequest {
	return &TaobaoUsergrowthOfflineConvertionSummaryOneGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUsergrowthOfflineConvertionSummaryOneGetAPIRequest) GetApiMethodName() string {
	return "taobao.usergrowth.offline.convertion.summary.one.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUsergrowthOfflineConvertionSummaryOneGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Query Setter
// 入参
func (r *TaobaoUsergrowthOfflineConvertionSummaryOneGetAPIRequest) SetQuery(_query *OfflineMapiQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// Get Query Getter
func (r TaobaoUsergrowthOfflineConvertionSummaryOneGetAPIRequest) GetQuery() *OfflineMapiQuery {
	return r._query
}
