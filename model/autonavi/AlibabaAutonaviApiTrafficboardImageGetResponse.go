package autonavi

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
交通看板-栅格情报获取 APIResponse
alibaba.autonavi.api.trafficboard.image.get

获取指定情报板ID的二进制数据（图片）
*/
type AlibabaAutonaviApiTrafficboardImageGetAPIResponse struct {
    model.CommonResponse
    AlibabaAutonaviApiTrafficboardImageGetResponse
}

type AlibabaAutonaviApiTrafficboardImageGetResponse struct {
    XMLName xml.Name `xml:"alibaba_autonavi_api_trafficboard_image_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 二进制图片流(png)
    
    RespResult   []*model.File `json:"resp_result,omitempty" xml:"resp_result>*model.File,omitempty"`
    
    
}
