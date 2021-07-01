package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据推广单元id获取搜索溢价人群 API请求
taobao.simba.serchcrowd.get

根据推广单元id获取搜索溢价人群
*/
type TaobaoSimbaSerchcrowdGetAPIRequest struct {
    model.Params
    // 被操作者的淘宝昵称
    _nick   string
    // 推广单元id
    _adgroupId   int64
}

// 初始化TaobaoSimbaSerchcrowdGetAPIRequest对象
func NewTaobaoSimbaSerchcrowdGetRequest() *TaobaoSimbaSerchcrowdGetAPIRequest{
    return &TaobaoSimbaSerchcrowdGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaSerchcrowdGetAPIRequest) GetApiMethodName() string {
    return "taobao.simba.serchcrowd.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaSerchcrowdGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 被操作者的淘宝昵称
func (r *TaobaoSimbaSerchcrowdGetAPIRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaSerchcrowdGetAPIRequest) GetNick() string {
    return r._nick
}
// AdgroupId Setter
// 推广单元id
func (r *TaobaoSimbaSerchcrowdGetAPIRequest) SetAdgroupId(_adgroupId int64) error {
    r._adgroupId = _adgroupId
    r.Set("adgroup_id", _adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSimbaSerchcrowdGetAPIRequest) GetAdgroupId() int64 {
    return r._adgroupId
}
