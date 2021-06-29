package tmallnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取消上门揽件 APIRequest
tmall.nr.fulfill.cancel

新零售门店业务取消上门揽件
*/
type TmallNrFulfillCancelRequest struct {
    model.Params

    // 入参
    req   *NrCancelFulfillReqDto 

}

func NewTmallNrFulfillCancelRequest() *TmallNrFulfillCancelRequest{
    return &TmallNrFulfillCancelRequest{
        Params: model.NewParams(),
    }
}

func (r TmallNrFulfillCancelRequest) GetApiMethodName() string {
    return "tmall.nr.fulfill.cancel"
}

func (r TmallNrFulfillCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallNrFulfillCancelRequest) SetReq(req *NrCancelFulfillReqDto) error {
    r.req = req
    r.Set("req", req)
    return nil
}

func (r TmallNrFulfillCancelRequest) GetReq() *NrCancelFulfillReqDto {
    return r.req
}

