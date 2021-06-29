package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
信息流单元下查看绑定资源位 API请求
taobao.feedflow.item.adgroup.adzone.page

信息流单元下查看绑定资源位
*/
type TaobaoFeedflowItemAdgroupAdzonePageRequest struct {
    model.Params
    // 查询条件
    _adzoneBindQuery   *AdzoneBindQueryDTO
}

// 初始化TaobaoFeedflowItemAdgroupAdzonePageRequest对象
func NewTaobaoFeedflowItemAdgroupAdzonePageRequest() *TaobaoFeedflowItemAdgroupAdzonePageRequest{
    return &TaobaoFeedflowItemAdgroupAdzonePageRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemAdgroupAdzonePageRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.adgroup.adzone.page"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemAdgroupAdzonePageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AdzoneBindQuery Setter
// 查询条件
func (r *TaobaoFeedflowItemAdgroupAdzonePageRequest) SetAdzoneBindQuery(_adzoneBindQuery *AdzoneBindQueryDTO) error {
    r._adzoneBindQuery = _adzoneBindQuery
    r.Set("adzone_bind_query", _adzoneBindQuery)
    return nil
}

// AdzoneBindQuery Getter
func (r TaobaoFeedflowItemAdgroupAdzonePageRequest) GetAdzoneBindQuery() *AdzoneBindQueryDTO {
    return r._adzoneBindQuery
}
