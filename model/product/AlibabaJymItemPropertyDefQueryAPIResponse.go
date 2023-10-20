package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymItemPropertyDefQueryAPIResponse 交易猫商品属性定义查询 API返回值
// alibaba.jym.item.property.def.query
//
// 查询商品发布属性定义详情
type AlibabaJymItemPropertyDefQueryAPIResponse struct {
	model.CommonResponse
	AlibabaJymItemPropertyDefQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaJymItemPropertyDefQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaJymItemPropertyDefQueryAPIResponseModel).Reset()
}

// AlibabaJymItemPropertyDefQueryAPIResponseModel is 交易猫商品属性定义查询 成功返回结果
type AlibabaJymItemPropertyDefQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_jym_item_property_def_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 状态码
	StateCode string `json:"state_code,omitempty" xml:"state_code,omitempty"`
	// 错误信息
	ExtraErrMsg string `json:"extra_err_msg,omitempty" xml:"extra_err_msg,omitempty"`
	// 商品发布填写资料定义DTO
	Result *GoodsPublishPropertyDefDetailDto `json:"result,omitempty" xml:"result,omitempty"`
	// 是否成功
	Succeed bool `json:"succeed,omitempty" xml:"succeed,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaJymItemPropertyDefQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.StateCode = ""
	m.ExtraErrMsg = ""
	m.Result = nil
	m.Succeed = false
}

var poolAlibabaJymItemPropertyDefQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaJymItemPropertyDefQueryAPIResponse)
	},
}

// GetAlibabaJymItemPropertyDefQueryAPIResponse 从 sync.Pool 获取 AlibabaJymItemPropertyDefQueryAPIResponse
func GetAlibabaJymItemPropertyDefQueryAPIResponse() *AlibabaJymItemPropertyDefQueryAPIResponse {
	return poolAlibabaJymItemPropertyDefQueryAPIResponse.Get().(*AlibabaJymItemPropertyDefQueryAPIResponse)
}

// ReleaseAlibabaJymItemPropertyDefQueryAPIResponse 将 AlibabaJymItemPropertyDefQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaJymItemPropertyDefQueryAPIResponse(v *AlibabaJymItemPropertyDefQueryAPIResponse) {
	v.Reset()
	poolAlibabaJymItemPropertyDefQueryAPIResponse.Put(v)
}
