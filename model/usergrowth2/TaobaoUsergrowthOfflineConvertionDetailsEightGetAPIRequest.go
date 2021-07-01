package usergrowth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUsergrowthOfflineConvertionDetailsEightGetAPIRequest
手淘线下拉新业务 t+8转化明细数据 API请求
taobao.usergrowth.offline.convertion.details.eight.get

手淘线下拉新业务 给合作渠道返回t+8转化明细数据 */
type TaobaoUsergrowthOfflineConvertionDetailsEightGetAPIRequest struct {
	model.Params
	// 入参
	_query *OfflineMapiQuery
}

// NewTaobaoUsergrowthOfflineConvertionDetailsEightGetRequest 初始化TaobaoUsergrowthOfflineConvertionDetailsEightGetAPIRequest对象
func NewTaobaoUsergrowthOfflineConvertionDetailsEightGetRequest() *TaobaoUsergrowthOfflineConvertionDetailsEightGetAPIRequest {
	return &TaobaoUsergrowthOfflineConvertionDetailsEightGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUsergrowthOfflineConvertionDetailsEightGetAPIRequest) GetApiMethodName() string {
	return "taobao.usergrowth.offline.convertion.details.eight.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUsergrowthOfflineConvertionDetailsEightGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Query Setter
// 入参
func (r *TaobaoUsergrowthOfflineConvertionDetailsEightGetAPIRequest) SetQuery(_query *OfflineMapiQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// Get Query Getter
func (r TaobaoUsergrowthOfflineConvertionDetailsEightGetAPIRequest) GetQuery() *OfflineMapiQuery {
	return r._query
}
