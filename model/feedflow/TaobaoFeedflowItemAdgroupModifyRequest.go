package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
信息流单元修改 API请求
taobao.feedflow.item.adgroup.modify

信息流单元修改
*/
type TaobaoFeedflowItemAdgroupModifyAPIRequest struct {
    model.Params
    // 单元信息
    _adgroup   *AdgroupDTO
}

// 初始化TaobaoFeedflowItemAdgroupModifyAPIRequest对象
func NewTaobaoFeedflowItemAdgroupModifyRequest() *TaobaoFeedflowItemAdgroupModifyAPIRequest{
    return &TaobaoFeedflowItemAdgroupModifyAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemAdgroupModifyAPIRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.adgroup.modify"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemAdgroupModifyAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Adgroup Setter
// 单元信息
func (r *TaobaoFeedflowItemAdgroupModifyAPIRequest) SetAdgroup(_adgroup *AdgroupDTO) error {
    r._adgroup = _adgroup
    r.Set("adgroup", _adgroup)
    return nil
}

// Adgroup Getter
func (r TaobaoFeedflowItemAdgroupModifyAPIRequest) GetAdgroup() *AdgroupDTO {
    return r._adgroup
}
