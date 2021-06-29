package oversea

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
汇率信息获取 APIRequest
alibaba.oversea.exchagerate.get

提供外部汇率查询接口
*/
type AlibabaOverseaExchagerateGetRequest struct {
    model.Params

    // 业务类型
    bizCode   string 

    // 原始币种
    baseCode   string 

    // 目标币种
    targetCode   string 

}

func NewAlibabaOverseaExchagerateGetRequest() *AlibabaOverseaExchagerateGetRequest{
    return &AlibabaOverseaExchagerateGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaOverseaExchagerateGetRequest) GetApiMethodName() string {
    return "alibaba.oversea.exchagerate.get"
}

func (r AlibabaOverseaExchagerateGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaOverseaExchagerateGetRequest) SetBizCode(bizCode string) error {
    r.bizCode = bizCode
    r.Set("biz_code", bizCode)
    return nil
}

func (r AlibabaOverseaExchagerateGetRequest) GetBizCode() string {
    return r.bizCode
}

func (r *AlibabaOverseaExchagerateGetRequest) SetBaseCode(baseCode string) error {
    r.baseCode = baseCode
    r.Set("base_code", baseCode)
    return nil
}

func (r AlibabaOverseaExchagerateGetRequest) GetBaseCode() string {
    return r.baseCode
}

func (r *AlibabaOverseaExchagerateGetRequest) SetTargetCode(targetCode string) error {
    r.targetCode = targetCode
    r.Set("target_code", targetCode)
    return nil
}

func (r AlibabaOverseaExchagerateGetRequest) GetTargetCode() string {
    return r.targetCode
}

