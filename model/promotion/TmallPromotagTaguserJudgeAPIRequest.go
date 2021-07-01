package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallPromotagTaguserJudgeAPIRequest
用户标签判断接口 API请求
tmall.promotag.taguser.judge

查询用户是否有标签 */
type TmallPromotagTaguserJudgeAPIRequest struct {
	model.Params
	// 标签ID
	_tagId int64
	// 买家昵称
	_nick string
}

// New
