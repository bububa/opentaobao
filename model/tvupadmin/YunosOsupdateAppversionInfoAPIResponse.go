package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosOsupdateAppversionInfoAPIResponse 获取应用升级详情 API返回值
// yunos.osupdate.appversion.info
//
// 获取应用升级详情
type YunosOsupdateAppversionInfoAPIResponse struct {
	model.CommonResponse
	YunosOsupdateAppversionInfoAPIResponseModel
}

// Reset 清空结构体
func (m *YunosOsupdateAppversionInfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosOsupdateAppversionInfoAPIResponseModel).Reset()
}

// YunosOsupdateAppversionInfoAPIResponseModel is 获取应用升级详情 成功返回结果
type YunosOsupdateAppversionInfoAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_osupdate_appversion_info_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Data *TvAppVersion `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *YunosOsupdateAppversionInfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolYunosOsupdateAppversionInfoAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosOsupdateAppversionInfoAPIResponse)
	},
}

// GetYunosOsupdateAppversionInfoAPIResponse 从 sync.Pool 获取 YunosOsupdateAppversionInfoAPIResponse
func GetYunosOsupdateAppversionInfoAPIResponse() *YunosOsupdateAppversionInfoAPIResponse {
	return poolYunosOsupdateAppversionInfoAPIResponse.Get().(*YunosOsupdateAppversionInfoAPIResponse)
}

// ReleaseYunosOsupdateAppversionInfoAPIResponse 将 YunosOsupdateAppversionInfoAPIResponse 保存到 sync.Pool
func ReleaseYunosOsupdateAppversionInfoAPIResponse(v *YunosOsupdateAppversionInfoAPIResponse) {
	v.Reset()
	poolYunosOsupdateAppversionInfoAPIResponse.Put(v)
}
