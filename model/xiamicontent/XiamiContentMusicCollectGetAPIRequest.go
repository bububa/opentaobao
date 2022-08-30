package xiamicontent

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// XiamiContentMusicCollectGetAPIRequest 获取歌单歌曲 API请求
// xiami.content.music.collect.get
//
// 获取歌单歌曲
type XiamiContentMusicCollectGetAPIRequest struct {
	model.Params
	// 歌单id
	_collectId int64
	// 每页数量
	_pageSize int64
	// 页码
	_page int64
}

// NewXiamiContentMusicCollectGetRequest 初始化XiamiContentMusicCollectGetAPIRequest对象
func NewXiamiContentMusicCollectGetRequest() *XiamiContentMusicCollectGetAPIRequest {
	return &XiamiContentMusicCollectGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r XiamiContentMusicCollectGetAPIRequest) GetApiMethodName() string {
	return "xiami.content.music.collect.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r XiamiContentMusicCollectGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCollectId is CollectId Setter
// 歌单id
func (r *XiamiContentMusicCollectGetAPIRequest) SetCollectId(_collectId int64) error {
	r._collectId = _collectId
	r.Set("collect_id", _collectId)
	return nil
}

// GetCollectId CollectId Getter
func (r XiamiContentMusicCollectGetAPIRequest) GetCollectId() int64 {
	return r._collectId
}

// SetPageSize is PageSize Setter
// 每页数量
func (r *XiamiContentMusicCollectGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r XiamiContentMusicCollectGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPage is Page Setter
// 页码
func (r *XiamiContentMusicCollectGetAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r XiamiContentMusicCollectGetAPIRequest) GetPage() int64 {
	return r._page
}
