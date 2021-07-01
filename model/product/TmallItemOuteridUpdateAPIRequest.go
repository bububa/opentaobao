package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallItemOuteridUpdateAPIRequest
天猫商品/SKU商家编码更新接口 API请求
tmall.item.outerid.update

天猫商品/SKU商家编码更新接口；支持商品、SKU的商家编码同时更新；支持同一商品下的SKU批量更新。（感谢sample小雨提供接口命名） */
type TmallItemOuteridUpdateAPIRequest struct {
	model.Params
	// 商品ID
	_itemId int64
	// 商品维度商家编码，如果不修改可以不传；清空请设置空串
	_outerId string
	// 商品SKU更新OuterId时候用的数据
	_skuOuters []UpdateSkuOuterId
}

// New
