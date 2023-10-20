package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabatianmaolanpeiuopcreateAPIResponse 阿里巴巴.天猫家装.揽配.履约订单.创建 API返回值
// alibaba.tianmao.lanpei.uop.create
//
// 阿里巴巴.天猫家装.揽配.履约订单.创建
type AlibabatianmaolanpeiuopcreateAPIResponse struct {
	model.CommonResponse
	AlibabatianmaolanpeiuopcreateAPIResponseModel
}

// AlibabatianmaolanpeiuopcreateAPIResponseModel is 阿里巴巴.天猫家装.揽配.履约订单.创建 成功返回结果
type AlibabatianmaolanpeiuopcreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tianmao_lanpei_uop_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
