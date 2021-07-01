package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmOpenAssertRefundAPIRequest
资产核销回退接口 API请求
alibaba.alsc.crm.open.assert.refund

回退已经核销的储值积分券资产 */
type AlibabaAlscCrmOpenAssertRefundAPIRequest struct {
	model.Params
	// 入参
	_paramPropertyRefundOpenReq *PropertyRefundOpenReq
}

// New
