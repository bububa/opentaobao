package westcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWestcrmAccountIdGetAPIRequest
根据支付宝id查询IB的id API请求
alibaba.westcrm.account.id.get

根据支付宝id查询IB的id */
type AlibabaWestcrmAccountIdGetAPIRequest struct {
	model.Params
	// 支付宝id
	_alipayId string
}

// New
