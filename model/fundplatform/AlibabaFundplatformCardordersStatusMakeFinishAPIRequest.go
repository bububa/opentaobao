package fundplatform

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaFundplatformCardordersStatusMakeFinishAPIRequest
制卡商通知制卡完成 API请求
alibaba.fundplatform.cardorders.status.make.finish

当制卡完成后，制卡商需要调用该接口，通知我们制卡已完成。 */
type AlibabaFundplatformCardordersStatusMakeFinishAPIRequest struct {
	model.Params
	// 子制卡单ID
	_cardOrderId int64
}

// New
