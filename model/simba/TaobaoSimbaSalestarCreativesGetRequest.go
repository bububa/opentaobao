package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
（新）批量获取创意 API请求
taobao.simba.salestar.creatives.get

取得一个推广组的所有创意或者根据一个创意Id列表取得一组创意；<br/>如果同时提供了推广组Id和创意id列表，则优先使用推广组Id；
*/
type TaobaoSimbaSalestarCreativesGetRequest struct {
    model.Params
    // 主人昵称
    _nick   string
    // 创意Id数组，最多200个
    _creativeIds   []int64
    // 推广组Id
    _adgroupId   int64
}

// 初始化TaobaoSimbaSalestarCreativesGetRequest对象
func NewTaobaoSimbaSalestarCreativesGetRequest() *TaobaoSimbaSalestarCreativesGetRequest{
    return &TaobaoSimbaSalestarCreativesGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaSalestarCreativesGetRequest) GetApiMethodName() string {
    return "taobao.simba.salestar.creatives.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaSalestarCreativesGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaSalestarCreativesGetRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaSalestarCreativesGetRequest) GetNick() string {
    return r._nick
}
// CreativeIds Setter
// 创意Id数组，最多200个
func (r *TaobaoSimbaSalestarCreativesGetRequest) SetCreativeIds(_creativeIds []int64) error {
    r._creativeIds = _creativeIds
    r.Set("creative_ids", _creativeIds)
    return nil
}

// CreativeIds Getter
func (r TaobaoSimbaSalestarCreativesGetRequest) GetCreativeIds() []int64 {
    return r._creativeIds
}
// AdgroupId Setter
// 推广组Id
func (r *TaobaoSimbaSalestarCreativesGetRequest) SetAdgroupId(_adgroupId int64) error {
    r._adgroupId = _adgroupId
    r.Set("adgroup_id", _adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSimbaSalestarCreativesGetRequest) GetAdgroupId() int64 {
    return r._adgroupId
}
