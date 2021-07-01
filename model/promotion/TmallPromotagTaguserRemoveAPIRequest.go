package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallPromotagTaguserRemoveAPIRequest
给用户移除优惠标签 API请求
tmall.promotag.taguser.remove

给用户载体去标 */
type TmallPromotagTaguserRemoveAPIRequest struct {
	model.Params
	// 标签ID
	_tagId int64
	// 买家昵称
	_nick string
}

// New
