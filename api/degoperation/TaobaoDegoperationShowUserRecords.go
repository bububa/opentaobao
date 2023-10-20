package degoperation

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/degoperation"
)

// TaobaoDegoperationShowUserRecords 用户中奖记录
// taobao.degoperation.show.user.records
//
// 用户中奖记录
func TaobaoDegoperationShowUserRecords(clt *core.SDKClient, req *degoperation.TaobaoDegoperationShowUserRecordsAPIRequest, resp *degoperation.TaobaoDegoperationShowUserRecordsAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
