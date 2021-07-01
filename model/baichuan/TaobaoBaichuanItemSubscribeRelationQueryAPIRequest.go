package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBaichuanItemSubscribeRelationQueryAPIRequest
查询单个订阅关系 API请求
taobao.baichuan.item.subscribe.relation.query

查询单个订阅关系 */
type TaobaoBaichuanItemSubscribeRelationQueryAPIRequest struct {
	model.Params
	// 商品Id
	_itemId int64
}

// New
