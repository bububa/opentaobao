package shop

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoshopupdateAPIResponse 更新店铺基本信息 API返回值
// taobao.shop.update
//
// 目前只支持标题、公告和描述的更新
type TaobaoshopupdateAPIResponse struct {
	model.CommonResponse
	TaobaoshopupdateAPIResponseModel
}

// TaobaoshopupdateAPIResponseModel is 更新店铺基本信息 成功返回结果
type TaobaoshopupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"shop_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 店铺信息
	Shop *Shop `json:"shop,omitempty" xml:"shop,omitempty"`
}
