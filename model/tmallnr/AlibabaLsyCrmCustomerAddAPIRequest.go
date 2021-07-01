package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLsyCrmCustomerAddAPIRequest
私域导购添加活动留资入口 API请求
alibaba.lsy.crm.customer.add

私域导购添加活动留资入口 */
type AlibabaLsyCrmCustomerAddAPIRequest struct {
	model.Params
	// 入参对象
	_reqDto *NrtCrmCreateCustomerReq
}

// New
