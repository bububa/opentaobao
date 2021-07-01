package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
信息流增加单元 API请求
taobao.feedflow.item.adgroup.add

信息流增加单元
*/
type TaobaoFeedflowItemAdgroupAddAPIRequest struct {
    model.Params
    // 单元信息
    _adgroup   *AdgroupDTO
}

// 初始化TaobaoFeedflowItemAdgroupAddAPIRequest对象
func NewTaobaoFeedflowItemAdgroupAddRequest() *TaobaoFeedflowItemAdgroupAddAPIRequest{
    return &TaobaoFeedflowItemAdgroupAddAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemAdgroupAddAPIRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.adgroup.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemAdgroupAddAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Adgroup Setter
// 单元信息
func (r *TaobaoFeedflowItemAdgroupAddAPIRequest) SetAdgroup(_adgroup *AdgroupDTO) error {
    r._adgroup = _adgroup
    r.Set("adgroup", _adgroup)
    return nil
}

// Adgroup Getter
func (r TaobaoFeedflowItemAdgroupAddAPIRequest) GetAdgroup() *AdgroupDTO {
    return r._adgroup
}
