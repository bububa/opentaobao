package degoperation

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/degoperation"
)

// TaobaoDegoperationDoLuckydraw 激励抽奖
// taobao.degoperation.do.luckydraw
//
// 激励平台抽奖接口。用户可以通过接口完成抽奖功能
func TaobaoDegoperationDoLuckydraw(ctx context.Context, clt *core.SDKClient, req *degoperation.TaobaoDegoperationDoLuckydrawAPIRequest, resp *degoperation.TaobaoDegoperationDoLuckydrawAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
