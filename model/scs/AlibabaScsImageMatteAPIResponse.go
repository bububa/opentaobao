package scs

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScsImageMatteAPIResponse 阿里妈妈智能创意平台在线抠图 API返回值
// alibaba.scs.image.matte
//
// 该API对外输出一个在线抠图(Deep Image Matting)接口，合作方可以通过该接口利用深度学习抠图算法从图片中抠出目标对象(比如商品或者人物轮廓)
type AlibabaScsImageMatteAPIResponse struct {
	model.CommonResponse
	AlibabaScsImageMatteAPIResponseModel
}

// AlibabaScsImageMatteAPIResponseModel is 阿里妈妈智能创意平台在线抠图 成功返回结果
type AlibabaScsImageMatteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scs_image_matte_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	Errno string `json:"errno,omitempty" xml:"errno,omitempty"`
	// 分组数据
	DataList []Array `json:"data_list,omitempty" xml:"data_list>array,omitempty"`
	// 错误提示信息
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	// 会话ID
	SessionId string `json:"session_id,omitempty" xml:"session_id,omitempty"`
	// 桶号ID (调用方暂不用关心)
	Bucketid string `json:"bucketid,omitempty" xml:"bucketid,omitempty"`
}
