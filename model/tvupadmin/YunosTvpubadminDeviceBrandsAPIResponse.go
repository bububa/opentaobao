package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDeviceBrandsAPIResponse 获取终端类型下品牌列表 API返回值
// yunos.tvpubadmin.device.brands
//
// 获取终端类型下品牌列表
type YunosTvpubadminDeviceBrandsAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminDeviceBrandsAPIResponseModel
}

// YunosTvpubadminDeviceBrandsAPIResponseModel is 获取终端类型下品牌列表 成功返回结果
type YunosTvpubadminDeviceBrandsAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_device_brands_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// object
	List []BrandDo `json:"list,omitempty" xml:"list>brand_do,omitempty"`
}
