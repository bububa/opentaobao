package baichuanctg

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
微博输出头条数据 APIRequest
alibaba.baichuan.ctg.toutiao.content

百川头条内容获取
*/
type AlibabaBaichuanCtgToutiaoContentRequest struct {
    model.Params

    // param0
    param0   *CtgRequest 

}

func NewAlibabaBaichuanCtgToutiaoContentRequest() *AlibabaBaichuanCtgToutiaoContentRequest{
    return &AlibabaBaichuanCtgToutiaoContentRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaBaichuanCtgToutiaoContentRequest) GetApiMethodName() string {
    return "alibaba.baichuan.ctg.toutiao.content"
}

func (r AlibabaBaichuanCtgToutiaoContentRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaBaichuanCtgToutiaoContentRequest) SetParam0(param0 *CtgRequest) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlibabaBaichuanCtgToutiaoContentRequest) GetParam0() *CtgRequest {
    return r.param0
}

