package viapi

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
绿网-内容安全 API返回值 
aliyun.viapi.imageaudit.scanimage

绿网-内容安全技术是基于阿里云视觉分析技术和深度识别技术，并经过在阿里经济体内和云上客户的多领域、多场景的广泛应用和不断优化，可提供风险和治理领域的图像识别、定位、检索等全面服务能力，不仅可以降低色情、涉恐、涉政、广告、垃圾信息等违规风险，而且能大幅度降低人工审核成本。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
*/
type AliyunViapiImageauditScanimageAPIResponse struct {
    model.CommonResponse
    AliyunViapiImageauditScanimageResponse
}

// 绿网-内容安全 成功返回结果
type AliyunViapiImageauditScanimageResponse struct {
    XMLName xml.Name `xml:"aliyun_viapi_imageaudit_scanimage_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 请求ID
    TaobaoRequestId   string `json:"taobao_request_id,omitempty" xml:"taobao_request_id,omitempty"`
    // 系统自动生成
    DataList   *AliyunViapiImageauditScanimageData `json:"data_list,omitempty" xml:"data_list,omitempty"`
}
