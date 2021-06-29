package viapi

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
人脸比对1：1 API返回值 
aliyun.viapi.facebody.compareface

基于输入的两张图片，人脸比对服务可检测两张图片中的人脸，并挑选两张图片的最大人脸进行比较，判断是否是同一人；人脸比对服务还返回了这两个人脸的矩形框、比对的置信度，以及不同误识率的置信度阈值。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
*/
type AliyunViapiFacebodyComparefaceAPIResponse struct {
    model.CommonResponse
    AliyunViapiFacebodyComparefaceResponse
}

// 人脸比对1：1 成功返回结果
type AliyunViapiFacebodyComparefaceResponse struct {
    XMLName xml.Name `xml:"aliyun_viapi_facebody_compareface_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 请求ID
    TaobaoRequestId   string `json:"taobao_request_id,omitempty" xml:"taobao_request_id,omitempty"`
    // 系统自动生成
    Data   *AliyunViapiFacebodyComparefaceData `json:"data,omitempty" xml:"data,omitempty"`
}
