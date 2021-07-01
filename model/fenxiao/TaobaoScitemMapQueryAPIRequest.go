package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoScitemMapQueryAPIRequest
查找IC商品或分销商品与后端商品的关联信息 API请求
taobao.scitem.map.query

查找IC商品或分销商品与后端商品的关联信息。skuId如果不传就查找该itemId下所有的sku */
type TaobaoScitemMapQueryAPIRequest struct {
	model.Params
	// map_type为1：前台ic商品id<br/>map_type为2：分销productid
	_itemId int64
	// map_type为1：前台ic商品skuid <br/>map_type为2：分销商品的skuid
	_skuId int64
}

// New
