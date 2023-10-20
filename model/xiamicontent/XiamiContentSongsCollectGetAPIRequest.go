package xiamicontent

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// XiamicontentsongscollectgetAPIRequest 获取歌单详情接口 API请求
// xiami.content.songs.collect.get
//
// 根据歌单id，获取歌单详情
type XiamicontentsongscollectgetAPIRequest struct {
	model.Params
	// 歌单id
	_collectId int64
	// 分页信息
	_page *PagingVo
}

// NewXiamicontentsongscollectgetRequest 初始化XiamicontentsongscollectgetAPIRequest对象
func NewXiamicontentsongscollectgetRequest() *XiamicontentsongscollectgetAPIRequest {
	return &XiamicontentsongscollectgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r XiamicontentsongscollectgetAPIRequest) GetApiMethodName() string {
	return "xiami.content.songs.collect.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r XiamicontentsongscollectgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r XiamicontentsongscollectgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCollectId is CollectId Setter
// 歌单id
func (r *XiamicontentsongscollectgetAPIRequest) SetCollectId(_collectId int64) error {
	r._collectId = _collectId
	r.Set("collect_id", _collectId)
	return nil
}

// GetCollectId CollectId Getter
func (r XiamicontentsongscollectgetAPIRequest) GetCollectId() int64 {
	return r._collectId
}

// SetPage is Page Setter
// 分页信息
func (r *XiamicontentsongscollectgetAPIRequest) SetPage(_page *PagingVo) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r XiamicontentsongscollectgetAPIRequest) GetPage() *PagingVo {
	return r._page
}
