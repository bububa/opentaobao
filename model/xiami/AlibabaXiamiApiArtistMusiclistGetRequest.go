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
type AlibabaXiamiApiArtistMusiclistGetAPIRequest struct {
    model.Params
    // 语种, 有all, chinese, musician, english, japanese, korea
    _type   string
    // 所有、男、女、组合分别为(all、male、female、combination)
    _order   string
}

// 初始化AlibabaXiamiApiArtistMusiclistGetAPIRequest对象
func NewAlibabaXiamiApiArtistMusiclistGetRequest() *AlibabaXiamiApiArtistMusiclistGetAPIRequest{
    return &AlibabaXiamiApiArtistMusiclistGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaXiamiApiArtistMusiclistGetAPIRequest) GetApiMethodName() string {
    return "alibaba.xiami.api.artist.musiclist.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaXiamiApiArtistMusiclistGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Type Setter
// 语种, 有all, chinese, musician, english, japanese, korea
func (r *AlibabaXiamiApiArtistMusiclistGetAPIRequest) SetType(_type string) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r AlibabaXiamiApiArtistMusiclistGetAPIRequest) GetType() string {
    return r._type
}
// Order Setter
// 所有、男、女、组合分别为(all、male、female、combination)
func (r *AlibabaXiamiApiArtistMusiclistGetAPIRequest) SetOrder(_order string) error {
    r._order = _order
    r.Set("order", _order)
    return nil
}

// Order Getter
func (r AlibabaXiamiApiArtistMusiclistGetAPIRequest) GetOrder() string {
    return r._order
}
