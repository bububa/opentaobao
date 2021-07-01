package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceFlowTaxGetAPIRequest
查询税控开通工单详情 API请求
alibaba.einvoice.flow.tax.get

查询税控开通工单详情，接口返回工单状态、开票商户信息以及税控设备信息。
场景使用：1、业务前台收到入驻成功消息后，调用此接口查询最终的商户信息和设备信息；2、主动补偿查询：当工单长时间未收到事件通知，可能存在丢消息的情况，此时可主动查询该工单，更新本地工单状态。 */
type AlibabaEinvoiceFlowTaxGetAPIRequest struct {
	model.Params
	// 入驻开通工单ID
	_flowId string
}

// New
