package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceProdApplyGet 查询发票申请
// alibaba.einvoice.prod.apply.get
//
// 查询申请的详细信息，包含申请所关联的发票摘要信息+板式文件+预览图；
// 场景使用：1、业务前台收到申请状态变更消息后，调用此接口查询申请详情；2、主动补偿查询：当指定了自动开票，且发票申请长时间未收到状态变更通知时，可能存在丢消息的情况，此时可主动查询该申请，然后更新本地工单状态。
func AlibabaEinvoiceProdApplyGet(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceProdApplyGetAPIRequest, session string) (*einvoice.AlibabaEinvoiceProdApplyGetAPIResponse, error) {
	var resp einvoice.AlibabaEinvoiceProdApplyGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
