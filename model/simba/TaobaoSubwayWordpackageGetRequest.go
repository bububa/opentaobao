package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取词包列表 API请求
taobao.subway.wordpackage.get

获取流量智选、捡漏词包等词包列表
*/
type TaobaoSubwayWordpackageGetRequest struct {
    model.Params
    // 主人昵称
    nick   string
    // 推广组id
    adgroupId   int64
}

// 初始化TaobaoSubwayWordpackageGetRequest对象
func NewTaobaoSubwayWordpackageGetRequest() *TaobaoSubwayWordpackageGetRequest{
    return &TaobaoSubwayWordpackageGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSubwayWordpackageGetRequest) GetApiMethodName() string {
    return "taobao.subway.wordpackage.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSubwayWordpackageGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 主人昵称
func (r *TaobaoSubwayWordpackageGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSubwayWordpackageGetRequest) GetNick() string {
    return r.nick
}
// AdgroupId Setter
// 推广组id
func (r *TaobaoSubwayWordpackageGetRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSubwayWordpackageGetRequest) GetAdgroupId() int64 {
    return r.adgroupId
}
