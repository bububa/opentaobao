package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaServicecenterFulfiltaskBuyeraddressChangeAPIRequest
修改消费者服务地址 API请求
alibaba.servicecenter.fulfiltask.buyeraddress.change

当消费者反馈自己的服务地址错误时，可以电话联系服务商修改为正确地址，服务商只能修改派给自己的单子 */
type AlibabaServicecenterFulfiltaskBuyeraddressChangeAPIRequest struct {
	model.Params
	// 核销单id
	_fulfilTaskId int64
	// 详细地址
	_addressDetail string
	// 地址编码
	_location int64
}

// New
