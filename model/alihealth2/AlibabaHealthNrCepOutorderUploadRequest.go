package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
线上订单收货验收单、出入库单据生成接口 APIRequest
alibaba.health.nr.cep.outorder.upload

线上订单收货验收单、出入库单据生成接口
*/
type AlibabaHealthNrCepOutorderUploadRequest struct {
    model.Params

    // 出库单对象
    topWarOutDto   *TopWarOutDTO 

}

func NewAlibabaHealthNrCepOutorderUploadRequest() *AlibabaHealthNrCepOutorderUploadRequest{
    return &AlibabaHealthNrCepOutorderUploadRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaHealthNrCepOutorderUploadRequest) GetApiMethodName() string {
    return "alibaba.health.nr.cep.outorder.upload"
}

func (r AlibabaHealthNrCepOutorderUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaHealthNrCepOutorderUploadRequest) SetTopWarOutDto(topWarOutDto *TopWarOutDTO) error {
    r.topWarOutDto = topWarOutDto
    r.Set("top_war_out_dto", topWarOutDto)
    return nil
}

func (r AlibabaHealthNrCepOutorderUploadRequest) GetTopWarOutDto() *TopWarOutDTO {
    return r.topWarOutDto
}

