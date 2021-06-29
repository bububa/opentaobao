package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询单元列表 API请求
taobao.feedflow.item.adgroup.page

通过计划id查询单元信息
*/
type TaobaoFeedflowItemAdgroupPageRequest struct {
    model.Params
    // 系统自动生成
    _adgroupQuery   *AdgroupQueryDTO
}

// 初始化TaobaoFeedflowItemAdgroupPageRequest对象
func NewTaobaoFeedflowItemAdgroupPageRequest() *TaobaoFeedflowItemAdgroupPageRequest{
    return &TaobaoFeedflowItemAdgroupPageRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemAdgroupPageRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.adgroup.page"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemAdgroupPageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AdgroupQuery Setter
// 系统自动生成
func (r *TaobaoFeedflowItemAdgroupPageRequest) SetAdgroupQuery(_adgroupQuery *AdgroupQueryDTO) error {
    r._adgroupQuery = _adgroupQuery
    r.Set("adgroup_query", _adgroupQuery)
    return nil
}

// AdgroupQuery Getter
func (r TaobaoFeedflowItemAdgroupPageRequest) GetAdgroupQuery() *AdgroupQueryDTO {
    return r._adgroupQuery
}
