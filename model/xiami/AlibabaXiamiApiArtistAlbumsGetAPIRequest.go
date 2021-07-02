package xiami

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaXiamiApiArtistAlbumsGetAPIRequest 艺人专辑 API请求
// alibaba.xiami.api.artist.albums.get
//
// 艺人专辑
type AlibabaXiamiApiArtistAlbumsGetAPIRequest struct {
	model.Params
	// 歌曲数量
	_limit int64
	// 页码
	_page int64
	// 艺人id
	_id int64
}

// NewAlibabaXiamiApiArtistAlbumsGetRequest 初始化AlibabaXiamiApiArtistAlbumsGetAPIRequest对象
func NewAlibabaXiamiApiArtistAlbumsGetRequest() *AlibabaXiamiApiArtistAlbumsGetAPIRequest {
	return &AlibabaXiamiApiArtistAlbumsGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaXiamiApiArtistAlbumsGetAPIRequest) GetApiMethodName() string {
	return "alibaba.xiami.api.artist.albums.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaXiamiApiArtistAlbumsGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Limit Setter
// 歌曲数量
func (r *AlibabaXiamiApiArtistAlbumsGetAPIRequest) SetLimit(_limit int64) error {
	r._limit = _limit
	r.Set("limit", _limit)
	return nil
}

// Get Limit Getter
func (r AlibabaXiamiApiArtistAlbumsGetAPIRequest) GetLimit() int64 {
	return r._limit
}

// Set is Page Setter
// 页码
func (r *AlibabaXiamiApiArtistAlbumsGetAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// Get Page Getter
func (r AlibabaXiamiApiArtistAlbumsGetAPIRequest) GetPage() int64 {
	return r._page
}

// Set is Id Setter
// 艺人id
func (r *AlibabaXiamiApiArtistAlbumsGetAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// Get Id Getter
func (r AlibabaXiamiApiArtistAlbumsGetAPIRequest) GetId() int64 {
	return r._id
}
