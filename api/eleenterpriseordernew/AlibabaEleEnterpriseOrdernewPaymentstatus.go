package eleenterpriseordernew

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterpriseordernew"
)

// AlibabaEleEnterpriseOrdernewPaymentstatus 设置订单支付
// alibaba.ele.enterprise.ordernew.paymentstatus
//
// 设置订单支付成功
func AlibabaEleEnterpriseOrdernewPaymentstatus(ctx context.Context, clt *core.SDKClient, req *eleenterpriseordernew.AlibabaEleEnterpriseOrdernewPaymentstatusAPIRequest, resp *eleenterpriseordernew.AlibabaEleEnterpriseOrdernewPaymentstatusAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
