package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosOsupdateModelSearchAPIResponse 机型检索 API返回值
// yunos.osupdate.model.search
//
// 机型检索
type YunosOsupdateModelSearchAPIResponse struct {
	model.CommonResponse
	YunosOsupdateModelSearchAPIResponseModel
}

// Reset 清空结构体
func (m *YunosOsupdateModelSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosOsupdateModelSearchAPIResponseModel).Reset()
}

// YunosOsupdateModelSearchAPIResponseModel is 机型检索 成功返回结果
type YunosOsupdateModelSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_osupdate_model_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 机型列表
	ModelList []DeviceEntryDto `json:"model_list,omitempty" xml:"model_list>device_entry_dto,omitempty"`
}

// Reset 清空结构体
func (m *YunosOsupdateModelSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ModelList = m.ModelList[:0]
}

var poolYunosOsupdateModelSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosOsupdateModelSearchAPIResponse)
	},
}

// GetYunosOsupdateModelSearchAPIResponse 从 sync.Pool 获取 YunosOsupdateModelSearchAPIResponse
func GetYunosOsupdateModelSearchAPIResponse() *YunosOsupdateModelSearchAPIResponse {
	return poolYunosOsupdateModelSearchAPIResponse.Get().(*YunosOsupdateModelSearchAPIResponse)
}

// ReleaseYunosOsupdateModelSearchAPIResponse 将 YunosOsupdateModelSearchAPIResponse 保存到 sync.Pool
func ReleaseYunosOsupdateModelSearchAPIResponse(v *YunosOsupdateModelSearchAPIResponse) {
	v.Reset()
	poolYunosOsupdateModelSearchAPIResponse.Put(v)
}
