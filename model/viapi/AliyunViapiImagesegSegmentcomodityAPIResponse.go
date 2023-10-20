package viapi

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliyunviapiimagesegsegmentcomodityAPIResponse 商品分割 API返回值
// aliyun.viapi.imageseg.segmentcomodity
//
// 识别输入图像中的商品轮廓，与背景进行分离，返回分割后的前景商品图（4通道png透明图），适应单商品/多商品、复杂背景。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
type AliyunviapiimagesegsegmentcomodityAPIResponse struct {
	model.CommonResponse
	AliyunviapiimagesegsegmentcomodityAPIResponseModel
}

// AliyunviapiimagesegsegmentcomodityAPIResponseModel is 商品分割 成功返回结果
type AliyunviapiimagesegsegmentcomodityAPIResponseModel struct {
	XMLName xml.Name `xml:"aliyun_viapi_imageseg_segmentcomodity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求ID
	TaobaoRequestId string `json:"taobao_request_id,omitempty" xml:"taobao_request_id,omitempty"`
	// 系统自动生成
	Data *AliyunviapiimagesegsegmentcomodityData `json:"data,omitempty" xml:"data,omitempty"`
}
