package caipiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCaipiaoShopInfoInputAPIResponse 录入参加送彩票店铺信息 API返回值
// taobao.caipiao.shop.info.input
//
// 录入参加送彩票店铺信息，如果录入成功，返回true，如果录入失败，返回false，后端会根据卖家id与赠送类型（sellerid_presenttype_uk）来决定是新增数据还是修改数据。
type TaobaoCaipiaoShopInfoInputAPIResponse struct {
	model.CommonResponse
	TaobaoCaipiaoShopInfoInputAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoCaipiaoShopInfoInputAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoCaipiaoShopInfoInputAPIResponseModel).Reset()
}

// TaobaoCaipiaoShopInfoInputAPIResponseModel is 录入参加送彩票店铺信息 成功返回结果
type TaobaoCaipiaoShopInfoInputAPIResponseModel struct {
	XMLName xml.Name `xml:"caipiao_shop_info_input_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 录入操作是否成功
	InputResult bool `json:"input_result,omitempty" xml:"input_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoCaipiaoShopInfoInputAPIResponseModel) Reset() {
	m.RequestId = ""
	m.InputResult = false
}

var poolTaobaoCaipiaoShopInfoInputAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoCaipiaoShopInfoInputAPIResponse)
	},
}

// GetTaobaoCaipiaoShopInfoInputAPIResponse 从 sync.Pool 获取 TaobaoCaipiaoShopInfoInputAPIResponse
func GetTaobaoCaipiaoShopInfoInputAPIResponse() *TaobaoCaipiaoShopInfoInputAPIResponse {
	return poolTaobaoCaipiaoShopInfoInputAPIResponse.Get().(*TaobaoCaipiaoShopInfoInputAPIResponse)
}

// ReleaseTaobaoCaipiaoShopInfoInputAPIResponse 将 TaobaoCaipiaoShopInfoInputAPIResponse 保存到 sync.Pool
func ReleaseTaobaoCaipiaoShopInfoInputAPIResponse(v *TaobaoCaipiaoShopInfoInputAPIResponse) {
	v.Reset()
	poolTaobaoCaipiaoShopInfoInputAPIResponse.Put(v)
}
