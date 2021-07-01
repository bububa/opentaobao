package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniorderStoreSdtstatusAPIRequest
菜鸟裹裹运单状态查询 API请求
taobao.omniorder.store.sdtstatus

提供给商家查询运力单的状态。 */
type TaobaoOmniorderStoreSdtstatusAPIRequest struct {
	model.Params
	// 菜鸟裹裹的包裹ID
	_packageId int64
}

// New
