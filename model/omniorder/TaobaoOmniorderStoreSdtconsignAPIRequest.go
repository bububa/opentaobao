package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniorderStoreSdtconsignAPIRequest
通知菜鸟裹裹发货 API请求
taobao.omniorder.store.sdtconsign

ISV取完单号后通知菜鸟裹裹发货 */
type TaobaoOmniorderStoreSdtconsignAPIRequest struct {
	model.Params
	// 取号接口返回的包裹id
	_packageId string
	// 发货标签号
	_tagCode string
}

// New
