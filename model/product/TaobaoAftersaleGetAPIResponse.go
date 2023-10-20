package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoaftersalegetAPIResponse 查询用户售后服务模板 API返回值
// taobao.aftersale.get
//
// 查询用户设置的售后服务模板，仅返回标题和id
type TaobaoaftersalegetAPIResponse struct {
	model.CommonResponse
	TaobaoaftersalegetAPIResponseModel
}

// TaobaoaftersalegetAPIResponseModel is 查询用户售后服务模板 成功返回结果
type TaobaoaftersalegetAPIResponseModel struct {
	XMLName xml.Name `xml:"aftersale_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 售后服务返回对象
	AfterSales []AfterSale `json:"after_sales,omitempty" xml:"after_sales>after_sale,omitempty"`
}
