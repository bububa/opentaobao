package xiami

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
排行榜歌曲获取 APIRequest
alibaba.xiami.api.rank.songs.get

获取歌曲排行榜
*/
type AlibabaXiamiApiRankSongsGetRequest struct {
    model.Params

    // 榜单类型:<br/>虾米榜 music_all,music_oumei,music_ri,music_han,music_huayu;<br/>虾米新歌榜 newmusic_all,newmusc_oumei,newmusic_ri,newmusic_han,newmusic_huayu;<br/>全球媒体榜 hito,jingge,uk,billboard,oricon,mnet;<br/>原创榜 music_original;<br/>demo榜 music_demo;<br/>陌陌榜 momo;
    type   string 

}

func NewAlibabaXiamiApiRankSongsGetRequest() *AlibabaXiamiApiRankSongsGetRequest{
    return &AlibabaXiamiApiRankSongsGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaXiamiApiRankSongsGetRequest) GetApiMethodName() string {
    return "alibaba.xiami.api.rank.songs.get"
}

func (r AlibabaXiamiApiRankSongsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaXiamiApiRankSongsGetRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r AlibabaXiamiApiRankSongsGetRequest) GetType() string {
    return r.type
}

