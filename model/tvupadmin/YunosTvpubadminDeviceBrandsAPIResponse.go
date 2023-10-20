package tvupadmin

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *YunosTvpubadminDeviceBrandsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminDeviceBrandsAPIResponseModel).Reset()
}

// YunosTvpubadminDeviceBrandsAPIResponseModel is 获取终端类型下品牌列表 成功返回结果
type YunosTvpubadminDeviceBrandsAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_device_brands_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// object
	List []BrandDo `json:"list,omitempty" xml:"list>brand_do,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminDeviceBrandsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.List = m.List[:0]
}

var poolYunosTvpubadminDeviceBrandsAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminDeviceBrandsAPIResponse)
	},
}

// GetYunosTvpubadminDeviceBrandsAPIResponse 从 sync.Pool 获取 YunosTvpubadminDeviceBrandsAPIResponse
func GetYunosTvpubadminDeviceBrandsAPIResponse() *YunosTvpubadminDeviceBrandsAPIResponse {
	return poolYunosTvpubadminDeviceBrandsAPIResponse.Get().(*YunosTvpubadminDeviceBrandsAPIResponse)
}

// ReleaseYunosTvpubadminDeviceBrandsAPIResponse 将 YunosTvpubadminDeviceBrandsAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminDeviceBrandsAPIResponse(v *YunosTvpubadminDeviceBrandsAPIResponse) {
	v.Reset()
	poolYunosTvpubadminDeviceBrandsAPIResponse.Put(v)
}
