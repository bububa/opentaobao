package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkHrworkbenchMokaEntryReceiptWrite 摩卡确认入职后往入职单据表写数据接口
// alibaba.wdk.hrworkbench.moka.entry.receipt.write
//
// 摩卡确认入职后往入职单据表写数据接口
func AlibabaWdkHrworkbenchMokaEntryReceiptWrite(clt *core.SDKClient, req *wdk.AlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIRequest, resp *wdk.AlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
