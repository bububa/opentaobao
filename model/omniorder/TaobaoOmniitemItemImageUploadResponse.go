package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
全渠道商品上传图片 APIResponse
taobao.omniitem.item.image.upload

全渠道商品上传图片
*/
type TaobaoOmniitemItemImageUploadAPIResponse struct {
    model.CommonResponse
    TaobaoOmniitemItemImageUploadResponse
}

type TaobaoOmniitemItemImageUploadResponse struct {
    XMLName xml.Name `xml:"omniitem_item_image_upload_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TaobaoOmniitemItemImageUploadResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
