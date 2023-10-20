package degoperation

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/degoperation"
)

// TaobaoDegoperationShowUserRecords 用户中奖记录
// taobao.degoperation.show.user.records
//
// 用户中奖记录
func TaobaoDegoperationShowUserRecords(clt *core.SDKClient, req *degoperation.TaobaoDegoperationShowUserRecordsAPIRequest, session string) (*degoperation.TaobaoDegoperationShowUserRecordsAPIResponse, error) {
	var resp degoperation.TaobaoDegoperationShowUserRecordsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
