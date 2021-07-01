package fundplatform

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaFundplatformCardordersStatusSendedAPIRequest
制卡商通知实体卡发货完成 API请求
alibaba.fundplatform.cardorders.status.sended

当制卡商将实体卡发货完成后，需要调用该接口，通知我们已发货。 */
type AlibabaFundplatformCardordersStatusSendedAPIRequest struct {
	model.Params
	// 子制卡单ID
	_cardOrderId int64
	// 物流单号
	_logisticsOrderId string
	// 物流商名称
	_logisticsCompany string
}

// New
