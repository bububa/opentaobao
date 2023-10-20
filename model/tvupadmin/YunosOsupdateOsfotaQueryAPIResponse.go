package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosOsupdateOsfotaQueryAPIResponse 系统升级分页查询 API返回值
// yunos.osupdate.osfota.query
//
// 分页查询osoupdate系统升级列表
type YunosOsupdateOsfotaQueryAPIResponse struct {
	model.CommonResponse
	YunosOsupdateOsfotaQueryAPIResponseModel
}

// Reset 清空结构体
func (m *YunosOsupdateOsfotaQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosOsupdateOsfotaQueryAPIResponseModel).Reset()
}

// YunosOsupdateOsfotaQueryAPIResponseModel is 系统升级分页查询 成功返回结果
type YunosOsupdateOsfotaQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_osupdate_osfota_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *YunosOsupdateOsfotaQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *YunosOsupdateOsfotaQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolYunosOsupdateOsfotaQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosOsupdateOsfotaQueryAPIResponse)
	},
}

// GetYunosOsupdateOsfotaQueryAPIResponse 从 sync.Pool 获取 YunosOsupdateOsfotaQueryAPIResponse
func GetYunosOsupdateOsfotaQueryAPIResponse() *YunosOsupdateOsfotaQueryAPIResponse {
	return poolYunosOsupdateOsfotaQueryAPIResponse.Get().(*YunosOsupdateOsfotaQueryAPIResponse)
}

// ReleaseYunosOsupdateOsfotaQueryAPIResponse 将 YunosOsupdateOsfotaQueryAPIResponse 保存到 sync.Pool
func ReleaseYunosOsupdateOsfotaQueryAPIResponse(v *YunosOsupdateOsfotaQueryAPIResponse) {
	v.Reset()
	poolYunosOsupdateOsfotaQueryAPIResponse.Put(v)
}
