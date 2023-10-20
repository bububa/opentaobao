package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosOsupdateAppversionPublishAPIResponse 发布应用升级 API返回值
// yunos.osupdate.appversion.publish
//
// 发布应用升级任务
type YunosOsupdateAppversionPublishAPIResponse struct {
	model.CommonResponse
	YunosOsupdateAppversionPublishAPIResponseModel
}

// Reset 清空结构体
func (m *YunosOsupdateAppversionPublishAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosOsupdateAppversionPublishAPIResponseModel).Reset()
}

// YunosOsupdateAppversionPublishAPIResponseModel is 发布应用升级 成功返回结果
type YunosOsupdateAppversionPublishAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_osupdate_appversion_publish_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *YunosOsupdateAppversionPublishResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *YunosOsupdateAppversionPublishAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolYunosOsupdateAppversionPublishAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosOsupdateAppversionPublishAPIResponse)
	},
}

// GetYunosOsupdateAppversionPublishAPIResponse 从 sync.Pool 获取 YunosOsupdateAppversionPublishAPIResponse
func GetYunosOsupdateAppversionPublishAPIResponse() *YunosOsupdateAppversionPublishAPIResponse {
	return poolYunosOsupdateAppversionPublishAPIResponse.Get().(*YunosOsupdateAppversionPublishAPIResponse)
}

// ReleaseYunosOsupdateAppversionPublishAPIResponse 将 YunosOsupdateAppversionPublishAPIResponse 保存到 sync.Pool
func ReleaseYunosOsupdateAppversionPublishAPIResponse(v *YunosOsupdateAppversionPublishAPIResponse) {
	v.Reset()
	poolYunosOsupdateAppversionPublishAPIResponse.Put(v)
}
