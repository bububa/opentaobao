package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除单品人群 API请求
taobao.feedflow.item.crowd.delete

删除单品人群
*/
type TaobaoFeedflowItemCrowdDeleteAPIRequest struct {
    model.Params
    // 人群结构
    _crowds   []CrowdDTO
    // 单元id
    _adgroupId   int64
}

// 初始化TaobaoFeedflowItemCrowdDeleteAPIRequest对象
func NewTaobaoFeedflowItemCrowdDeleteRequest() *TaobaoFeedflowItemCrowdDeleteAPIRequest{
    return &TaobaoFeedflowItemCrowdDeleteAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCrowdDeleteAPIRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.crowd.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCrowdDeleteAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Crowds Setter
// 人群结构
func (r *TaobaoFeedflowItemCrowdDeleteAPIRequest) SetCrowds(_crowds []CrowdDTO) error {
    r._crowds = _crowds
    r.Set("crowds", _crowds)
    return nil
}

// Crowds Getter
func (r TaobaoFeedflowItemCrowdDeleteAPIRequest) GetCrowds() []CrowdDTO {
    return r._crowds
}
// AdgroupId Setter
// 单元id
func (r *TaobaoFeedflowItemCrowdDeleteAPIRequest) SetAdgroupId(_adgroupId int64) error {
    r._adgroupId = _adgroupId
    r.Set("adgroup_id", _adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoFeedflowItemCrowdDeleteAPIRequest) GetAdgroupId() int64 {
    return r._adgroupId
}
