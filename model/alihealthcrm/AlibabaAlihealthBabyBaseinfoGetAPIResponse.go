package alihealthcrm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
三方从我们这获取宝宝信息 API返回值 
alibaba.alihealth.baby.baseinfo.get

三方从我们这获取宝宝信息
*/
type AlibabaAlihealthBabyBaseinfoGetAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthBabyBaseinfoGetAPIResponseModel
}

// 三方从我们这获取宝宝信息 成功返回结果
type AlibabaAlihealthBabyBaseinfoGetAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_alihealth_baby_baseinfo_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *BabyInfoResult `json:"result,omitempty" xml:"result,omitempty"`
}
