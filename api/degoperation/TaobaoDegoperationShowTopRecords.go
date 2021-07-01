package degoperation

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/degoperation"
)

/* 
活动中奖记录 
taobao.degoperation.show.top.records

活动中奖记录
*/
func TaobaoDegoperationShowTopRecords(clt *core.SDKClient, req *degoperation.TaobaoDegoperationShowTopRecordsAPIRequest, session string) (*degoperation.TaobaoDegoperationShowTopRecordsAPIResponse, error) {
    var resp degoperation.TaobaoDegoperationShowTopRecordsAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
