package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMediaVideoListAPIRequest
获取商家视频列表 API请求
taobao.media.video.list

用于获取授权商家的视频列表 */
type TaobaoMediaVideoListAPIRequest struct {
	model.Params
	// 搜索条件
	_searchCondition *VideoSearchCondition2
}

// New
