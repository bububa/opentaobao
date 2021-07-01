package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
信息流单元下查看创意 API请求
taobao.feedflow.item.adgroup.creative.page

信息流单元下查看创意
*/
type TaobaoFeedflowItemAdgroupCreativePageAPIRequest struct {
    model.Params
    // 绑定查询条件
    _creativeBindQuery   *CreativeBindQueryDTO
}

// 初始化TaobaoFeedflowItemAdgroupCreativePageAPIRequest对象
func NewTaobaoFeedflowItemAdgroupCreativePageRequest() *TaobaoFeedflowItemAdgroupCreativePageAPIRequest{
    return &TaobaoFeedflowItemAdgroupCreativePageAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemAdgroupCreativePageAPIRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.adgroup.creative.page"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemAdgroupCreativePageAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CreativeBindQuery Setter
// 绑定查询条件
func (r *TaobaoFeedflowItemAdgroupCreativePageAPIRequest) SetCreativeBindQuery(_creativeBindQuery *CreativeBindQueryDTO) error {
    r._creativeBindQuery = _creativeBindQuery
    r.Set("creative_bind_query", _creativeBindQuery)
    return nil
}

// CreativeBindQuery Getter
func (r TaobaoFeedflowItemAdgroupCreativePageAPIRequest) GetCreativeBindQuery() *CreativeBindQueryDTO {
    return r._creativeBindQuery
}
