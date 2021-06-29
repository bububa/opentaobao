package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单品单元下，新增定向人群 API请求
taobao.feedflow.item.crowd.add

单品单元下，新增定向人群
*/
type TaobaoFeedflowItemCrowdAddRequest struct {
    model.Params
    // 人群列表
    _crowds   []CrowdDTO
    // 单元id
    _adgroupId   int64
}

// 初始化TaobaoFeedflowItemCrowdAddRequest对象
func NewTaobaoFeedflowItemCrowdAddRequest() *TaobaoFeedflowItemCrowdAddRequest{
    return &TaobaoFeedflowItemCrowdAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCrowdAddRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.crowd.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCrowdAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Crowds Setter
// 人群列表
func (r *TaobaoFeedflowItemCrowdAddRequest) SetCrowds(_crowds []CrowdDTO) error {
    r._crowds = _crowds
    r.Set("crowds", _crowds)
    return nil
}

// Crowds Getter
func (r TaobaoFeedflowItemCrowdAddRequest) GetCrowds() []CrowdDTO {
    return r._crowds
}
// AdgroupId Setter
// 单元id
func (r *TaobaoFeedflowItemCrowdAddRequest) SetAdgroupId(_adgroupId int64) error {
    r._adgroupId = _adgroupId
    r.Set("adgroup_id", _adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoFeedflowItemCrowdAddRequest) GetAdgroupId() int64 {
    return r._adgroupId
}
