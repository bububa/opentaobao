package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymItemExternalGoodsBatchModifypriceAPIResponse 交易猫外部商家批量商品改价接口 API返回值
// alibaba.jym.item.external.goods.batch.modifyprice
//
// 供外部B端商家接入，提交批量商品改价请求，返回批量改价任务结果
type AlibabaJymItemExternalGoodsBatchModifypriceAPIResponse struct {
	model.CommonResponse
	AlibabaJymItemExternalGoodsBatchModifypriceAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaJymItemExternalGoodsBatchModifypriceAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaJymItemExternalGoodsBatchModifypriceAPIResponseModel).Reset()
}

// AlibabaJymItemExternalGoodsBatchModifypriceAPIResponseModel is 交易猫外部商家批量商品改价接口 成功返回结果
type AlibabaJymItemExternalGoodsBatchModifypriceAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_jym_item_external_goods_batch_modifyprice_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 状态码
	StateCode string `json:"state_code,omitempty" xml:"state_code,omitempty"`
	// 错误信息
	ExtraErrMsg string `json:"extra_err_msg,omitempty" xml:"extra_err_msg,omitempty"`
	// 商品批量改价接口返回
	Result *GoodsBatchResultDto `json:"result,omitempty" xml:"result,omitempty"`
	// 是否成功
	Succeed bool `json:"succeed,omitempty" xml:"succeed,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaJymItemExternalGoodsBatchModifypriceAPIResponseModel) Reset() {
	m.RequestId = ""
	m.StateCode = ""
	m.ExtraErrMsg = ""
	m.Result = nil
	m.Succeed = false
}

var poolAlibabaJymItemExternalGoodsBatchModifypriceAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaJymItemExternalGoodsBatchModifypriceAPIResponse)
	},
}

// GetAlibabaJymItemExternalGoodsBatchModifypriceAPIResponse 从 sync.Pool 获取 AlibabaJymItemExternalGoodsBatchModifypriceAPIResponse
func GetAlibabaJymItemExternalGoodsBatchModifypriceAPIResponse() *AlibabaJymItemExternalGoodsBatchModifypriceAPIResponse {
	return poolAlibabaJymItemExternalGoodsBatchModifypriceAPIResponse.Get().(*AlibabaJymItemExternalGoodsBatchModifypriceAPIResponse)
}

// ReleaseAlibabaJymItemExternalGoodsBatchModifypriceAPIResponse 将 AlibabaJymItemExternalGoodsBatchModifypriceAPIResponse 保存到 sync.Pool
func ReleaseAlibabaJymItemExternalGoodsBatchModifypriceAPIResponse(v *AlibabaJymItemExternalGoodsBatchModifypriceAPIResponse) {
	v.Reset()
	poolAlibabaJymItemExternalGoodsBatchModifypriceAPIResponse.Put(v)
}
