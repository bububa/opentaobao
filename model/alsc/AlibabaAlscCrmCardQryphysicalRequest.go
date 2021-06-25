package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询物理卡 APIRequest
alibaba.alsc.crm.card.qryphysical

查询物理卡
*/
type AlibabaAlscCrmCardQryphysicalRequest struct {
    model.Params

    // 入参
    paramQueryPhyCardOpenReq   *QueryPhyCardOpenReq 

}

func NewAlibabaAlscCrmCardQryphysicalRequest() *AlibabaAlscCrmCardQryphysicalRequest{
    return &AlibabaAlscCrmCardQryphysicalRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmCardQryphysicalRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.card.qryphysical"
}

func (r AlibabaAlscCrmCardQryphysicalRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmCardQryphysicalRequest) SetParamQueryPhyCardOpenReq(paramQueryPhyCardOpenReq *QueryPhyCardOpenReq) error {
    r.paramQueryPhyCardOpenReq = paramQueryPhyCardOpenReq
    r.Set("param_query_phy_card_open_req", paramQueryPhyCardOpenReq)
    return nil
}

func (r AlibabaAlscCrmCardQryphysicalRequest) GetParamQueryPhyCardOpenReq() *QueryPhyCardOpenReq {
    return r.paramQueryPhyCardOpenReq
}

