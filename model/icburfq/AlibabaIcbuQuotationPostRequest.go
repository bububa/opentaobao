package icburfq

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商提交报价 APIRequest
alibaba.icbu.quotation.post

供应商对RFQ进行报价
*/
type AlibabaIcbuQuotationPostRequest struct {
    model.Params

    // 验证
    md5key   string 

    // 报价DTO
    dto   *RfqQuotationRemoteDto 

}

func NewAlibabaIcbuQuotationPostRequest() *AlibabaIcbuQuotationPostRequest{
    return &AlibabaIcbuQuotationPostRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIcbuQuotationPostRequest) GetApiMethodName() string {
    return "alibaba.icbu.quotation.post"
}

func (r AlibabaIcbuQuotationPostRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIcbuQuotationPostRequest) SetMd5key(md5key string) error {
    r.md5key = md5key
    r.Set("md5key", md5key)
    return nil
}

func (r AlibabaIcbuQuotationPostRequest) GetMd5key() string {
    return r.md5key
}

func (r *AlibabaIcbuQuotationPostRequest) SetDto(dto *RfqQuotationRemoteDto) error {
    r.dto = dto
    r.Set("dto", dto)
    return nil
}

func (r AlibabaIcbuQuotationPostRequest) GetDto() *RfqQuotationRemoteDto {
    return r.dto
}

