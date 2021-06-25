package xiami

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
乐馆mv列表 APIRequest
alibaba.xiami.api.mv.musiclist.get

乐馆mv列表
*/
type AlibabaXiamiApiMvMusiclistGetRequest struct {
    model.Params

    // 语种, 有all, chinese, musician, english, japanese, korea
    type   string 

    // 分组, 有recommend、hot、new
    order   string 

    // 每页记录
    limit   int64 

    // 当前页
    page   int64 

}

func NewAlibabaXiamiApiMvMusiclistGetRequest() *AlibabaXiamiApiMvMusiclistGetRequest{
    return &AlibabaXiamiApiMvMusiclistGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaXiamiApiMvMusiclistGetRequest) GetApiMethodName() string {
    return "alibaba.xiami.api.mv.musiclist.get"
}

func (r AlibabaXiamiApiMvMusiclistGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaXiamiApiMvMusiclistGetRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r AlibabaXiamiApiMvMusiclistGetRequest) GetType() string {
    return r.type
}

func (r *AlibabaXiamiApiMvMusiclistGetRequest) SetOrder(order string) error {
    r.order = order
    r.Set("order", order)
    return nil
}

func (r AlibabaXiamiApiMvMusiclistGetRequest) GetOrder() string {
    return r.order
}

func (r *AlibabaXiamiApiMvMusiclistGetRequest) SetLimit(limit int64) error {
    r.limit = limit
    r.Set("limit", limit)
    return nil
}

func (r AlibabaXiamiApiMvMusiclistGetRequest) GetLimit() int64 {
    return r.limit
}

func (r *AlibabaXiamiApiMvMusiclistGetRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

func (r AlibabaXiamiApiMvMusiclistGetRequest) GetPage() int64 {
    return r.page
}

