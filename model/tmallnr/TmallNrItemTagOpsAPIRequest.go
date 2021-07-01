package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallNrItemTagOpsAPIRequest
区域零售商品打标去标 API请求
tmall.nr.item.tag.ops

参加区域零售的商品，需要批量打标或去标，方便后续设置商品库存 */
type TmallNrItemTagOpsAPIRequest struct {
	model.Params
	// 请求入参
	_tagReqDTO *TagReqDto
}

// New
