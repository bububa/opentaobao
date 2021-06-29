package ihome

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
售后上门信息同步 API返回值 
alibaba.ihome.ctom.postsale.onsite.sync

用于三维家同步售后单上门人员和时间信息
*/
type AlibabaIhomeCtomPostsaleOnsiteSyncAPIResponse struct {
    model.CommonResponse
    AlibabaIhomeCtomPostsaleOnsiteSyncResponse
}

// 售后上门信息同步 成功返回结果
type AlibabaIhomeCtomPostsaleOnsiteSyncResponse struct {
    XMLName xml.Name `xml:"alibaba_ihome_ctom_postsale_onsite_sync_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // true
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`
}
