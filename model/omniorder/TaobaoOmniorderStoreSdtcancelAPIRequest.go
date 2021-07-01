package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniorderStoreSdtcancelAPIRequest
通知速店通取消取号 API请求
taobao.omniorder.store.sdtcancel

通知速店通取消取号 */
type TaobaoOmniorderStoreSdtcancelAPIRequest struct {
	model.Params
	// 取号返回的packageId
	_packageId int64
}

// New
