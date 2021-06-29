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
type TaobaoFeedflowItemAdgroupAddRequest struct {
    model.Params
    // 单元信息
    _adgroup   *AdgroupDto
}

// 初始化TaobaoFeedflowItemAdgroupAddRequest对象
func NewTaobaoFeedflowItemAdgroupAddRequest() *TaobaoFeedflowItemAdgroupAddRequest{
    return &TaobaoFeedflowItemAdgroupAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemAdgroupAddRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.adgroup.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemAdgroupAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Adgroup Setter
// 单元信息
func (r *TaobaoFeedflowItemAdgroupAddRequest) SetAdgroup(_adgroup *AdgroupDto) error {
    r._adgroup = _adgroup
    r.Set("adgroup", _adgroup)
    return nil
}

// Adgroup Getter
func (r TaobaoFeedflowItemAdgroupAddRequest) GetAdgroup() *AdgroupDto {
    return r._adgroup
}
