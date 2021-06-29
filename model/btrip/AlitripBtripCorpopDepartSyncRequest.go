package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
外部部门同步 APIRequest
alitrip.btrip.corpop.depart.sync

同步外部平台部门信息至商旅内部
*/
type AlitripBtripCorpopDepartSyncRequest struct {
    model.Params

    // 同步部门请求
    rq   *BtripDepartSyncRq 

}

func NewAlitripBtripCorpopDepartSyncRequest() *AlitripBtripCorpopDepartSyncRequest{
    return &AlitripBtripCorpopDepartSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripCorpopDepartSyncRequest) GetApiMethodName() string {
    return "alitrip.btrip.corpop.depart.sync"
}

func (r AlitripBtripCorpopDepartSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripCorpopDepartSyncRequest) SetRq(rq *BtripDepartSyncRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripCorpopDepartSyncRequest) GetRq() *BtripDepartSyncRq {
    return r.rq
}

