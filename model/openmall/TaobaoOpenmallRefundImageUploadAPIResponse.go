package openmall

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
OpenMall退款图片上传 API返回值 
taobao.openmall.refund.image.upload

OpenMall退款图片上传
*/
type TaobaoOpenmallRefundImageUploadAPIResponse struct {
    model.CommonResponse
    TaobaoOpenmallRefundImageUploadAPIResponseModel
}

// OpenMall退款图片上传 成功返回结果
type TaobaoOpenmallRefundImageUploadAPIResponseModel struct {
    XMLName xml.Name `xml:"openmall_refund_image_upload_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 图片上传对应Token，用于提交留言接口
    Result   string `json:"result,omitempty" xml:"result,omitempty"`
}
