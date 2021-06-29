package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
导购员判断 APIRequest
taobao.omniorder.print.sale.judge

用于判断当前子账号是否导购员
*/
type TaobaoOmniorderPrintSaleJudgeRequest struct {
    model.Params

    // 用户子账号ID
    subUid   int64 

}

func NewTaobaoOmniorderPrintSaleJudgeRequest() *TaobaoOmniorderPrintSaleJudgeRequest{
    return &TaobaoOmniorderPrintSaleJudgeRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOmniorderPrintSaleJudgeRequest) GetApiMethodName() string {
    return "taobao.omniorder.print.sale.judge"
}

func (r TaobaoOmniorderPrintSaleJudgeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOmniorderPrintSaleJudgeRequest) SetSubUid(subUid int64) error {
    r.subUid = subUid
    r.Set("sub_uid", subUid)
    return nil
}

func (r TaobaoOmniorderPrintSaleJudgeRequest) GetSubUid() int64 {
    return r.subUid
}

