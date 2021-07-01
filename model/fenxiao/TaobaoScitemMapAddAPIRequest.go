package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoScitemMapAddAPIRequest
创建IC商品与后端商品的映射关系 API请求
taobao.scitem.map.add

创建IC商品或分销商品与后端商品的映射关系 */
type TaobaoScitemMapAddAPIRequest struct {
	model.Params
	// 前台ic商品id
	_itemId int64
	// 前台ic商品skuid
	_skuId int64
	// sc_item_id和outer_code 其中一个不能为空
	_scItemId int64
	// sc_item_id和outer_code 其中一个不能为空
	_outerCode string
	// 默认值为false<br/>true:进行高级校验,前端商品或SKU的商家编码必须与后端商品的商家编码一致，否则会拒绝关联<br/>false:不进行高级校验
	_needCheck bool
}

// New
