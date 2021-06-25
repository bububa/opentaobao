package wangwang

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
客服评价详情接口 APIRequest
taobao.qianniu.kefueval.get

获取买家对客服的服务评价
*/
type TaobaoQianniuKefuevalGetRequest struct {
    model.Params

    // 客服的nick，多个用逗号分隔，不要超过10个，带cntaobao的长nick
    queryIds   string 

    // 开始时间，格式yyyyMMddHHmmss
    btime   string 

    // 结束时间，格式yyyyMMddHHmmss
    etime   string 

}

func NewTaobaoQianniuKefuevalGetRequest() *TaobaoQianniuKefuevalGetRequest{
    return &TaobaoQianniuKefuevalGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQianniuKefuevalGetRequest) GetApiMethodName() string {
    return "taobao.qianniu.kefueval.get"
}

func (r TaobaoQianniuKefuevalGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQianniuKefuevalGetRequest) SetQueryIds(queryIds string) error {
    r.queryIds = queryIds
    r.Set("query_ids", queryIds)
    return nil
}

func (r TaobaoQianniuKefuevalGetRequest) GetQueryIds() string {
    return r.queryIds
}

func (r *TaobaoQianniuKefuevalGetRequest) SetBtime(btime string) error {
    r.btime = btime
    r.Set("btime", btime)
    return nil
}

func (r TaobaoQianniuKefuevalGetRequest) GetBtime() string {
    return r.btime
}

func (r *TaobaoQianniuKefuevalGetRequest) SetEtime(etime string) error {
    r.etime = etime
    r.Set("etime", etime)
    return nil
}

func (r TaobaoQianniuKefuevalGetRequest) GetEtime() string {
    return r.etime
}

