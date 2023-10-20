package bus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripbusinsurancerecommendAPIResponse 汽车票保险推荐接口 API返回值
// alitrip.bus.insurance.recommend
//
// 汽车票保险推荐接口-供商家售卖飞猪保险使用
type AlitripbusinsurancerecommendAPIResponse struct {
	model.CommonResponse
	AlitripbusinsurancerecommendAPIResponseModel
}

// AlitripbusinsurancerecommendAPIResponseModel is 汽车票保险推荐接口 成功返回结果
type AlitripbusinsurancerecommendAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_bus_insurance_recommend_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 保险产品列表
	InsProductList []InsuranceProductVo `json:"ins_product_list,omitempty" xml:"ins_product_list>insurance_product_vo,omitempty"`
	// 接口业务成功时为空
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 错误信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// true:成功；false:失败
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
