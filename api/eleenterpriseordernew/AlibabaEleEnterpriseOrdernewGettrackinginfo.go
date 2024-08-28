package eleenterpriseordernew

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterpriseordernew"
)

// AlibabaEleEnterpriseOrdernewGettrackinginfo 订单配送信息跟踪
// alibaba.ele.enterprise.ordernew.gettrackinginfo
//
// 订单配送信息跟踪
func AlibabaEleEnterpriseOrdernewGettrackinginfo(ctx context.Context, clt *core.SDKClient, req *eleenterpriseordernew.AlibabaEleEnterpriseOrdernewGettrackinginfoAPIRequest, resp *eleenterpriseordernew.AlibabaEleEnterpriseOrdernewGettrackinginfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
