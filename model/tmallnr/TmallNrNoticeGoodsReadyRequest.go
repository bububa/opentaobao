package tmallnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
同步天猫配送人员信息 APIRequest
tmall.nr.notice.goods.ready

接收商家的配送人员信息，和第三公司信息及提货码
*/
type TmallNrNoticeGoodsReadyRequest struct {
    model.Params

    // 入参
    param0   *NrZqsGoodsReadyReqDto 

}

func NewTmallNrNoticeGoodsReadyRequest() *TmallNrNoticeGoodsReadyRequest{
    return &TmallNrNoticeGoodsReadyRequest{
        Params: model.NewParams(),
    }
}

func (r TmallNrNoticeGoodsReadyRequest) GetApiMethodName() string {
    return "tmall.nr.notice.goods.ready"
}

func (r TmallNrNoticeGoodsReadyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallNrNoticeGoodsReadyRequest) SetParam0(param0 *NrZqsGoodsReadyReqDto) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r TmallNrNoticeGoodsReadyRequest) GetParam0() *NrZqsGoodsReadyReqDto {
    return r.param0
}

