package legalcase

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
委托 API请求
alibaba.legal.case.entrust.get

获取委托案件的基本信息
*/
type AlibabaLegalCaseEntrustGetRequest struct {
    model.Params
    // 委托id
    _entrustId   int64
}

// 初始化AlibabaLegalCaseEntrustGetRequest对象
func NewAlibabaLegalCaseEntrustGetRequest() *AlibabaLegalCaseEntrustGetRequest{
    return &AlibabaLegalCaseEntrustGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLegalCaseEntrustGetRequest) GetApiMethodName() string {
    return "alibaba.legal.case.entrust.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLegalCaseEntrustGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// EntrustId Setter
// 委托id
func (r *AlibabaLegalCaseEntrustGetRequest) SetEntrustId(_entrustId int64) error {
    r._entrustId = _entrustId
    r.Set("entrust_id", _entrustId)
    return nil
}

// EntrustId Getter
func (r AlibabaLegalCaseEntrustGetRequest) GetEntrustId() int64 {
    return r._entrustId
}
