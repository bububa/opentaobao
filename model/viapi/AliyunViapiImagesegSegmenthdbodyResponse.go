package viapi

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
高清人体分割 API返回值 
aliyun.viapi.imageseg.segmenthdbody

对输入图像中包含的人进行分割，输出结果透明图。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
*/
type AliyunViapiImagesegSegmenthdbodyAPIResponse struct {
    model.CommonResponse
    AliyunViapiImagesegSegmenthdbodyResponse
}

// 高清人体分割 成功返回结果
type AliyunViapiImagesegSegmenthdbodyResponse struct {
    XMLName xml.Name `xml:"aliyun_viapi_imageseg_segmenthdbody_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 请求ID
    TaobaoRequestId   string `json:"taobao_request_id,omitempty" xml:"taobao_request_id,omitempty"`
    // 系统自动生成
    Data   *AliyunViapiImagesegSegmenthdbodyData `json:"data,omitempty" xml:"data,omitempty"`
}
