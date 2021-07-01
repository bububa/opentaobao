package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelDistributionOrderDetailSearchAPIRequest
分销渠道订单详情查询 API请求
taobao.xhotel.distribution.order.detail.search

该接口用于提供酒店分销渠道的订单详情查询 */
type TaobaoXhotelDistributionOrderDetailSearchAPIRequest struct {
	model.Params
	// 外部分销渠道的订单号（与tid必填其一）
	_distributionOid string
	// 传入用户对应的openId
	_openId string
	// 飞猪的订单号（与distribution_oid必填其一）
	_tid int64
}

// New
