package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建及绑定互动实例 API返回值 
alibaba.interact.isvadmin.bind

创建互动实例，并绑定奖池
*/
type AlibabaInteractIsvadminBindAPIResponse struct {
    model.CommonResponse
    AlibabaInteractIsvadminBindAPIResponseModel
}

// 创建及绑定互动实例 成功返回结果
type AlibabaInteractIsvadminBindAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_interact_isvadmin_bind_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回创建并且绑定成功的互动实例
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
