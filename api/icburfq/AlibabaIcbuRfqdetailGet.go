package icburfq

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icburfq"
)

// AlibabaIcbuRfqdetailGet 获取RFQ详情
// alibaba.icbu.rfqdetail.get
//
// 查看RFQ的详情信息
func AlibabaIcbuRfqdetailGet(ctx context.Context, clt *core.SDKClient, req *icburfq.AlibabaIcbuRfqdetailGetAPIRequest, resp *icburfq.AlibabaIcbuRfqdetailGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
