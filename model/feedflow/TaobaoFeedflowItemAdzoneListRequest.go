package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量查询可用广告位列表 API请求
taobao.feedflow.item.adzone.list

批量查询可用广告位列表
*/
type TaobaoFeedflowItemAdzoneListRequest struct {
    model.Params
    // 广告位查询条件
    _adzoneQuery   *AdzoneQueryDto
}

// 初始化TaobaoFeedflowItemAdzoneListRequest对象
func NewTaobaoFeedflowItemAdzoneListRequest() *TaobaoFeedflowItemAdzoneListRequest{
    return &TaobaoFeedflowItemAdzoneListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemAdzoneListRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.adzone.list"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemAdzoneListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AdzoneQuery Setter
// 广告位查询条件
func (r *TaobaoFeedflowItemAdzoneListRequest) SetAdzoneQuery(_adzoneQuery *AdzoneQueryDto) error {
    r._adzoneQuery = _adzoneQuery
    r.Set("adzone_query", _adzoneQuery)
    return nil
}

// AdzoneQuery Getter
func (r TaobaoFeedflowItemAdzoneListRequest) GetAdzoneQuery() *AdzoneQueryDto {
    return r._adzoneQuery
}
