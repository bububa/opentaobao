package icburfq

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
是否已读RFQ APIRequest
alibaba.icbu.rfq.read

是否已读RFQ
*/
type AlibabaIcbuRfqReadRequest struct {
    model.Params

    // 查询RFQID列表
    rfqIdList   []string 

}

func NewAlibabaIcbuRfqReadRequest() *AlibabaIcbuRfqReadRequest{
    return &AlibabaIcbuRfqReadRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIcbuRfqReadRequest) GetApiMethodName() string {
    return "alibaba.icbu.rfq.read"
}

func (r AlibabaIcbuRfqReadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIcbuRfqReadRequest) SetRfqIdList(rfqIdList []string) error {
    r.rfqIdList = rfqIdList
    r.Set("rfq_id_list", rfqIdList)
    return nil
}

func (r AlibabaIcbuRfqReadRequest) GetRfqIdList() []string {
    return r.rfqIdList
}

