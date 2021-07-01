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
type AlibabaIcbuQuotationPostAPIRequest struct {
    model.Params
    // 验证
    _md5key   string
    // 报价DTO
    _dto   *RfqQuotationRemoteDto
}

// 初始化AlibabaIcbuQuotationPostAPIRequest对象
func NewAlibabaIcbuQuotationPostRequest() *AlibabaIcbuQuotationPostAPIRequest{
    return &AlibabaIcbuQuotationPostAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIcbuQuotationPostAPIRequest) GetApiMethodName() string {
    return "alibaba.icbu.quotation.post"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIcbuQuotationPostAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Md5key Setter
// 验证
func (r *AlibabaIcbuQuotationPostAPIRequest) SetMd5key(_md5key string) error {
    r._md5key = _md5key
    r.Set("md5key", _md5key)
    return nil
}

// Md5key Getter
func (r AlibabaIcbuQuotationPostAPIRequest) GetMd5key() string {
    return r._md5key
}
// Dto Setter
// 报价DTO
func (r *AlibabaIcbuQuotationPostAPIRequest) SetDto(_dto *RfqQuotationRemoteDto) error {
    r._dto = _dto
    r.Set("dto", _dto)
    return nil
}

// Dto Getter
func (r AlibabaIcbuQuotationPostAPIRequest) GetDto() *RfqQuotationRemoteDto {
    return r._dto
}
