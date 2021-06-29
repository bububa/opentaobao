package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除项目 APIResponse
alitrip.btrip.project.delete

删除项目
*/
type AlitripBtripProjectDeleteAPIResponse struct {
    model.CommonResponse
    AlitripBtripProjectDeleteResponse
}

type AlitripBtripProjectDeleteResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_project_delete_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
    // 结果码
    
    ResultCode   int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // 结果信息
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
}
