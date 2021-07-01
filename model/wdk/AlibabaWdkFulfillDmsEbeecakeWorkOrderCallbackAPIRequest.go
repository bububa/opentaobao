package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackAPIRequest
北京小蜜蜂配作业回传 API请求
alibaba.wdk.fulfill.dms.ebeecake.work.order.callback

北京小蜜蜂配作业回传。 */
type AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackAPIRequest struct {
	model.Params
	// 作业单回传对象
	_callbackOrder *EbeecakeO2OCallbackOrder
}

// New
