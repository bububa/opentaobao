package openmall

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
OpenMall退款图片上传 APIResponse
taobao.openmall.refund.image.upload

OpenMall退款图片上传
*/
type TaobaoOpenmallRefundImageUploadAPIResponse struct {
    model.CommonResponse
    TaobaoOpenmallRefundImageUploadResponse
}

type TaobaoOpenmallRefundImageUploadResponse struct {
    XMLName xml.Name `xml:"openmall_refund_image_upload_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 图片上传对应Token，用于提交留言接口
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
