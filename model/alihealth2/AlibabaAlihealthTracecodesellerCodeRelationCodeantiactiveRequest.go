package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
解除码的关联关系 APIRequest
alibaba.alihealth.tracecodeseller.code.relation.codeantiactive

解除码的关联关系
*/
type AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveRequest struct {
    model.Params

    // 顶层码
    topCode   string 

    // 淘宝名
    tbUserId   string 

    // 企业id
    entInfoId   int64 

}

func NewAlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveRequest() *AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveRequest{
    return &AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveRequest) GetApiMethodName() string {
    return "alibaba.alihealth.tracecodeseller.code.relation.codeantiactive"
}

func (r AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveRequest) SetTopCode(topCode string) error {
    r.topCode = topCode
    r.Set("top_code", topCode)
    return nil
}

func (r AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveRequest) GetTopCode() string {
    return r.topCode
}

func (r *AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveRequest) SetTbUserId(tbUserId string) error {
    r.tbUserId = tbUserId
    r.Set("tb_user_id", tbUserId)
    return nil
}

func (r AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveRequest) GetTbUserId() string {
    return r.tbUserId
}

func (r *AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveRequest) SetEntInfoId(entInfoId int64) error {
    r.entInfoId = entInfoId
    r.Set("ent_info_id", entInfoId)
    return nil
}

func (r AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveRequest) GetEntInfoId() int64 {
    return r.entInfoId
}

