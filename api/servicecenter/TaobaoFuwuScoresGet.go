package servicecenter

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/servicecenter"
)

/* 
服务平台评价查询接口 
taobao.fuwu.scores.get

根据日期、查询appkey对应服务评价，每次调用只能查询某一天服务评价信息，可设置分页查询，页大小最大为100，非实时接口，延迟时间为30分钟
*/
func TaobaoFuwuScoresGet(clt *core.SDKClient, req *servicecenter.TaobaoFuwuScoresGetRequest, session string) (*servicecenter.TaobaoFuwuScoresGetResponse, error) {
    var resp servicecenter.TaobaoFuwuScoresGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
