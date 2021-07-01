package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleRentOrderCheckstatusUploadAPIRequest
上传验收结果 API请求
alibaba.idle.rent.order.checkstatus.upload

上传验收结果 */
type AlibabaIdleRentOrderCheckstatusUploadAPIRequest struct {
	model.Params
	// 订单id
	_orderId int64
	// 校验结果
	_checkResult *CheckResultDto
}

// New
