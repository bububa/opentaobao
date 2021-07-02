package usergrowth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUsergrowthOfflineConvertionDetailsGetAPIRequest 淘系用增线下转化明细 API请求
// taobao.usergrowth.offline.convertion.details.get
//
// 淘系用增增长-线下拉新：为渠道提供返回拉新转化数据接口
type TaobaoUsergrowthOfflineConvertionDetailsGetAPIRequest struct {
	model.Params
	// 入参
	_query *OfflineMapiQuery
}

// NewTaobaoUsergrowthOfflineConvertionDetailsGetRequest 初始化TaobaoUsergrowthOfflineConvertionDetailsGetAPIRequest对象
func NewTaobaoUsergrowthOfflineConvertionDetailsGetRequest() *TaobaoUsergrowthOfflineConvertionDetailsGetAPIRequest {
	return &TaobaoUsergrowthOfflineConvertionDetailsGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUsergrowthOfflineConvertionDetailsGetAPIRequest) GetApiMethodName() string {
	return "taobao.usergrowth.offline.convertion.details.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUsergrowthOfflineConvertionDetailsGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Query Setter
// 入参
func (r *TaobaoUsergrowthOfflineConvertionDetailsGetAPIRequest) SetQuery(_query *OfflineMapiQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// Get Query Getter
func (r TaobaoUsergrowthOfflineConvertionDetailsGetAPIRequest) GetQuery() *OfflineMapiQuery {
	return r._query
}
