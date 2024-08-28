package eleenterpriseordernew

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterpriseordernew"
)

// AlibabaEleEnterpriseOrdernewGet 查询订单详情
// alibaba.ele.enterprise.ordernew.get
//
// 查询订单详情
func AlibabaEleEnterpriseOrdernewGet(ctx context.Context, clt *core.SDKClient, req *eleenterpriseordernew.AlibabaEleEnterpriseOrdernewGetAPIRequest, resp *eleenterpriseordernew.AlibabaEleEnterpriseOrdernewGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
