package usergrowth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUsergrowthOfflineConvertionSyncInfoGetAPIRequest 查询 用增线下业务  转化数据是否同步完成 API请求
// taobao.usergrowth.offline.convertion.sync.info.get
//
// 为手淘线下合作的渠道，提供对外查询数据是否更新完毕接口
type TaobaoUsergrowthOfflineConvertionSyncInfoGetAPIRequest struct {
	model.Params
	// 入参
	_query *OfflineConvertionSyncInfoQuery
}

// NewTaobaoUsergrowthOfflineConvertionSyncInfoGetRequest 初始化TaobaoUsergrowthOfflineConvertionSyncInfoGetAPIRequest对象
func NewTaobaoUsergrowthOfflineConvertionSyncInfoGetRequest() *TaobaoUsergrowthOfflineConvertionSyncInfoGetAPIRequest {
	return &TaobaoUsergrowthOfflineConvertionSyncInfoGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUsergrowthOfflineConvertionSyncInfoGetAPIRequest) GetApiMethodName() string {
	return "taobao.usergrowth.offline.convertion.sync.info.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUsergrowthOfflineConvertionSyncInfoGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetQuery is Query Setter
// 入参
func (r *TaobaoUsergrowthOfflineConvertionSyncInfoGetAPIRequest) SetQuery(_query *OfflineConvertionSyncInfoQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r TaobaoUsergrowthOfflineConvertionSyncInfoGetAPIRequest) GetQuery() *OfflineConvertionSyncInfoQuery {
	return r._query
}
