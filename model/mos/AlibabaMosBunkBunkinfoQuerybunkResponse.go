package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据合同号查询铺位信息 APIResponse
alibaba.mos.bunk.bunkinfo.querybunk

根据合同号查询铺位信息
*/
type AlibabaMosBunkBunkinfoQuerybunkAPIResponse struct {
    model.CommonResponse
    AlibabaMosBunkBunkinfoQuerybunkResponse
}

type AlibabaMosBunkBunkinfoQuerybunkResponse struct {
    XMLName xml.Name `xml:"alibaba_mos_bunk_bunkinfo_querybunk_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *AlibabaMosBunkBunkinfoQuerybunkResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
