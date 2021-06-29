package wenyuvideo

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
只看TA APIResponse
youku.wenyuvideo.seeta.get

只看Ta对外输出
*/
type YoukuWenyuvideoSeetaGetAPIResponse struct {
    model.CommonResponse
    YoukuWenyuvideoSeetaGetResponse
}

type YoukuWenyuvideoSeetaGetResponse struct {
    XMLName xml.Name `xml:"youku_wenyuvideo_seeta_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *YoukuWenyuvideoSeetaGetResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
