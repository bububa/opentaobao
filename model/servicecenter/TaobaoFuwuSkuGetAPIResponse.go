package servicecenter

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaofuwuskugetAPIResponse 获取内购服务及SKU详情 API返回值
// taobao.fuwu.sku.get
//
// 通过服务code和用户nick，获取该服务对应的收费项目的sku信息，包括价格、可购买周期、用户能否购买等信息
type TaobaofuwuskugetAPIResponse struct {
	model.CommonResponse
	TaobaofuwuskugetAPIResponseModel
}

// TaobaofuwuskugetAPIResponseModel is 获取内购服务及SKU详情 成功返回结果
type TaobaofuwuskugetAPIResponseModel struct {
	XMLName xml.Name `xml:"fuwu_sku_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 内购服务及SKU详情
	Result *ArticleViewResult `json:"result,omitempty" xml:"result,omitempty"`
}
