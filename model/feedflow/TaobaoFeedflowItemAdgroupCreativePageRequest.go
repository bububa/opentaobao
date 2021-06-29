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
type TaobaoFeedflowItemAdgroupCreativePageRequest struct {
    model.Params
    // 绑定查询条件
    creativeBindQuery   *CreativeBindQueryDto
}

// 初始化TaobaoFeedflowItemAdgroupCreativePageRequest对象
func NewTaobaoFeedflowItemAdgroupCreativePageRequest() *TaobaoFeedflowItemAdgroupCreativePageRequest{
    return &TaobaoFeedflowItemAdgroupCreativePageRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemAdgroupCreativePageRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.adgroup.creative.page"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemAdgroupCreativePageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CreativeBindQuery Setter
// 绑定查询条件
func (r *TaobaoFeedflowItemAdgroupCreativePageRequest) SetCreativeBindQuery(creativeBindQuery *CreativeBindQueryDto) error {
    r.creativeBindQuery = creativeBindQuery
    r.Set("creative_bind_query", creativeBindQuery)
    return nil
}

// CreativeBindQuery Getter
func (r TaobaoFeedflowItemAdgroupCreativePageRequest) GetCreativeBindQuery() *CreativeBindQueryDto {
    return r.creativeBindQuery
}
