package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinFlowWalletGradeAPIRequest
获取流量档位 API请求
alibaba.aliqin.flow.wallet.grade

获取直充流量档位 */
type AlibabaAliqinFlowWalletGradeAPIRequest struct {
	model.Params
	// 手机号码
	_phoneNum string
}

// New
