package xiamicontent

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// XiamiContentSongsCollectGetAPIRequest 获取歌单详情接口 API请求
// xiami.content.songs.collect.get
//
// 根据歌单id，获取歌单详情
type XiamiContentSongsCollectGetAPIRequest struct {
	model.Params
	// 歌单id
	_collectId int64
	// 分页信息
	_page *PagingVO
}

// NewXiamiContentSongsCollectGetRequest 初始化XiamiContentSongsCollectGetAPIRequest对象
func NewXiamiContentSongsCollectGetRequest() *XiamiContentSongsCollectGetAPIRequest {
	return &XiamiContentSongsCollectGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r XiamiContentSongsCollectGetAPIRequest) GetApiMethodName() string {
	return "xiami.content.songs.collect.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r XiamiContentSongsCollectGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCollectId is CollectId Setter
// 歌单id
func (r *XiamiContentSongsCollectGetAPIRequest) SetCollectId(_collectId int64) error {
	r._collectId = _collectId
	r.Set("collect_id", _collectId)
	return nil
}

// GetCollectId CollectId Getter
func (r XiamiContentSongsCollectGetAPIRequest) GetCollectId() int64 {
	return r._collectId
}

// SetPage is Page Setter
// 分页信息
func (r *XiamiContentSongsCollectGetAPIRequest) SetPage(_page *PagingVO) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r XiamiContentSongsCollectGetAPIRequest) GetPage() *PagingVO {
	return r._page
}
