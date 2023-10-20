package degoperation

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/degoperation"
)

// TaobaoDegoperationShowTopRecords 活动中奖记录
// taobao.degoperation.show.top.records
//
// 活动中奖记录
func TaobaoDegoperationShowTopRecords(clt *core.SDKClient, req *degoperation.TaobaoDegoperationShowTopRecordsAPIRequest, resp *degoperation.TaobaoDegoperationShowTopRecordsAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
