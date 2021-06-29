package wdkitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
图片上传接口 API返回值 
alibaba.wdk.picture.upload

上传图片
*/
type AlibabaWdkPictureUploadAPIResponse struct {
    model.CommonResponse
    AlibabaWdkPictureUploadResponse
}

// 图片上传接口 成功返回结果
type AlibabaWdkPictureUploadResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_picture_upload_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // apiresult
    Result   *AlibabaWdkPictureUploadApiResult `json:"result,omitempty" xml:"result,omitempty"`
}
