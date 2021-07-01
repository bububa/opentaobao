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
type TaobaoFeedflowItemCrowdPageAPIRequest struct {
    model.Params
    // 查询条件
    _crowdQuery   *CrowdQueryDto
}

// 初始化TaobaoFeedflowItemCrowdPageAPIRequest对象
func NewTaobaoFeedflowItemCrowdPageRequest() *TaobaoFeedflowItemCrowdPageAPIRequest{
    return &TaobaoFeedflowItemCrowdPageAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCrowdPageAPIRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.crowd.page"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCrowdPageAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CrowdQuery Setter
// 查询条件
func (r *TaobaoFeedflowItemCrowdPageAPIRequest) SetCrowdQuery(_crowdQuery *CrowdQueryDto) error {
    r._crowdQuery = _crowdQuery
    r.Set("crowd_query", _crowdQuery)
    return nil
}

// CrowdQuery Getter
func (r TaobaoFeedflowItemCrowdPageAPIRequest) GetCrowdQuery() *CrowdQueryDto {
    return r._crowdQuery
}
