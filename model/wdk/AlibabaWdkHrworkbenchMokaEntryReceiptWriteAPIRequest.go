package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIRequest
摩卡确认入职后往入职单据表写数据接口 API请求
alibaba.wdk.hrworkbench.moka.entry.receipt.write

摩卡确认入职后往入职单据表写数据接口 */
type AlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIRequest struct {
	model.Params
	// 确认入职后入职单据请求
	_paramSyncEntryReceiptRequest *SyncEntryReceiptRequest
}

// New
