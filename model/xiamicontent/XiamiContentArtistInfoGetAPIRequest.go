package xiamicontent

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// XiamiContentArtistInfoGetAPIRequest 获取艺人信息 API请求
// xiami.content.artist.info.get
//
// (批量)获取艺人信息
type XiamiContentArtistInfoGetAPIRequest struct {
	model.Params
	// 艺人id
	_artistIds int64
}

// NewXiamiContentArtistInfoGetRequest 初始化XiamiContentArtistInfoGetAPIRequest对象
func NewXiamiContentArtistInfoGetRequest() *XiamiContentArtistInfoGetAPIRequest {
	return &XiamiContentArtistInfoGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r XiamiContentArtistInfoGetAPIRequest) GetApiMethodName() string {
	return "xiami.content.artist.info.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r XiamiContentArtistInfoGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetArtistIds is ArtistIds Setter
// 艺人id
func (r *XiamiContentArtistInfoGetAPIRequest) SetArtistIds(_artistIds int64) error {
	r._artistIds = _artistIds
	r.Set("artist_ids", _artistIds)
	return nil
}

// GetArtistIds ArtistIds Getter
func (r XiamiContentArtistInfoGetAPIRequest) GetArtistIds() int64 {
	return r._artistIds
}
