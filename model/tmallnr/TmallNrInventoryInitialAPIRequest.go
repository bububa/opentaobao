package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallNrInventoryInitialAPIRequest
门店库存初始化前后端商品绑定 API请求
tmall.nr.inventory.initial

用于门店业务的商品的初始化，前端商品和后端商品绑定，走后端库存模式 */
type TmallNrInventoryInitialAPIRequest struct {
	model.Params
	// 请求入参
	_param0 *NrStoreInvItemInitialReqDto
}

// New
