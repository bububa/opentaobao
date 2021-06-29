package alitrippoi

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
穷游内容写入接口 APIResponse
alitrip.platform.content.raw.add

穷游内容写入飞猪接口
*/
type AlitripPlatformContentRawAddAPIResponse struct {
    model.CommonResponse
    AlitripPlatformContentRawAddResponse
}

type AlitripPlatformContentRawAddResponse struct {
    XMLName xml.Name `xml:"alitrip_platform_content_raw_add_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回包装类
    
    Result   *BaseResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
