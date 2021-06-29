package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
解除码的关联关系 API请求
alibaba.alihealth.tracecodeseller.code.relation.codeantiactive

解除码的关联关系
*/
type AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveRequest struct {
    model.Params
    // 顶层码
    _topCode   string
    // 淘宝名
    _tbUserId   string
    // 企业id
    _entInfoId   int64
}

// 初始化AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveRequest对象
func NewAlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveRequest() *AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveRequest{
    return &AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveRequest) GetApiMethodName() string {
    return "alibaba.alihealth.tracecodeseller.code.relation.codeantiactive"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TopCode Setter
// 顶层码
func (r *AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveRequest) SetTopCode(_topCode string) error {
    r._topCode = _topCode
    r.Set("top_code", _topCode)
    return nil
}

// TopCode Getter
func (r AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveRequest) GetTopCode() string {
    return r._topCode
}
// TbUserId Setter
// 淘宝名
func (r *AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveRequest) SetTbUserId(_tbUserId string) error {
    r._tbUserId = _tbUserId
    r.Set("tb_user_id", _tbUserId)
    return nil
}

// TbUserId Getter
func (r AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveRequest) GetTbUserId() string {
    return r._tbUserId
}
// EntInfoId Setter
// 企业id
func (r *AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveRequest) SetEntInfoId(_entInfoId int64) error {
    r._entInfoId = _entInfoId
    r.Set("ent_info_id", _entInfoId)
    return nil
}

// EntInfoId Getter
func (r AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveRequest) GetEntInfoId() int64 {
    return r._entInfoId
}
