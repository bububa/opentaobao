package icburfq

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取RFQ详情 APIRequest
alibaba.icbu.rfqdetail.get

查看RFQ的详情信息
*/
type AlibabaIcbuRfqdetailGetRequest struct {
    model.Params

    // 验证
    md5key   string 

    // 查询RFQ详情DTO
    rfqQueryDto   *RfqDetailSearchQueryDto 

}

func NewAlibabaIcbuRfqdetailGetRequest() *AlibabaIcbuRfqdetailGetRequest{
    return &AlibabaIcbuRfqdetailGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIcbuRfqdetailGetRequest) GetApiMethodName() string {
    return "alibaba.icbu.rfqdetail.get"
}

func (r AlibabaIcbuRfqdetailGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIcbuRfqdetailGetRequest) SetMd5key(md5key string) error {
    r.md5key = md5key
    r.Set("md5key", md5key)
    return nil
}

func (r AlibabaIcbuRfqdetailGetRequest) GetMd5key() string {
    return r.md5key
}

func (r *AlibabaIcbuRfqdetailGetRequest) SetRfqQueryDto(rfqQueryDto *RfqDetailSearchQueryDto) error {
    r.rfqQueryDto = rfqQueryDto
    r.Set("rfq_query_dto", rfqQueryDto)
    return nil
}

func (r AlibabaIcbuRfqdetailGetRequest) GetRfqQueryDto() *RfqDetailSearchQueryDto {
    return r.rfqQueryDto
}

