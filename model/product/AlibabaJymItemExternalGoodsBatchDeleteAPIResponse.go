package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymItemExternalGoodsBatchDeleteAPIResponse 交易猫外部商家批量删除商品接口 API返回值
// alibaba.jym.item.external.goods.batch.delete
//
// 交易猫外部商家批量删除商品接口
type AlibabaJymItemExternalGoodsBatchDeleteAPIResponse struct {
	model.CommonResponse
	AlibabaJymItemExternalGoodsBatchDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaJymItemExternalGoodsBatchDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaJymItemExternalGoodsBatchDeleteAPIResponseModel).Reset()
}

// AlibabaJymItemExternalGoodsBatchDeleteAPIResponseModel is 交易猫外部商家批量删除商品接口 成功返回结果
type AlibabaJymItemExternalGoodsBatchDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_jym_item_external_goods_batch_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 状态码
	StateCode string `json:"state_code,omitempty" xml:"state_code,omitempty"`
	// 错误信息
	ExtraErrMsg string `json:"extra_err_msg,omitempty" xml:"extra_err_msg,omitempty"`
	// 商品批量删除接口返回
	Result *GoodsBatchResultDto `json:"result,omitempty" xml:"result,omitempty"`
	// 是否成功
	Succeed bool `json:"succeed,omitempty" xml:"succeed,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaJymItemExternalGoodsBatchDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.StateCode = ""
	m.ExtraErrMsg = ""
	m.Result = nil
	m.Succeed = false
}

var poolAlibabaJymItemExternalGoodsBatchDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaJymItemExternalGoodsBatchDeleteAPIResponse)
	},
}

// GetAlibabaJymItemExternalGoodsBatchDeleteAPIResponse 从 sync.Pool 获取 AlibabaJymItemExternalGoodsBatchDeleteAPIResponse
func GetAlibabaJymItemExternalGoodsBatchDeleteAPIResponse() *AlibabaJymItemExternalGoodsBatchDeleteAPIResponse {
	return poolAlibabaJymItemExternalGoodsBatchDeleteAPIResponse.Get().(*AlibabaJymItemExternalGoodsBatchDeleteAPIResponse)
}

// ReleaseAlibabaJymItemExternalGoodsBatchDeleteAPIResponse 将 AlibabaJymItemExternalGoodsBatchDeleteAPIResponse 保存到 sync.Pool
func ReleaseAlibabaJymItemExternalGoodsBatchDeleteAPIResponse(v *AlibabaJymItemExternalGoodsBatchDeleteAPIResponse) {
	v.Reset()
	poolAlibabaJymItemExternalGoodsBatchDeleteAPIResponse.Put(v)
}
