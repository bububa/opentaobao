package xiami

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
热门艺人 APIRequest
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

func NewAlibabaXiamiApiArtistMusiclistGetRequest() *AlibabaXiamiApiArtistMusiclistGetRequest{
    return &AlibabaXiamiApiArtistMusiclistGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaXiamiApiArtistMusiclistGetRequest) GetApiMethodName() string {
    return "alibaba.xiami.api.artist.musiclist.get"
}

func (r AlibabaXiamiApiArtistMusiclistGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaXiamiApiArtistMusiclistGetRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r AlibabaXiamiApiArtistMusiclistGetRequest) GetType() string {
    return r.type
}

func (r *AlibabaXiamiApiArtistMusiclistGetRequest) SetOrder(order string) error {
    r.order = order
    r.Set("order", order)
    return nil
}

func (r AlibabaXiamiApiArtistMusiclistGetRequest) GetOrder() string {
    return r.order
}

