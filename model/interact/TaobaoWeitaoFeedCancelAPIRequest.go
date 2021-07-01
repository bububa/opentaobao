package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWeitaoFeedCancelAPIRequest
取消广播在timeline、广场中展示 API请求
taobao.weitao.feed.cancel

取消广播在timeline和广场中的展示。 */
type TaobaoWeitaoFeedCancelAPIRequest struct {
	model.Params
	// 广播id
	_feedId int64
	// 三方活动ID
	_bizId string
	// 是否彻底删除（店铺动态不可见，等同卖家广播后台删除），默认false
	_delete bool
}

// New
