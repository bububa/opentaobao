package degoperation

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/degoperation"
)

// TaobaoDegoperationGetInfoUuid 根据uuid用户抽奖次数限制
// taobao.degoperation.get.info.uuid
//
// 根据uuid用户抽奖次数限制
func TaobaoDegoperationGetInfoUuid(ctx context.Context, clt *core.SDKClient, req *degoperation.TaobaoDegoperationGetInfoUuidAPIRequest, resp *degoperation.TaobaoDegoperationGetInfoUuidAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
