package tmallcar

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallAliautoServiceItemGetAPIResponse 查询服务商门店已上架服务商品列表 API返回值
// tmall.aliauto.service.item.get
//
// 根据门店自定义门店编码查询门店【已上架】服务商品列表
type TmallAliautoServiceItemGetAPIResponse struct {
	model.CommonResponse
	TmallAliautoServiceItemGetAPIResponseModel
}

// Reset 清空结构体
func (m *TmallAliautoServiceItemGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallAliautoServiceItemGetAPIResponseModel).Reset()
}

// TmallAliautoServiceItemGetAPIResponseModel is 查询服务商门店已上架服务商品列表 成功返回结果
type TmallAliautoServiceItemGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_aliauto_service_item_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务出参
	Data *StoreServiceItem4IsvDto `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TmallAliautoServiceItemGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolTmallAliautoServiceItemGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallAliautoServiceItemGetAPIResponse)
	},
}

// GetTmallAliautoServiceItemGetAPIResponse 从 sync.Pool 获取 TmallAliautoServiceItemGetAPIResponse
func GetTmallAliautoServiceItemGetAPIResponse() *TmallAliautoServiceItemGetAPIResponse {
	return poolTmallAliautoServiceItemGetAPIResponse.Get().(*TmallAliautoServiceItemGetAPIResponse)
}

// ReleaseTmallAliautoServiceItemGetAPIResponse 将 TmallAliautoServiceItemGetAPIResponse 保存到 sync.Pool
func ReleaseTmallAliautoServiceItemGetAPIResponse(v *TmallAliautoServiceItemGetAPIResponse) {
	v.Reset()
	poolTmallAliautoServiceItemGetAPIResponse.Put(v)
}
