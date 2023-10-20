package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymItemExternalGoodsBatchOffsaleAPIResponse 交易猫外部商家批量下架商品接口 API返回值
// alibaba.jym.item.external.goods.batch.offsale
//
// 供外部B端商家接入，提交批量下架商品请求，返回批量下架任务结果
type AlibabaJymItemExternalGoodsBatchOffsaleAPIResponse struct {
	model.CommonResponse
	AlibabaJymItemExternalGoodsBatchOffsaleAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaJymItemExternalGoodsBatchOffsaleAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaJymItemExternalGoodsBatchOffsaleAPIResponseModel).Reset()
}

// AlibabaJymItemExternalGoodsBatchOffsaleAPIResponseModel is 交易猫外部商家批量下架商品接口 成功返回结果
type AlibabaJymItemExternalGoodsBatchOffsaleAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_jym_item_external_goods_batch_offsale_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 状态码
	StateCode string `json:"state_code,omitempty" xml:"state_code,omitempty"`
	// 错误信息
	ExtraErrMsg string `json:"extra_err_msg,omitempty" xml:"extra_err_msg,omitempty"`
	// 商品批量下架接口返回
	Result *GoodsBatchResultDto `json:"result,omitempty" xml:"result,omitempty"`
	// 是否成功
	Succeed bool `json:"succeed,omitempty" xml:"succeed,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaJymItemExternalGoodsBatchOffsaleAPIResponseModel) Reset() {
	m.RequestId = ""
	m.StateCode = ""
	m.ExtraErrMsg = ""
	m.Result = nil
	m.Succeed = false
}

var poolAlibabaJymItemExternalGoodsBatchOffsaleAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaJymItemExternalGoodsBatchOffsaleAPIResponse)
	},
}

// GetAlibabaJymItemExternalGoodsBatchOffsaleAPIResponse 从 sync.Pool 获取 AlibabaJymItemExternalGoodsBatchOffsaleAPIResponse
func GetAlibabaJymItemExternalGoodsBatchOffsaleAPIResponse() *AlibabaJymItemExternalGoodsBatchOffsaleAPIResponse {
	return poolAlibabaJymItemExternalGoodsBatchOffsaleAPIResponse.Get().(*AlibabaJymItemExternalGoodsBatchOffsaleAPIResponse)
}

// ReleaseAlibabaJymItemExternalGoodsBatchOffsaleAPIResponse 将 AlibabaJymItemExternalGoodsBatchOffsaleAPIResponse 保存到 sync.Pool
func ReleaseAlibabaJymItemExternalGoodsBatchOffsaleAPIResponse(v *AlibabaJymItemExternalGoodsBatchOffsaleAPIResponse) {
	v.Reset()
	poolAlibabaJymItemExternalGoodsBatchOffsaleAPIResponse.Put(v)
}
