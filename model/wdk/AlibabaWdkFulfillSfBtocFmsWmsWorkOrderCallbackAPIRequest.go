package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackAPIRequest
顺丰仓作业回传 API请求
alibaba.wdk.fulfill.sf.btoc.fms.wms.work.order.callback

顺丰仓作业单回传接口 */
type AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackAPIRequest struct {
	model.Params
	// 作业单回传对象
	_callbackOrder *SfB2CFmsCallbackOrder
}

// New
