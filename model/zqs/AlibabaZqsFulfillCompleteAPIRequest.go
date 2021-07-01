package zqs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaZqsFulfillCompleteAPIRequest
周期购履约完成接口 API请求
alibaba.zqs.fulfill.complete

周期购履约完成接口 */
type AlibabaZqsFulfillCompleteAPIRequest struct {
	model.Params
	// 第几期
	_sequenceNo int64
	// 交易单号
	_mainBizOrderId int64
	// 交易子单号
	_subBizOrderId int64
}

// New
