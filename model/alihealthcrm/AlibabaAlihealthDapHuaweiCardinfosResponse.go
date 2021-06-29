package alihealthcrm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
华为负一屏卡片查询 API返回值 
alibaba.alihealth.dap.huawei.cardinfos

医疗健康频道卡片华为负一屏
*/
type AlibabaAlihealthDapHuaweiCardinfosAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDapHuaweiCardinfosResponse
}

// 华为负一屏卡片查询 成功返回结果
type AlibabaAlihealthDapHuaweiCardinfosResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_dap_huawei_cardinfos_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 和三方交互最外层model对象
    Result   *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
