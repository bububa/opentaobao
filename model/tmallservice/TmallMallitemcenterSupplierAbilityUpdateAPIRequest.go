package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallMallitemcenterSupplierAbilityUpdateAPIRequest
门店服务能力授权接口 API请求
tmall.mallitemcenter.supplier.ability.update

门店服务能力授权 */
type TmallMallitemcenterSupplierAbilityUpdateAPIRequest struct {
	model.Params
	// 入参
	_param0 *EnableServiceStoreRequestDto
}

// New
