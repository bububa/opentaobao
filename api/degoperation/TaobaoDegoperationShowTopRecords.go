package degoperation

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/degoperation"
)

// TaobaoDegoperationShowTopRecords 活动中奖记录
// taobao.degoperation.show.top.records
//
// 活动中奖记录
func TaobaoDegoperationShowTopRecords(ctx context.Context, clt *core.SDKClient, req *degoperation.TaobaoDegoperationShowTopRecordsAPIRequest, resp *degoperation.TaobaoDegoperationShowTopRecordsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
