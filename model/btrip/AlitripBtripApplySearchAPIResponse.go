package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
搜索审批单 API返回值 
alitrip.btrip.apply.search

外部企业调用获取本企业审批单列表数据
*/
type AlitripBtripApplySearchAPIResponse struct {
    model.CommonResponse
    AlitripBtripApplySearchAPIResponseModel
}

// 搜索审批单 成功返回结果
type AlitripBtripApplySearchAPIResponseModel struct {
    XMLName xml.Name `xml:"alitrip_btrip_apply_search_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *BtriphomeResultSupport `json:"result,omitempty" xml:"result,omitempty"`
}
