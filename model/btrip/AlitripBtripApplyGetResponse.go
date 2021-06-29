package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取单个审批单 API返回值 
alitrip.btrip.apply.get

获取单个审批单的详情数据
*/
type AlitripBtripApplyGetAPIResponse struct {
    model.CommonResponse
    AlitripBtripApplyGetResponse
}

// 获取单个审批单 成功返回结果
type AlitripBtripApplyGetResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_apply_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *BtriphomeResultSupport `json:"result,omitempty" xml:"result,omitempty"`
}
