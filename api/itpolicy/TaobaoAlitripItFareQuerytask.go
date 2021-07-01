package itpolicy

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/itpolicy"
)

/* 
【国际机票自有政策】批量操作结果查询 
taobao.alitrip.it.fare.querytask

批量操作同步返回任务id，后台生成异步任务，通过此接口查询批量操作的执行结果
*/
func TaobaoAlitripItFareQuerytask(clt *core.SDKClient, req *itpolicy.TaobaoAlitripItFareQuerytaskAPIRequest, session string) (*itpolicy.TaobaoAlitripItFareQuerytaskAPIResponse, error) {
    var resp itpolicy.TaobaoAlitripItFareQuerytaskAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
