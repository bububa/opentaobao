package alihealthcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
点击标签后排序接口 API请求
alibaba.alihealth.pregnancy.taginfo.addrank

备孕管理--点击标签后排序接口
*/
type AlibabaAlihealthPregnancyTaginfoAddrankRequest struct {
    model.Params
    // 用户id
    _userId   int64
    // 标签编码，例如备孕标签为5122
    _tagCode   string
}

// 初始化AlibabaAlihealthPregnancyTaginfoAddrankRequest对象
func NewAlibabaAlihealthPregnancyTaginfoAddrankRequest() *AlibabaAlihealthPregnancyTaginfoAddrankRequest{
    return &AlibabaAlihealthPregnancyTaginfoAddrankRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthPregnancyTaginfoAddrankRequest) GetApiMethodName() string {
    return "alibaba.alihealth.pregnancy.taginfo.addrank"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthPregnancyTaginfoAddrankRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserId Setter
// 用户id
func (r *AlibabaAlihealthPregnancyTaginfoAddrankRequest) SetUserId(_userId int64) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r AlibabaAlihealthPregnancyTaginfoAddrankRequest) GetUserId() int64 {
    return r._userId
}
// TagCode Setter
// 标签编码，例如备孕标签为5122
func (r *AlibabaAlihealthPregnancyTaginfoAddrankRequest) SetTagCode(_tagCode string) error {
    r._tagCode = _tagCode
    r.Set("tag_code", _tagCode)
    return nil
}

// TagCode Getter
func (r AlibabaAlihealthPregnancyTaginfoAddrankRequest) GetTagCode() string {
    return r._tagCode
}
