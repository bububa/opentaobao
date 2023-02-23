package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMediaVideoListAPIRequest 获取商家视频列表 API请求
// taobao.media.video.list
//
// 用于获取授权商家的视频列表
type TaobaoMediaVideoListAPIRequest struct {
	model.Params
	// 搜索条件
	_searchCondition *VideoSearchCondition2
}

// NewTaobaoMediaVideoListRequest 初始化TaobaoMediaVideoListAPIRequest对象
func NewTaobaoMediaVideoListRequest() *TaobaoMediaVideoListAPIRequest {
	return &TaobaoMediaVideoListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMediaVideoListAPIRequest) GetApiMethodName() string {
	return "taobao.media.video.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMediaVideoListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMediaVideoListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSearchCondition is SearchCondition Setter
// 搜索条件
func (r *TaobaoMediaVideoListAPIRequest) SetSearchCondition(_searchCondition *VideoSearchCondition2) error {
	r._searchCondition = _searchCondition
	r.Set("search_condition", _searchCondition)
	return nil
}

// GetSearchCondition SearchCondition Getter
func (r TaobaoMediaVideoListAPIRequest) GetSearchCondition() *VideoSearchCondition2 {
	return r._searchCondition
}
