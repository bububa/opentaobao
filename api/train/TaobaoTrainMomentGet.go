package train

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/train"
)

/* 
火车票时刻表 
taobao.train.moment.get

查询火车票车次时刻表
*/
func TaobaoTrainMomentGet(clt *core.SDKClient, req *train.TaobaoTrainMomentGetRequest, session string) (*train.TaobaoTrainMomentGetAPIResponse, error) {
    var resp train.TaobaoTrainMomentGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
