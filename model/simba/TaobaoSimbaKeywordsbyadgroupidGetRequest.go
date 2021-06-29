package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取得一个推广组的所有关键词 API请求
taobao.simba.keywordsbyadgroupid.get

取得一个推广组的所有关键词
*/
type TaobaoSimbaKeywordsbyadgroupidGetRequest struct {
    model.Params
    // 主人昵称
    _nick   string
    // 推广组Id
    _adgroupId   int64
}

// 初始化TaobaoSimbaKeywordsbyadgroupidGetRequest对象
func NewTaobaoSimbaKeywordsbyadgroupidGetRequest() *TaobaoSimbaKeywordsbyadgroupidGetRequest{
    return &TaobaoSimbaKeywordsbyadgroupidGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaKeywordsbyadgroupidGetRequest) GetApiMethodName() string {
    return "taobao.simba.keywordsbyadgroupid.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaKeywordsbyadgroupidGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaKeywordsbyadgroupidGetRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaKeywordsbyadgroupidGetRequest) GetNick() string {
    return r._nick
}
// AdgroupId Setter
// 推广组Id
func (r *TaobaoSimbaKeywordsbyadgroupidGetRequest) SetAdgroupId(_adgroupId int64) error {
    r._adgroupId = _adgroupId
    r.Set("adgroup_id", _adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSimbaKeywordsbyadgroupidGetRequest) GetAdgroupId() int64 {
    return r._adgroupId
}
