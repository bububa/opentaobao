package xiamicontent

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// XiamicontentmusiccollectgetAPIRequest 获取歌单歌曲 API请求
// xiami.content.music.collect.get
//
// 获取歌单歌曲
type XiamicontentmusiccollectgetAPIRequest struct {
	model.Params
	// 歌单id
	_collectId int64
	// 每页数量
	_pageSize int64
	// 页码
	_page int64
}

// NewXiamicontentmusiccollectgetRequest 初始化XiamicontentmusiccollectgetAPIRequest对象
func NewXiamicontentmusiccollectgetRequest() *XiamicontentmusiccollectgetAPIRequest {
	return &XiamicontentmusiccollectgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r XiamicontentmusiccollectgetAPIRequest) GetApiMethodName() string {
	return "xiami.content.music.collect.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r XiamicontentmusiccollectgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r XiamicontentmusiccollectgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCollectId is CollectId Setter
// 歌单id
func (r *XiamicontentmusiccollectgetAPIRequest) SetCollectId(_collectId int64) error {
	r._collectId = _collectId
	r.Set("collect_id", _collectId)
	return nil
}

// GetCollectId CollectId Getter
func (r XiamicontentmusiccollectgetAPIRequest) GetCollectId() int64 {
	return r._collectId
}

// SetPageSize is PageSize Setter
// 每页数量
func (r *XiamicontentmusiccollectgetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r XiamicontentmusiccollectgetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPage is Page Setter
// 页码
func (r *XiamicontentmusiccollectgetAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r XiamicontentmusiccollectgetAPIRequest) GetPage() int64 {
	return r._page
}
