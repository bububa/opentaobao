package degoperation

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/degoperation"
)

// TaobaoDegoperationShowUserRecords 用户中奖记录
// taobao.degoperation.show.user.records
//
// 用户中奖记录
func TaobaoDegoperationShowUserRecords(ctx context.Context, clt *core.SDKClient, req *degoperation.TaobaoDegoperationShowUserRecordsAPIRequest, resp *degoperation.TaobaoDegoperationShowUserRecordsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
