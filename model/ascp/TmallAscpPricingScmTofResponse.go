package ascp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
TOF&SCM营销域对接-成本录入设置 APIResponse
tmall.ascp.pricing.scm.tof

TOF&SCM营销域对接-成本录入设置
*/
type TmallAscpPricingScmTofAPIResponse struct {
    model.CommonResponse
    TmallAscpPricingScmTofResponse
}

type TmallAscpPricingScmTofResponse struct {
    XMLName xml.Name `xml:"tmall_ascp_pricing_scm_tof_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 成功
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
    // reasonsForPartSucc
    
    ReasonsForPartSuccList   []string `json:"reasons_for_part_succ_list,omitempty" xml:"reasons_for_part_succ_list>string,omitempty"`
    
    
    // 成功
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
}
