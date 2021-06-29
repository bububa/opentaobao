package alihealthcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
点击标签后排序接口 APIRequest
alibaba.alihealth.pregnancy.taginfo.addrank

备孕管理--点击标签后排序接口
*/
type AlibabaAlihealthPregnancyTaginfoAddrankRequest struct {
    model.Params

    // 用户id
    userId   int64 

    // 标签编码，例如备孕标签为5122
    tagCode   string 

}

func NewAlibabaAlihealthPregnancyTaginfoAddrankRequest() *AlibabaAlihealthPregnancyTaginfoAddrankRequest{
    return &AlibabaAlihealthPregnancyTaginfoAddrankRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthPregnancyTaginfoAddrankRequest) GetApiMethodName() string {
    return "alibaba.alihealth.pregnancy.taginfo.addrank"
}

func (r AlibabaAlihealthPregnancyTaginfoAddrankRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthPregnancyTaginfoAddrankRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlibabaAlihealthPregnancyTaginfoAddrankRequest) GetUserId() int64 {
    return r.userId
}

func (r *AlibabaAlihealthPregnancyTaginfoAddrankRequest) SetTagCode(tagCode string) error {
    r.tagCode = tagCode
    r.Set("tag_code", tagCode)
    return nil
}

func (r AlibabaAlihealthPregnancyTaginfoAddrankRequest) GetTagCode() string {
    return r.tagCode
}

