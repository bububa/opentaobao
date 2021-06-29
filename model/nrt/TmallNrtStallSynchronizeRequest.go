package nrt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
摊位信息同步 APIRequest
tmall.nrt.stall.synchronize

摊位信息同步
*/
type TmallNrtStallSynchronizeRequest struct {
    model.Params

    // 参数对象
    stall   *NrtStoreDTO 

}

func NewTmallNrtStallSynchronizeRequest() *TmallNrtStallSynchronizeRequest{
    return &TmallNrtStallSynchronizeRequest{
        Params: model.NewParams(),
    }
}

func (r TmallNrtStallSynchronizeRequest) GetApiMethodName() string {
    return "tmall.nrt.stall.synchronize"
}

func (r TmallNrtStallSynchronizeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallNrtStallSynchronizeRequest) SetStall(stall *NrtStoreDTO) error {
    r.stall = stall
    r.Set("stall", stall)
    return nil
}

func (r TmallNrtStallSynchronizeRequest) GetStall() *NrtStoreDTO {
    return r.stall
}

