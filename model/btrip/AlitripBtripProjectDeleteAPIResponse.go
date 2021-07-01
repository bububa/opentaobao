package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除项目 API返回值 
alitrip.btrip.project.delete

删除项目
*/
type AlitripBtripProjectDeleteAPIResponse struct {
    model.CommonResponse
    AlitripBtripProjectDeleteAPIResponseModel
}

// 删除项目 成功返回结果
type AlitripBtripProjectDeleteAPIResponseModel struct {
    XMLName xml.Name `xml:"alitrip_btrip_project_delete_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`
    // 结果码
    ResultCode   int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
    // 结果信息
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
}
