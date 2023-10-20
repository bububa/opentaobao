package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaomediavideolistAPIRequest 获取商家视频列表 API请求
// taobao.media.video.list
//
// 用于获取授权商家的视频列表
type TaobaomediavideolistAPIRequest struct {
	model.Params
	// 搜索条件
	_searchCondition *VideoSearchCondition2
}

// NewTaobaomediavideolistRequest 初始化TaobaomediavideolistAPIRequest对象
func NewTaobaomediavideolistRequest() *TaobaomediavideolistAPIRequest {
	return &TaobaomediavideolistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaomediavideolistAPIRequest) GetApiMethodName() string {
	return "taobao.media.video.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaomediavideolistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaomediavideolistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSearchCondition is SearchCondition Setter
// 搜索条件
func (r *TaobaomediavideolistAPIRequest) SetSearchCondition(_searchCondition *VideoSearchCondition2) error {
	r._searchCondition = _searchCondition
	r.Set("search_condition", _searchCondition)
	return nil
}

// GetSearchCondition SearchCondition Getter
func (r TaobaomediavideolistAPIRequest) GetSearchCondition() *VideoSearchCondition2 {
	return r._searchCondition
}
