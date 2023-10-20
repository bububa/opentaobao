package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymItemExternalGoodsDetailQueryAPIResponse 交易猫外部商家商品详情查询接口 API返回值
// alibaba.jym.item.external.goods.detail.query
//
// 供外部B端商家接入，请求查询商品详情，返回商品详情查询结果
type AlibabaJymItemExternalGoodsDetailQueryAPIResponse struct {
	model.CommonResponse
	AlibabaJymItemExternalGoodsDetailQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaJymItemExternalGoodsDetailQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaJymItemExternalGoodsDetailQueryAPIResponseModel).Reset()
}

// AlibabaJymItemExternalGoodsDetailQueryAPIResponseModel is 交易猫外部商家商品详情查询接口 成功返回结果
type AlibabaJymItemExternalGoodsDetailQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_jym_item_external_goods_detail_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 状态码
	StateCode string `json:"state_code,omitempty" xml:"state_code,omitempty"`
	// 错误信息
	ExtraErrMsg string `json:"extra_err_msg,omitempty" xml:"extra_err_msg,omitempty"`
	// 商品详情查询接口返回
	Result *GoodsDetailResultDto `json:"result,omitempty" xml:"result,omitempty"`
	// 是否成功
	Succeed bool `json:"succeed,omitempty" xml:"succeed,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaJymItemExternalGoodsDetailQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.StateCode = ""
	m.ExtraErrMsg = ""
	m.Result = nil
	m.Succeed = false
}

var poolAlibabaJymItemExternalGoodsDetailQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaJymItemExternalGoodsDetailQueryAPIResponse)
	},
}

// GetAlibabaJymItemExternalGoodsDetailQueryAPIResponse 从 sync.Pool 获取 AlibabaJymItemExternalGoodsDetailQueryAPIResponse
func GetAlibabaJymItemExternalGoodsDetailQueryAPIResponse() *AlibabaJymItemExternalGoodsDetailQueryAPIResponse {
	return poolAlibabaJymItemExternalGoodsDetailQueryAPIResponse.Get().(*AlibabaJymItemExternalGoodsDetailQueryAPIResponse)
}

// ReleaseAlibabaJymItemExternalGoodsDetailQueryAPIResponse 将 AlibabaJymItemExternalGoodsDetailQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaJymItemExternalGoodsDetailQueryAPIResponse(v *AlibabaJymItemExternalGoodsDetailQueryAPIResponse) {
	v.Reset()
	poolAlibabaJymItemExternalGoodsDetailQueryAPIResponse.Put(v)
}
