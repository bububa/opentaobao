package normalvisa

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/normalvisa"
)

/* 
更新签证办理进度 
taobao.alitrip.travel.normalvisa.updatepersonstauts

更新签证办理进度
*/
func TaobaoAlitripTravelNormalvisaUpdatepersonstauts(clt *core.SDKClient, req *normalvisa.TaobaoAlitripTravelNormalvisaUpdatepersonstautsRequest, session string) (*normalvisa.TaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIResponse, error) {
    var resp normalvisa.TaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
