package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// Alibabaascpuopsupplierreverseorderinstorageresult 逆向销退入库单到仓结果回告
// alibaba.ascp.uop.supplier.reverseorder.instorage.result
//
// ERP回告销退入库单到仓信息回告
func Alibabaascpuopsupplierreverseorderinstorageresult(clt *core.SDKClient, req *ascpchannel.AlibabaascpuopsupplierreverseorderinstorageresultAPIRequest, session string) (*ascpchannel.AlibabaascpuopsupplierreverseorderinstorageresultAPIResponse, error) {
	var resp ascpchannel.AlibabaascpuopsupplierreverseorderinstorageresultAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
