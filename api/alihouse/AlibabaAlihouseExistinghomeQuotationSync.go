package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeQuotationSync 二手房行情数据同步
// alibaba.alihouse.existinghome.quotation.sync
//
// 二手房行情数据同步
func AlibabaAlihouseExistinghomeQuotationSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeQuotationSyncAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeQuotationSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
