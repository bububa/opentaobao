package fundplatform

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaFundplatformCardorderMakeSuccessAPIRequest
通知制卡成功 API请求
alibaba.fundplatform.cardorder.make.success

当外部业务方调用资金平台异步制卡接口后，资金平台制卡成功后通知异步通知业务方 */
type AlibabaFundplatformCardorderMakeSuccessAPIRequest struct {
	model.Params
	// 入参对象
	_request *AlibabaFundplatformCardorderMakeSuccessStruct
}

// New
