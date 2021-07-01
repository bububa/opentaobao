package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoInteractiveListGetbyuserAPIRequest
用户获取视频互动列表 API请求
taobao.interactive.list.getbyuser

根据用户来获取用户编辑的互动列表 */
type TaobaoInteractiveListGetbyuserAPIRequest struct {
	model.Params
	// 当前页
	_currentPage int64
	// 每页多少
	_pageSize int64
}

// New
