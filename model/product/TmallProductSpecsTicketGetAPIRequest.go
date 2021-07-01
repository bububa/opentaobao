package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallProductSpecsTicketGetAPIRequest
产品规格审核信息获取接口 API请求
tmall.product.specs.ticket.get

批量根据specId查询产品规格审核信息包括产品规格状态，申请人，拒绝原因等 */
type TmallProductSpecsTicketGetAPIRequest struct {
	model.Params
	// 产品规格ID，多个用逗号分隔
	_specIds string
}

// New
