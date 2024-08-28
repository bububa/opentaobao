package lstpos

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstpos"
)

// AlibabaLstPosOpenAccountCheckissettled 校验当前用户是否入驻了零售通门店接口
// alibaba.lst.pos.open.account.checkissettled
//
// 校验当前用户是否入驻了零售通门店接口
func AlibabaLstPosOpenAccountCheckissettled(ctx context.Context, clt *core.SDKClient, req *lstpos.AlibabaLstPosOpenAccountCheckissettledAPIRequest, resp *lstpos.AlibabaLstPosOpenAccountCheckissettledAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
