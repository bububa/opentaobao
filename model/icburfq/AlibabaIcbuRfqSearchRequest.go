package icburfq

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询RFQ API请求
alibaba.icbu.rfq.search

用于查询RFQ的信息
*/
type AlibabaIcbuRfqSearchRequest struct {
    model.Params
    // 验证
    _md5key   string
    // 查询条件
    _cond   *RfqRequestSearchCondDTO
}

// 初始化AlibabaIcbuRfqSearchRequest对象
func NewAlibabaIcbuRfqSearchRequest() *AlibabaIcbuRfqSearchRequest{
    return &AlibabaIcbuRfqSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIcbuRfqSearchRequest) GetApiMethodName() string {
    return "alibaba.icbu.rfq.search"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIcbuRfqSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Md5key Setter
// 验证
func (r *AlibabaIcbuRfqSearchRequest) SetMd5key(_md5key string) error {
    r._md5key = _md5key
    r.Set("md5key", _md5key)
    return nil
}

// Md5key Getter
func (r AlibabaIcbuRfqSearchRequest) GetMd5key() string {
    return r._md5key
}
// Cond Setter
// 查询条件
func (r *AlibabaIcbuRfqSearchRequest) SetCond(_cond *RfqRequestSearchCondDTO) error {
    r._cond = _cond
    r.Set("cond", _cond)
    return nil
}

// Cond Getter
func (r AlibabaIcbuRfqSearchRequest) GetCond() *RfqRequestSearchCondDTO {
    return r._cond
}
