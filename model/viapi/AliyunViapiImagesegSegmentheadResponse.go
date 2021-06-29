package viapi

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
头像分割 API返回值 
aliyun.viapi.imageseg.segmenthead

输入一张图片，对图中人头区域进行抠图解析，输出人头png透明图。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
*/
type AliyunViapiImagesegSegmentheadAPIResponse struct {
    model.CommonResponse
    AliyunViapiImagesegSegmentheadResponse
}

// 头像分割 成功返回结果
type AliyunViapiImagesegSegmentheadResponse struct {
    XMLName xml.Name `xml:"aliyun_viapi_imageseg_segmenthead_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 请求ID
    TaobaoRequestId   string `json:"taobao_request_id,omitempty" xml:"taobao_request_id,omitempty"`
    // 系统自动生成
    Data   *AliyunViapiImagesegSegmentheadData `json:"data,omitempty" xml:"data,omitempty"`
}
