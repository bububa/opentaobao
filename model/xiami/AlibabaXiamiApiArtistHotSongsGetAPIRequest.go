package xiami

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaXiamiApiArtistHotSongsGetAPIRequest
热门歌曲 API请求
alibaba.xiami.api.artist.hotSongs.get

热门歌曲 */
type AlibabaXiamiApiArtistHotSongsGetAPIRequest struct {
	model.Params
	// 艺人id
	_id int64
}

// NewAlibabaXiamiApiArtistHotSongsGetRequest 初始化AlibabaXiamiApiArtistHotSongsGetAPIRequest对象
func NewAlibabaXiamiApiArtistHotSongsGetRequest() *AlibabaXiamiApiArtistHotSongsGetAPIRequest {
	return &AlibabaXiamiApiArtistHotSongsGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaXiamiApiArtistHotSongsGetAPIRequest) GetApiMethodName() string {
	return "alibaba.xiami.api.artist.hotSongs.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaXiamiApiArtistHotSongsGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Id Setter
// 艺人id
func (r *AlibabaXiamiApiArtistHotSongsGetAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// Get Id Getter
func (r AlibabaXiamiApiArtistHotSongsGetAPIRequest) GetId() int64 {
	return r._id
}
