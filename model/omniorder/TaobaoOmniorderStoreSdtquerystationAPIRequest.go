package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniorderStoreSdtquerystationAPIRequest
速店通查询站点信息 API请求
taobao.omniorder.store.sdtquerystation

速店通查询站点信息 */
type TaobaoOmniorderStoreSdtquerystationAPIRequest struct {
	model.Params
	// 取号时返回的packageId
	_paramLong2 int64
}

// New
