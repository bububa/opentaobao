package xiami

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
排行榜歌曲获取 API请求
alibaba.xiami.api.rank.songs.get

获取歌曲排行榜
*/
type AlibabaXiamiApiRankSongsGetRequest struct {
    model.Params
    // 榜单类型:<br/>虾米榜 music_all,music_oumei,music_ri,music_han,music_huayu;<br/>虾米新歌榜 newmusic_all,newmusc_oumei,newmusic_ri,newmusic_han,newmusic_huayu;<br/>全球媒体榜 hito,jingge,uk,billboard,oricon,mnet;<br/>原创榜 music_original;<br/>demo榜 music_demo;<br/>陌陌榜 momo;
    type   string
}

// 初始化AlibabaXiamiApiRankSongsGetRequest对象
func NewAlibabaXiamiApiRankSongsGetRequest() *AlibabaXiamiApiRankSongsGetRequest{
    return &AlibabaXiamiApiRankSongsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaXiamiApiRankSongsGetRequest) GetApiMethodName() string {
    return "alibaba.xiami.api.rank.songs.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaXiamiApiRankSongsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Type Setter
// 榜单类型:<br/>虾米榜 music_all,music_oumei,music_ri,music_han,music_huayu;<br/>虾米新歌榜 newmusic_all,newmusc_oumei,newmusic_ri,newmusic_han,newmusic_huayu;<br/>全球媒体榜 hito,jingge,uk,billboard,oricon,mnet;<br/>原创榜 music_original;<br/>demo榜 music_demo;<br/>陌陌榜 momo;
func (r *AlibabaXiamiApiRankSongsGetRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r AlibabaXiamiApiRankSongsGetRequest) GetType() string {
    return r.type
}
