package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询单品单元下人群列表 API请求
taobao.feedflow.item.crowd.page

分页查询单品单元下人群列表
*/
type TaobaoFeedflowItemCrowdPageRequest struct {
    model.Params
    // 查询条件
    _crowdQuery   *CrowdQueryDTO
}

// 初始化TaobaoFeedflowItemCrowdPageRequest对象
func NewTaobaoFeedflowItemCrowdPageRequest() *TaobaoFeedflowItemCrowdPageRequest{
    return &TaobaoFeedflowItemCrowdPageRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCrowdPageRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.crowd.page"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCrowdPageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CrowdQuery Setter
// 查询条件
func (r *TaobaoFeedflowItemCrowdPageRequest) SetCrowdQuery(_crowdQuery *CrowdQueryDTO) error {
    r._crowdQuery = _crowdQuery
    r.Set("crowd_query", _crowdQuery)
    return nil
}

// CrowdQuery Getter
func (r TaobaoFeedflowItemCrowdPageRequest) GetCrowdQuery() *CrowdQueryDTO {
    return r._crowdQuery
}
