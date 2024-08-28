package fundplatform

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fundplatform"
)

// AlibabaFundplatformCardordersInfoQuery 根据制卡单分页查询卡信息
// alibaba.fundplatform.cardorders.info.query
//
// 该接口由汇金实现，外部调用。通过制卡单号分页查询卡信息
func AlibabaFundplatformCardordersInfoQuery(ctx context.Context, clt *core.SDKClient, req *fundplatform.AlibabaFundplatformCardordersInfoQueryAPIRequest, resp *fundplatform.AlibabaFundplatformCardordersInfoQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
