package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymItemExternalGoodsBatchtaskQueryAPIResponse 交易猫外部商家查询商品批量任务接口 API返回值
// alibaba.jym.item.external.goods.batchtask.query
//
// 供外部B端商家接入，请求查询商品批量任务，返回商品批量任务查询结果
type AlibabaJymItemExternalGoodsBatchtaskQueryAPIResponse struct {
	model.CommonResponse
	AlibabaJymItemExternalGoodsBatchtaskQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaJymItemExternalGoodsBatchtaskQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaJymItemExternalGoodsBatchtaskQueryAPIResponseModel).Reset()
}

// AlibabaJymItemExternalGoodsBatchtaskQueryAPIResponseModel is 交易猫外部商家查询商品批量任务接口 成功返回结果
type AlibabaJymItemExternalGoodsBatchtaskQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_jym_item_external_goods_batchtask_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 状态码
	StateCode string `json:"state_code,omitempty" xml:"state_code,omitempty"`
	// 错误信息
	ExtraErrMsg string `json:"extra_err_msg,omitempty" xml:"extra_err_msg,omitempty"`
	// 查询商品批量任务接口返回
	Result *GoodsBatchTaskResultDto `json:"result,omitempty" xml:"result,omitempty"`
	// 是否成功
	Succeed bool `json:"succeed,omitempty" xml:"succeed,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaJymItemExternalGoodsBatchtaskQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.StateCode = ""
	m.ExtraErrMsg = ""
	m.Result = nil
	m.Succeed = false
}

var poolAlibabaJymItemExternalGoodsBatchtaskQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaJymItemExternalGoodsBatchtaskQueryAPIResponse)
	},
}

// GetAlibabaJymItemExternalGoodsBatchtaskQueryAPIResponse 从 sync.Pool 获取 AlibabaJymItemExternalGoodsBatchtaskQueryAPIResponse
func GetAlibabaJymItemExternalGoodsBatchtaskQueryAPIResponse() *AlibabaJymItemExternalGoodsBatchtaskQueryAPIResponse {
	return poolAlibabaJymItemExternalGoodsBatchtaskQueryAPIResponse.Get().(*AlibabaJymItemExternalGoodsBatchtaskQueryAPIResponse)
}

// ReleaseAlibabaJymItemExternalGoodsBatchtaskQueryAPIResponse 将 AlibabaJymItemExternalGoodsBatchtaskQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaJymItemExternalGoodsBatchtaskQueryAPIResponse(v *AlibabaJymItemExternalGoodsBatchtaskQueryAPIResponse) {
	v.Reset()
	poolAlibabaJymItemExternalGoodsBatchtaskQueryAPIResponse.Put(v)
}
