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
type AlibabaIcbuRfqSearchAPIRequest struct {
    model.Params
    // 验证
    _md5key   string
    // 查询条件
    _cond   *RfqRequestSearchCondDto
}

// 初始化AlibabaIcbuRfqSearchAPIRequest对象
func NewAlibabaIcbuRfqSearchRequest() *AlibabaIcbuRfqSearchAPIRequest{
    return &AlibabaIcbuRfqSearchAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIcbuRfqSearchAPIRequest) GetApiMethodName() string {
    return "alibaba.icbu.rfq.search"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIcbuRfqSearchAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Md5key Setter
// 验证
func (r *AlibabaIcbuRfqSearchAPIRequest) SetMd5key(_md5key string) error {
    r._md5key = _md5key
    r.Set("md5key", _md5key)
    return nil
}

// Md5key Getter
func (r AlibabaIcbuRfqSearchAPIRequest) GetMd5key() string {
    return r._md5key
}
// Cond Setter
// 查询条件
func (r *AlibabaIcbuRfqSearchAPIRequest) SetCond(_cond *RfqRequestSearchCondDto) error {
    r._cond = _cond
    r.Set("cond", _cond)
    return nil
}

// Cond Getter
func (r AlibabaIcbuRfqSearchAPIRequest) GetCond() *RfqRequestSearchCondDto {
    return r._cond
}
