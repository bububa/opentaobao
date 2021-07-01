package fundplatform

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaFundplatformCardorderFetchCardAsyncAPIRequest
异步批量生成储值卡 API请求
alibaba.fundplatform.cardorder.fetch.card.async

外部业务方异步批量生成储值卡的接口。同步只返回接受成功，异步会通知制卡成功 */
type AlibabaFundplatformCardorderFetchCardAsyncAPIRequest struct {
	model.Params
	// 入参复杂对象
	_paramCardFetchAsyncRequest *CardFetchAsyncRequest
}

// New
