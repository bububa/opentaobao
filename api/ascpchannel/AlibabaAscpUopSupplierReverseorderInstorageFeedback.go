package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// Alibabaascpuopsupplierreverseorderinstoragefeedback 逆向销退入库单入库结果回告
// alibaba.ascp.uop.supplier.reverseorder.instorage.feedback
//
// ASCP按照逆向履约单纬度，通过该接口接收商家在退货完成时，自动创建销退单做入库回传。
func Alibabaascpuopsupplierreverseorderinstoragefeedback(clt *core.SDKClient, req *ascpchannel.AlibabaascpuopsupplierreverseorderinstoragefeedbackAPIRequest, session string) (*ascpchannel.AlibabaascpuopsupplierreverseorderinstoragefeedbackAPIResponse, error) {
	var resp ascpchannel.AlibabaascpuopsupplierreverseorderinstoragefeedbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
