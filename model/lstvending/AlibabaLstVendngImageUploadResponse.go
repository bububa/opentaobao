package lstvending

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
售货机商品图片上传 APIResponse
alibaba.lst.vendng.image.upload

零售通自动售货机商品图片上传接口，主要为ISV厂商提供图片同步的通道，从而建立统一的商品图片库。
*/
type AlibabaLstVendngImageUploadAPIResponse struct {
    model.CommonResponse
    AlibabaLstVendngImageUploadResponse
}

type AlibabaLstVendngImageUploadResponse struct {
    XMLName xml.Name `xml:"alibaba_lst_vendng_image_upload_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果集
    
    Result   *AlibabaLstVendngImageUploadResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
