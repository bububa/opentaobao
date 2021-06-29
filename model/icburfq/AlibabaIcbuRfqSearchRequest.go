package icburfq

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询RFQ APIRequest
alibaba.icbu.rfq.search

用于查询RFQ的信息
*/
type AlibabaIcbuRfqSearchRequest struct {
    model.Params

    // 验证
    md5key   string 

    // 查询条件
    cond   *RfqRequestSearchCondDto 

}

func NewAlibabaIcbuRfqSearchRequest() *AlibabaIcbuRfqSearchRequest{
    return &AlibabaIcbuRfqSearchRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIcbuRfqSearchRequest) GetApiMethodName() string {
    return "alibaba.icbu.rfq.search"
}

func (r AlibabaIcbuRfqSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIcbuRfqSearchRequest) SetMd5key(md5key string) error {
    r.md5key = md5key
    r.Set("md5key", md5key)
    return nil
}

func (r AlibabaIcbuRfqSearchRequest) GetMd5key() string {
    return r.md5key
}

func (r *AlibabaIcbuRfqSearchRequest) SetCond(cond *RfqRequestSearchCondDto) error {
    r.cond = cond
    r.Set("cond", cond)
    return nil
}

func (r AlibabaIcbuRfqSearchRequest) GetCond() *RfqRequestSearchCondDto {
    return r.cond
}

