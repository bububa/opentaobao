package baichuan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
用户 APIResponse
alibaba.baichuan.ctg.user.relation

提供给优酷查询道长和淘宝账户的绑定关系
*/
type AlibabaBaichuanCtgUserRelationAPIResponse struct {
    model.CommonResponse
    AlibabaBaichuanCtgUserRelationResponse
}

type AlibabaBaichuanCtgUserRelationResponse struct {
    XMLName xml.Name `xml:"alibaba_baichuan_ctg_user_relation_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回的整体结果
    
    Result   *AlibabaBaichuanCtgUserRelationResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
