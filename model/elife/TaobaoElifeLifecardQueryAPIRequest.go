package elife

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoElifeLifecardQueryAPIRequest
查询交易结果 API请求
taobao.elife.lifecard.query

卖家在交易状态不明的情况下, 查询交易结果. */
type TaobaoElifeLifecardQueryAPIRequest struct {
	model.Params
	// 入参
	_queryRequest *ConsumeRequest
}

// New
