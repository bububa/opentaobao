package fundplatform

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaFundplatformAccountChargeAPIRequest
资金平台余额账户充值 API请求
alibaba.fundplatform.account.charge

资金平台余额账户充值【创建账户&返回付款URL】 */
type AlibabaFundplatformAccountChargeAPIRequest struct {
	model.Params
	// 用户ID
	_paramLong int64
	// 入参对象
	_paramChargeRequest *ChargeRequest
}

// New
