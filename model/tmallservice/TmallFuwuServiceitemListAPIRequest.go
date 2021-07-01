package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallFuwuServiceitemListAPIRequest
获取服务商品扩展信息 API请求
tmall.fuwu.serviceitem.list

获取服务商品扩展信息 */
type TmallFuwuServiceitemListAPIRequest struct {
	model.Params
	// 商品所属卖家账号id
	_sellerId int64
	// 商品id列表，有数量限制
	_itemids []int64
}

// New
