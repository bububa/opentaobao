package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbItemBatchQueryAPIRequest
批次库存查询接口 API请求
taobao.wlb.item.batch.query

根据用户id，item id list和store code来查询商品库存信息和批次信息 */
type TaobaoWlbItemBatchQueryAPIRequest struct {
	model.Params
	// 仓库编号
	_storeCode string
	// 分页查询参数，指定查询页数，默认为1
	_pageNo int64
	// 分页查询参数，每页查询数量，默认20，最大值50,大于50时按照50条查询
	_pageSize int64
	// 需要查询的商品ID列表，以字符串表示，ID间以;隔开
	_itemIds string
}

// New
