package cainiaohandover

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改交接单 API返回值 
cainiao.global.handover.update

提供给ISV通过该接口修改交接单
*/
type CainiaoGlobalHandoverUpdateAPIResponse struct {
    model.CommonResponse
    CainiaoGlobalHandoverUpdateAPIResponseModel
}

// 修改交接单 成功返回结果
type CainiaoGlobalHandoverUpdateAPIResponseModel struct {
    XMLName xml.Name `xml:"cainiao_global_handover_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 请求结果
    Result   *HsfResult `json:"result,omitempty" xml:"result,omitempty"`
}
