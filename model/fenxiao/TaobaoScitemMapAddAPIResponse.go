package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoScitemMapAddAPIResponse 创建IC商品与后端商品的映射关系 API返回值
// taobao.scitem.map.add
//
// 创建IC商品或分销商品与后端商品的映射关系
type TaobaoScitemMapAddAPIResponse struct {
	model.CommonResponse
	TaobaoScitemMapAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoScitemMapAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoScitemMapAddAPIResponseModel).Reset()
}

// TaobaoScitemMapAddAPIResponseModel is 创建IC商品与后端商品的映射关系 成功返回结果
type TaobaoScitemMapAddAPIResponseModel struct {
	XMLName xml.Name `xml:"scitem_map_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口调用返回结果信息：商家编码
	OuterCode string `json:"outer_code,omitempty" xml:"outer_code,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoScitemMapAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.OuterCode = ""
}

var poolTaobaoScitemMapAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoScitemMapAddAPIResponse)
	},
}

// GetTaobaoScitemMapAddAPIResponse 从 sync.Pool 获取 TaobaoScitemMapAddAPIResponse
func GetTaobaoScitemMapAddAPIResponse() *TaobaoScitemMapAddAPIResponse {
	return poolTaobaoScitemMapAddAPIResponse.Get().(*TaobaoScitemMapAddAPIResponse)
}

// ReleaseTaobaoScitemMapAddAPIResponse 将 TaobaoScitemMapAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoScitemMapAddAPIResponse(v *TaobaoScitemMapAddAPIResponse) {
	v.Reset()
	poolTaobaoScitemMapAddAPIResponse.Put(v)
}
