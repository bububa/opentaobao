package xiami

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
热门艺人 API请求
alibaba.xiami.api.artist.musiclist.get

热门艺人
*/
type AlibabaXiamiApiArtistMusiclistGetRequest struct {
    model.Params
    // 语种, 有all, chinese, musician, english, japanese, korea
    type   string
    // 所有、男、女、组合分别为(all、male、female、combination)
    order   string
}

// 初始化AlibabaXiamiApiArtistMusiclistGetRequest对象
func NewAlibabaXiamiApiArtistMusiclistGetRequest() *AlibabaXiamiApiArtistMusiclistGetRequest{
    return &AlibabaXiamiApiArtistMusiclistGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaXiamiApiArtistMusiclistGetRequest) GetApiMethodName() string {
    return "alibaba.xiami.api.artist.musiclist.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaXiamiApiArtistMusiclistGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Type Setter
// 语种, 有all, chinese, musician, english, japanese, korea
func (r *AlibabaXiamiApiArtistMusiclistGetRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r AlibabaXiamiApiArtistMusiclistGetRequest) GetType() string {
    return r.type
}
// Order Setter
// 所有、男、女、组合分别为(all、male、female、combination)
func (r *AlibabaXiamiApiArtistMusiclistGetRequest) SetOrder(order string) error {
    r.order = order
    r.Set("order", order)
    return nil
}

// Order Getter
func (r AlibabaXiamiApiArtistMusiclistGetRequest) GetOrder() string {
    return r.order
}
