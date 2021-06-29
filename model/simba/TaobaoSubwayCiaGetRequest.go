package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询单元智能出价信息 API请求
taobao.subway.cia.get

查询单元智能出价信息
*/
type TaobaoSubwayCiaGetRequest struct {
    model.Params
    // 主人昵称
    nick   string
    // 推广组Id
    adgroupId   int64
}

// 初始化TaobaoSubwayCiaGetRequest对象
func NewTaobaoSubwayCiaGetRequest() *TaobaoSubwayCiaGetRequest{
    return &TaobaoSubwayCiaGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSubwayCiaGetRequest) GetApiMethodName() string {
    return "taobao.subway.cia.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSubwayCiaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 主人昵称
func (r *TaobaoSubwayCiaGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSubwayCiaGetRequest) GetNick() string {
    return r.nick
}
// AdgroupId Setter
// 推广组Id
func (r *TaobaoSubwayCiaGetRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSubwayCiaGetRequest) GetAdgroupId() int64 {
    return r.adgroupId
}
