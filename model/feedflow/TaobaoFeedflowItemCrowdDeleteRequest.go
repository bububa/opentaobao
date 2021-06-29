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
type TaobaoFeedflowItemCrowdDeleteRequest struct {
    model.Params
    // 人群结构
    _crowds   []CrowdDTO
    // 单元id
    _adgroupId   int64
}

// 初始化TaobaoFeedflowItemCrowdDeleteRequest对象
func NewTaobaoFeedflowItemCrowdDeleteRequest() *TaobaoFeedflowItemCrowdDeleteRequest{
    return &TaobaoFeedflowItemCrowdDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCrowdDeleteRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.crowd.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCrowdDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Crowds Setter
// 人群结构
func (r *TaobaoFeedflowItemCrowdDeleteRequest) SetCrowds(_crowds []CrowdDTO) error {
    r._crowds = _crowds
    r.Set("crowds", _crowds)
    return nil
}

// Crowds Getter
func (r TaobaoFeedflowItemCrowdDeleteRequest) GetCrowds() []CrowdDTO {
    return r._crowds
}
// AdgroupId Setter
// 单元id
func (r *TaobaoFeedflowItemCrowdDeleteRequest) SetAdgroupId(_adgroupId int64) error {
    r._adgroupId = _adgroupId
    r.Set("adgroup_id", _adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoFeedflowItemCrowdDeleteRequest) GetAdgroupId() int64 {
    return r._adgroupId
}
