package westcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWestcrmCustomerInfoGetAPIRequest
会员信息查询接口 API请求
alibaba.westcrm.customer.info.get

会员信息查询接口 */
type AlibabaWestcrmCustomerInfoGetAPIRequest struct {
	model.Params
	// 园区id
	_campusId int64
	// 会员id
	_ibUserId int64
	// 支付宝id
	_alipayId string
}

// New
