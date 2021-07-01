package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallNrInventoryUpdateAPIRequest
门店业务同步库存 API请求
tmall.nr.inventory.update

用于商家每日同步更新门店库存 */
type TmallNrInventoryUpdateAPIRequest struct {
	model.Params
	// 入参
	_param0 *NrInventoryUpdateReqDto
}

// New
