package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosOsupdateOsfotaPublishAPIResponse 系统升级发布 API返回值
// yunos.osupdate.osfota.publish
//
// 发布osupdate系统升级任务
type YunosOsupdateOsfotaPublishAPIResponse struct {
	model.CommonResponse
	YunosOsupdateOsfotaPublishAPIResponseModel
}

// Reset 清空结构体
func (m *YunosOsupdateOsfotaPublishAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosOsupdateOsfotaPublishAPIResponseModel).Reset()
}

// YunosOsupdateOsfotaPublishAPIResponseModel is 系统升级发布 成功返回结果
type YunosOsupdateOsfotaPublishAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_osupdate_osfota_publish_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *YunosOsupdateOsfotaPublishResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *YunosOsupdateOsfotaPublishAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolYunosOsupdateOsfotaPublishAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosOsupdateOsfotaPublishAPIResponse)
	},
}

// GetYunosOsupdateOsfotaPublishAPIResponse 从 sync.Pool 获取 YunosOsupdateOsfotaPublishAPIResponse
func GetYunosOsupdateOsfotaPublishAPIResponse() *YunosOsupdateOsfotaPublishAPIResponse {
	return poolYunosOsupdateOsfotaPublishAPIResponse.Get().(*YunosOsupdateOsfotaPublishAPIResponse)
}

// ReleaseYunosOsupdateOsfotaPublishAPIResponse 将 YunosOsupdateOsfotaPublishAPIResponse 保存到 sync.Pool
func ReleaseYunosOsupdateOsfotaPublishAPIResponse(v *YunosOsupdateOsfotaPublishAPIResponse) {
	v.Reset()
	poolYunosOsupdateOsfotaPublishAPIResponse.Put(v)
}
