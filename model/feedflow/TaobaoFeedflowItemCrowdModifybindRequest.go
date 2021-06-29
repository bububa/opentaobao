package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改人群出价或状态 API请求
taobao.feedflow.item.crowd.modifybind

修改人群出价或状态
*/
type TaobaoFeedflowItemCrowdModifybindRequest struct {
    model.Params
    // 人群信息
    _crowds   []CrowdDto
    // 单元id
    _adgroupId   int64
}

// 初始化TaobaoFeedflowItemCrowdModifybindRequest对象
func NewTaobaoFeedflowItemCrowdModifybindRequest() *TaobaoFeedflowItemCrowdModifybindRequest{
    return &TaobaoFeedflowItemCrowdModifybindRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCrowdModifybindRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.crowd.modifybind"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCrowdModifybindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Crowds Setter
// 人群信息
func (r *TaobaoFeedflowItemCrowdModifybindRequest) SetCrowds(_crowds []CrowdDto) error {
    r._crowds = _crowds
    r.Set("crowds", _crowds)
    return nil
}

// Crowds Getter
func (r TaobaoFeedflowItemCrowdModifybindRequest) GetCrowds() []CrowdDto {
    return r._crowds
}
// AdgroupId Setter
// 单元id
func (r *TaobaoFeedflowItemCrowdModifybindRequest) SetAdgroupId(_adgroupId int64) error {
    r._adgroupId = _adgroupId
    r.Set("adgroup_id", _adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoFeedflowItemCrowdModifybindRequest) GetAdgroupId() int64 {
    return r._adgroupId
}
