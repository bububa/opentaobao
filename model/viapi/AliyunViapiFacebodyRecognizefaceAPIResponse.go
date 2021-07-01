package viapi

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
人脸属性识别 API返回值 
aliyun.viapi.facebody.recognizeface

输入图片之后，在人脸检测定位返回结果的基础上，识别各个检测人脸的四种属性，返回性别（男/女）、年龄、表情（笑/不笑）、眼镜（戴/不戴）；并可返回人脸的1024维深度学习特征。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
*/
type AliyunViapiFacebodyRecognizefaceAPIResponse struct {
    model.CommonResponse
    AliyunViapiFacebodyRecognizefaceAPIResponseModel
}

// 人脸属性识别 成功返回结果
type AliyunViapiFacebodyRecognizefaceAPIResponseModel struct {
    XMLName xml.Name `xml:"aliyun_viapi_facebody_recognizeface_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 请求ID
    TaobaoRequestId   string `json:"taobao_request_id,omitempty" xml:"taobao_request_id,omitempty"`
    // 系统自动生成
    Data   *AliyunViapiFacebodyRecognizefaceData `json:"data,omitempty" xml:"data,omitempty"`
}
