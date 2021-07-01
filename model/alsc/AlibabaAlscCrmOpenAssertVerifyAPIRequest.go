package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmOpenAssertVerifyAPIRequest
资产核销接口 API请求
alibaba.alsc.crm.open.assert.verify

核销储值，积分，券资产 */
type AlibabaAlscCrmOpenAssertVerifyAPIRequest struct {
	model.Params
	// 入参
	_paramPropertyVerifyOpenReq *PropertyVerifyOpenReq
}

// New
