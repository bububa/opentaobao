package usergrowth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询 用增线下业务  转化数据是否同步完成 API请求
taobao.usergrowth.offline.convertion.sync.info.get

为手淘线下合作的渠道，提供对外查询数据是否更新完毕接口
*/
type TaobaoUsergrowthOfflineConvertionSyncInfoGetRequest struct {
    model.Params
    // 入参
    _query   *OfflineConvertionSyncInfoQuery
}

// 初始化TaobaoUsergrowthOfflineConvertionSyncInfoGetRequest对象
func NewTaobaoUsergrowthOfflineConvertionSyncInfoGetRequest() *TaobaoUsergrowthOfflineConvertionSyncInfoGetRequest{
    return &TaobaoUsergrowthOfflineConvertionSyncInfoGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUsergrowthOfflineConvertionSyncInfoGetRequest) GetApiMethodName() string {
    return "taobao.usergrowth.offline.convertion.sync.info.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUsergrowthOfflineConvertionSyncInfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Query Setter
// 入参
func (r *TaobaoUsergrowthOfflineConvertionSyncInfoGetRequest) SetQuery(_query *OfflineConvertionSyncInfoQuery) error {
    r._query = _query
    r.Set("query", _query)
    return nil
}

// Query Getter
func (r TaobaoUsergrowthOfflineConvertionSyncInfoGetRequest) GetQuery() *OfflineConvertionSyncInfoQuery {
    return r._query
}
