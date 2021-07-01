package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallPromotagTaguserSaveAPIRequest
给用户打上优惠标签 API请求
tmall.promotag.taguser.save

给用户载体打标 */
type TmallPromotagTaguserSaveAPIRequest struct {
	model.Params
	// 标签ID
	_tagId int64
	// 买家昵称
	_nick string
}

// New
