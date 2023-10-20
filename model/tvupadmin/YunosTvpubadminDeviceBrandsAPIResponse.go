package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmindevicebrandsAPIResponse 获取终端类型下品牌列表 API返回值
// yunos.tvpubadmin.device.brands
//
// 获取终端类型下品牌列表
type YunostvpubadmindevicebrandsAPIResponse struct {
	model.CommonResponse
	YunostvpubadmindevicebrandsAPIResponseModel
}

// YunostvpubadmindevicebrandsAPIResponseModel is 获取终端类型下品牌列表 成功返回结果
type YunostvpubadmindevicebrandsAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_device_brands_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// object
	List []BrandDo `json:"list,omitempty" xml:"list>brand_do,omitempty"`
}
