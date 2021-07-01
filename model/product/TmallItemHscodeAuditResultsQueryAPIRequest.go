package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallItemHscodeAuditResultsQueryAPIRequest
商品hscode信息审核状态查询接口 API请求
tmall.item.hscode.audit.results.query

通过此接口查询天猫跨境商品的hscode信息审核状态，卖家可以参考返回结果判断是否需要调整商品hscode相关信息。 */
type TmallItemHscodeAuditResultsQueryAPIRequest struct {
	model.Params
	// 商品ID
	_itemId int64
}

// New
