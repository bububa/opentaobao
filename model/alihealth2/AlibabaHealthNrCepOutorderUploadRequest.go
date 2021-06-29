package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
线上订单收货验收单、出入库单据生成接口 API请求
alibaba.health.nr.cep.outorder.upload

线上订单收货验收单、出入库单据生成接口
*/
type AlibabaHealthNrCepOutorderUploadRequest struct {
    model.Params
    // 出库单对象
    topWarOutDto   *TopWarOutDTO
}

// 初始化AlibabaHealthNrCepOutorderUploadRequest对象
func NewAlibabaHealthNrCepOutorderUploadRequest() *AlibabaHealthNrCepOutorderUploadRequest{
    return &AlibabaHealthNrCepOutorderUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHealthNrCepOutorderUploadRequest) GetApiMethodName() string {
    return "alibaba.health.nr.cep.outorder.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHealthNrCepOutorderUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TopWarOutDto Setter
// 出库单对象
func (r *AlibabaHealthNrCepOutorderUploadRequest) SetTopWarOutDto(topWarOutDto *TopWarOutDTO) error {
    r.topWarOutDto = topWarOutDto
    r.Set("top_war_out_dto", topWarOutDto)
    return nil
}

// TopWarOutDto Getter
func (r AlibabaHealthNrCepOutorderUploadRequest) GetTopWarOutDto() *TopWarOutDTO {
    return r.topWarOutDto
}
