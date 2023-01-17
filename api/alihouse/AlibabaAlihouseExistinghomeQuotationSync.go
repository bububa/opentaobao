package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeQuotationSync 二手房行情数据同步
// alibaba.alihouse.existinghome.quotation.sync
//
// 二手房行情数据同步
func AlibabaAlihouseExistinghomeQuotationSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeQuotationSyncAPIRequest, session string) (*alihouse.AlibabaAlihouseExistinghomeQuotationSyncAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseExistinghomeQuotationSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
