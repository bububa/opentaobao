package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallCarLeaseItemcarinfoAPIRequest
整车租赁商品四级车型信息 API请求
tmall.car.lease.itemcarinfo

整车租赁项目发布宝贝需要4级车型库，4级车型库信息需要回传 */
type TmallCarLeaseItemcarinfoAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
}

// New
