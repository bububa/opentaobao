package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosOsupdateAppversionQueryAPIResponse 分页获取桌面升级任务 API返回值
// yunos.osupdate.appversion.query
//
// 分页获取桌面升级任务
type YunosOsupdateAppversionQueryAPIResponse struct {
	model.CommonResponse
	YunosOsupdateAppversionQueryAPIResponseModel
}

// Reset 清空结构体
func (m *YunosOsupdateAppversionQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosOsupdateAppversionQueryAPIResponseModel).Reset()
}

// YunosOsupdateAppversionQueryAPIResponseModel is 分页获取桌面升级任务 成功返回结果
type YunosOsupdateAppversionQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_osupdate_appversion_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// resultList
	ResultList []TvAppVersion `json:"result_list,omitempty" xml:"result_list>tv_app_version,omitempty"`
	// 总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

// Reset 清空结构体
func (m *YunosOsupdateAppversionQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultList = m.ResultList[:0]
	m.TotalCount = 0
}

var poolYunosOsupdateAppversionQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosOsupdateAppversionQueryAPIResponse)
	},
}

// GetYunosOsupdateAppversionQueryAPIResponse 从 sync.Pool 获取 YunosOsupdateAppversionQueryAPIResponse
func GetYunosOsupdateAppversionQueryAPIResponse() *YunosOsupdateAppversionQueryAPIResponse {
	return poolYunosOsupdateAppversionQueryAPIResponse.Get().(*YunosOsupdateAppversionQueryAPIResponse)
}

// ReleaseYunosOsupdateAppversionQueryAPIResponse 将 YunosOsupdateAppversionQueryAPIResponse 保存到 sync.Pool
func ReleaseYunosOsupdateAppversionQueryAPIResponse(v *YunosOsupdateAppversionQueryAPIResponse) {
	v.Reset()
	poolYunosOsupdateAppversionQueryAPIResponse.Put(v)
}
