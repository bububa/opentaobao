package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoiceincomescanreturn 进项扫描状态回传
// alibaba.einvoice.income.scan.return
//
// 回传进项扫描每个阶段的状态，比如ocr开始，ocr结束，查验开始，查验结束等
func Alibabaeinvoiceincomescanreturn(clt *core.SDKClient, req *einvoice.AlibabaeinvoiceincomescanreturnAPIRequest, session string) (*einvoice.AlibabaeinvoiceincomescanreturnAPIResponse, error) {
	var resp einvoice.AlibabaeinvoiceincomescanreturnAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
