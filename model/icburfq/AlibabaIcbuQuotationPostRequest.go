package icburfq

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商提交报价 API请求
alibaba.icbu.quotation.post

供应商对RFQ进行报价
*/
type AlibabaIcbuQuotationPostRequest struct {
    model.Params
    // 验证
    _md5key   string
    // 报价DTO
    _dto   *RfqQuotationRemoteDto
}

// 初始化AlibabaIcbuQuotationPostRequest对象
func NewAlibabaIcbuQuotationPostRequest() *AlibabaIcbuQuotationPostRequest{
    return &AlibabaIcbuQuotationPostRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIcbuQuotationPostRequest) GetApiMethodName() string {
    return "alibaba.icbu.quotation.post"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIcbuQuotationPostRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Md5key Setter
// 验证
func (r *AlibabaIcbuQuotationPostRequest) SetMd5key(_md5key string) error {
    r._md5key = _md5key
    r.Set("md5key", _md5key)
    return nil
}

// Md5key Getter
func (r AlibabaIcbuQuotationPostRequest) GetMd5key() string {
    return r._md5key
}
// Dto Setter
// 报价DTO
func (r *AlibabaIcbuQuotationPostRequest) SetDto(_dto *RfqQuotationRemoteDto) error {
    r._dto = _dto
    r.Set("dto", _dto)
    return nil
}

// Dto Getter
func (r AlibabaIcbuQuotationPostRequest) GetDto() *RfqQuotationRemoteDto {
    return r._dto
}
