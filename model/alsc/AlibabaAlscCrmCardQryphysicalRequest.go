package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询物理卡 API请求
alibaba.alsc.crm.card.qryphysical

查询物理卡
*/
type AlibabaAlscCrmCardQryphysicalRequest struct {
    model.Params
    // 入参
    _paramQueryPhyCardOpenReq   *QueryPhyCardOpenReq
}

// 初始化AlibabaAlscCrmCardQryphysicalRequest对象
func NewAlibabaAlscCrmCardQryphysicalRequest() *AlibabaAlscCrmCardQryphysicalRequest{
    return &AlibabaAlscCrmCardQryphysicalRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCardQryphysicalRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.card.qryphysical"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCardQryphysicalRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamQueryPhyCardOpenReq Setter
// 入参
func (r *AlibabaAlscCrmCardQryphysicalRequest) SetParamQueryPhyCardOpenReq(_paramQueryPhyCardOpenReq *QueryPhyCardOpenReq) error {
    r._paramQueryPhyCardOpenReq = _paramQueryPhyCardOpenReq
    r.Set("param_query_phy_card_open_req", _paramQueryPhyCardOpenReq)
    return nil
}

// ParamQueryPhyCardOpenReq Getter
func (r AlibabaAlscCrmCardQryphysicalRequest) GetParamQueryPhyCardOpenReq() *QueryPhyCardOpenReq {
    return r._paramQueryPhyCardOpenReq
}
