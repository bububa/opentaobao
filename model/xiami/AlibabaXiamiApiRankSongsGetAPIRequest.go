package xiami

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaXiamiApiRankSongsGetAPIRequest 排行榜歌曲获取 API请求
// alibaba.xiami.api.rank.songs.get
//
// 获取歌曲排行榜
type AlibabaXiamiApiRankSongsGetAPIRequest struct {
	model.Params
	// 榜单类型:<br/>虾米榜 music_all,music_oumei,music_ri,music_han,music_huayu;<br/>虾米新歌榜 newmusic_all,newmusc_oumei,newmusic_ri,newmusic_han,newmusic_huayu;<br/>全球媒体榜 hito,jingge,uk,billboard,oricon,mnet;<br/>原创榜 music_original;<br/>demo榜 music_demo;<br/>陌陌榜 momo;
	_type string
}

// NewAlibabaXiamiApiRankSongsGetRequest 初始化AlibabaXiamiApiRankSongsGetAPIRequest对象
func NewAlibabaXiamiApiRankSongsGetRequest() *AlibabaXiamiApiRankSongsGetAPIRequest {
	return &AlibabaXiamiApiRankSongsGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaXiamiApiRankSongsGetAPIRequest) GetApiMethodName() string {
	return "alibaba.xiami.api.rank.songs.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaXiamiApiRankSongsGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetType is Type Setter
// 榜单类型:<br/>虾米榜 music_all,music_oumei,music_ri,music_han,music_huayu;<br/>虾米新歌榜 newmusic_all,newmusc_oumei,newmusic_ri,newmusic_han,newmusic_huayu;<br/>全球媒体榜 hito,jingge,uk,billboard,oricon,mnet;<br/>原创榜 music_original;<br/>demo榜 music_demo;<br/>陌陌榜 momo;
func (r *AlibabaXiamiApiRankSongsGetAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaXiamiApiRankSongsGetAPIRequest) GetType() string {
	return r._type
}
