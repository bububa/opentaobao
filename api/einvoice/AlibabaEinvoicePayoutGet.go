package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoicepayoutget 获取赔付计时列表数据
// alibaba.einvoice.payout.get
//
// 获取赔付计时列表数据
func Alibabaeinvoicepayoutget(clt *core.SDKClient, req *einvoice.AlibabaeinvoicepayoutgetAPIRequest, session string) (*einvoice.AlibabaeinvoicepayoutgetAPIResponse, error) {
	var resp einvoice.AlibabaeinvoicepayoutgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
