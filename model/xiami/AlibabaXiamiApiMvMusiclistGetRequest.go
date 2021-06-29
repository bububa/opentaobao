package xiami

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
乐馆mv列表 API请求
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

// 初始化AlibabaXiamiApiMvMusiclistGetRequest对象
func NewAlibabaXiamiApiMvMusiclistGetRequest() *AlibabaXiamiApiMvMusiclistGetRequest{
    return &AlibabaXiamiApiMvMusiclistGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaXiamiApiMvMusiclistGetRequest) GetApiMethodName() string {
    return "alibaba.xiami.api.mv.musiclist.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaXiamiApiMvMusiclistGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Type Setter
// 语种, 有all, chinese, musician, english, japanese, korea
func (r *AlibabaXiamiApiMvMusiclistGetRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r AlibabaXiamiApiMvMusiclistGetRequest) GetType() string {
    return r.type
}
// Order Setter
// 分组, 有recommend、hot、new
func (r *AlibabaXiamiApiMvMusiclistGetRequest) SetOrder(order string) error {
    r.order = order
    r.Set("order", order)
    return nil
}

// Order Getter
func (r AlibabaXiamiApiMvMusiclistGetRequest) GetOrder() string {
    return r.order
}
// Limit Setter
// 每页记录
func (r *AlibabaXiamiApiMvMusiclistGetRequest) SetLimit(limit int64) error {
    r.limit = limit
    r.Set("limit", limit)
    return nil
}

// Limit Getter
func (r AlibabaXiamiApiMvMusiclistGetRequest) GetLimit() int64 {
    return r.limit
}
// Page Setter
// 当前页
func (r *AlibabaXiamiApiMvMusiclistGetRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

// Page Getter
func (r AlibabaXiamiApiMvMusiclistGetRequest) GetPage() int64 {
    return r.page
}
