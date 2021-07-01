package westcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWestcrmUpdateAlipayCarnoAPIRequest
更新支付宝业务卡号 API请求
alibaba.westcrm.update.alipay.carno

更新支付宝业务卡号 */
type AlibabaWestcrmUpdateAlipayCarnoAPIRequest struct {
	model.Params
	// 商场id
	_mallId int64
	// 用户id
	_id int64
	// 2088102011918821
	_alipayCardNo string
	// appkey
	_westcrmAppKey string
}

// New
