package icburfq

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
是否已读RFQ API请求
alibaba.icbu.rfq.read

是否已读RFQ
*/
type AlibabaIcbuRfqReadAPIRequest struct {
    model.Params
    // 查询RFQID列表
    _rfqIdList   []string
}

// 初始化AlibabaIcbuRfqReadAPIRequest对象
func NewAlibabaIcbuRfqReadRequest() *AlibabaIcbuRfqReadAPIRequest{
    return &AlibabaIcbuRfqReadAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIcbuRfqReadAPIRequest) GetApiMethodName() string {
    return "alibaba.icbu.rfq.read"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIcbuRfqReadAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RfqIdList Setter
// 查询RFQID列表
func (r *AlibabaIcbuRfqReadAPIRequest) SetRfqIdList(_rfqIdList []string) error {
    r._rfqIdList = _rfqIdList
    r.Set("rfq_id_list", _rfqIdList)
    return nil
}

// RfqIdList Getter
func (r AlibabaIcbuRfqReadAPIRequest) GetRfqIdList() []string {
    return r._rfqIdList
}
