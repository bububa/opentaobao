package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceSerialnoBatchGenerateAPIRequest
开票流水号批量生成接口 API请求
alibaba.einvoice.serialno.batch.generate

批量获取开票流水号接口。此接口1次返回1000条开票流水号，每个应用每天限流1000次调用。
优先使用alibaba.einvoice.serial.generate。 */
type AlibabaEinvoiceSerialnoBatchGenerateAPIRequest struct {
	model.Params
}

// New
