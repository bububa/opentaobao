package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
(新)销量明星删除推广单元接口 API请求
taobao.simba.salestar.adgroup.delete

删除一个推广组
*/
type TaobaoSimbaSalestarAdgroupDeleteRequest struct {
    model.Params
    // 主人昵称
    _nick   string
    // 推广组Id
    _adgroupId   int64
}

// 初始化TaobaoSimbaSalestarAdgroupDeleteRequest对象
func NewTaobaoSimbaSalestarAdgroupDeleteRequest() *TaobaoSimbaSalestarAdgroupDeleteRequest{
    return &TaobaoSimbaSalestarAdgroupDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaSalestarAdgroupDeleteRequest) GetApiMethodName() string {
    return "taobao.simba.salestar.adgroup.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaSalestarAdgroupDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaSalestarAdgroupDeleteRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaSalestarAdgroupDeleteRequest) GetNick() string {
    return r._nick
}
// AdgroupId Setter
// 推广组Id
func (r *TaobaoSimbaSalestarAdgroupDeleteRequest) SetAdgroupId(_adgroupId int64) error {
    r._adgroupId = _adgroupId
    r.Set("adgroup_id", _adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSimbaSalestarAdgroupDeleteRequest) GetAdgroupId() int64 {
    return r._adgroupId
}
