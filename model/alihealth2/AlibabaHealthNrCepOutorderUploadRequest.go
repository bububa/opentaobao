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
type AlibabaHealthNrCepOutorderUploadAPIRequest struct {
    model.Params
    // 出库单对象
    _topWarOutDto   *TopWarOutDTO
}

// 初始化AlibabaHealthNrCepOutorderUploadAPIRequest对象
func NewAlibabaHealthNrCepOutorderUploadRequest() *AlibabaHealthNrCepOutorderUploadAPIRequest{
    return &AlibabaHealthNrCepOutorderUploadAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHealthNrCepOutorderUploadAPIRequest) GetApiMethodName() string {
    return "alibaba.health.nr.cep.outorder.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHealthNrCepOutorderUploadAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TopWarOutDto Setter
// 出库单对象
func (r *AlibabaHealthNrCepOutorderUploadAPIRequest) SetTopWarOutDto(_topWarOutDto *TopWarOutDTO) error {
    r._topWarOutDto = _topWarOutDto
    r.Set("top_war_out_dto", _topWarOutDto)
    return nil
}

// TopWarOutDto Getter
func (r AlibabaHealthNrCepOutorderUploadAPIRequest) GetTopWarOutDto() *TopWarOutDTO {
    return r._topWarOutDto
}
