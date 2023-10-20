package tmallcar

import (
	"encoding/xml"

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

// TmallAliautoServiceItemGetAPIResponseModel is 查询服务商门店已上架服务商品列表 成功返回结果
type TmallAliautoServiceItemGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_aliauto_service_item_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务出参
	Data *StoreServiceItem4isvDto `json:"data,omitempty" xml:"data,omitempty"`
}
