package oversea

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
汇率信息获取 API请求
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

// 初始化AlibabaOverseaExchagerateGetRequest对象
func NewAlibabaOverseaExchagerateGetRequest() *AlibabaOverseaExchagerateGetRequest{
    return &AlibabaOverseaExchagerateGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaOverseaExchagerateGetRequest) GetApiMethodName() string {
    return "alibaba.oversea.exchagerate.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaOverseaExchagerateGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizCode Setter
// 业务类型
func (r *AlibabaOverseaExchagerateGetRequest) SetBizCode(bizCode string) error {
    r.bizCode = bizCode
    r.Set("biz_code", bizCode)
    return nil
}

// BizCode Getter
func (r AlibabaOverseaExchagerateGetRequest) GetBizCode() string {
    return r.bizCode
}
// BaseCode Setter
// 原始币种
func (r *AlibabaOverseaExchagerateGetRequest) SetBaseCode(baseCode string) error {
    r.baseCode = baseCode
    r.Set("base_code", baseCode)
    return nil
}

// BaseCode Getter
func (r AlibabaOverseaExchagerateGetRequest) GetBaseCode() string {
    return r.baseCode
}
// TargetCode Setter
// 目标币种
func (r *AlibabaOverseaExchagerateGetRequest) SetTargetCode(targetCode string) error {
    r.targetCode = targetCode
    r.Set("target_code", targetCode)
    return nil
}

// TargetCode Getter
func (r AlibabaOverseaExchagerateGetRequest) GetTargetCode() string {
    return r.targetCode
}
