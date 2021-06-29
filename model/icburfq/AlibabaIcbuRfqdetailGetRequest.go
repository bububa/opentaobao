package icburfq

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取RFQ详情 API请求
alibaba.icbu.rfqdetail.get

查看RFQ的详情信息
*/
type AlibabaIcbuRfqdetailGetRequest struct {
    model.Params
    // 验证
    _md5key   string
    // 查询RFQ详情DTO
    _rfqQueryDto   *RfqDetailSearchQueryDTO
}

// 初始化AlibabaIcbuRfqdetailGetRequest对象
func NewAlibabaIcbuRfqdetailGetRequest() *AlibabaIcbuRfqdetailGetRequest{
    return &AlibabaIcbuRfqdetailGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIcbuRfqdetailGetRequest) GetApiMethodName() string {
    return "alibaba.icbu.rfqdetail.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIcbuRfqdetailGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Md5key Setter
// 验证
func (r *AlibabaIcbuRfqdetailGetRequest) SetMd5key(_md5key string) error {
    r._md5key = _md5key
    r.Set("md5key", _md5key)
    return nil
}

// Md5key Getter
func (r AlibabaIcbuRfqdetailGetRequest) GetMd5key() string {
    return r._md5key
}
// RfqQueryDto Setter
// 查询RFQ详情DTO
func (r *AlibabaIcbuRfqdetailGetRequest) SetRfqQueryDto(_rfqQueryDto *RfqDetailSearchQueryDTO) error {
    r._rfqQueryDto = _rfqQueryDto
    r.Set("rfq_query_dto", _rfqQueryDto)
    return nil
}

// RfqQueryDto Getter
func (r AlibabaIcbuRfqdetailGetRequest) GetRfqQueryDto() *RfqDetailSearchQueryDTO {
    return r._rfqQueryDto
}
