package fundplatform

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fundplatform"
)

// AlibabaFundplatformCardordersInfoQueryByCardno 通过卡号查询卡信息
// alibaba.fundplatform.cardorders.info.query.by.cardno
//
// 该接口由汇金实现，外部调用。通过制卡单号分页查询卡信息
func AlibabaFundplatformCardordersInfoQueryByCardno(ctx context.Context, clt *core.SDKClient, req *fundplatform.AlibabaFundplatformCardordersInfoQueryByCardnoAPIRequest, resp *fundplatform.AlibabaFundplatformCardordersInfoQueryByCardnoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
