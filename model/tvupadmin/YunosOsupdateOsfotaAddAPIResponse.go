package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosOsupdateOsfotaAddAPIResponse 添加系统升级任务 API返回值
// yunos.osupdate.osfota.add
//
// 添加osupdate系统升级任务
type YunosOsupdateOsfotaAddAPIResponse struct {
	model.CommonResponse
	YunosOsupdateOsfotaAddAPIResponseModel
}

// Reset 清空结构体
func (m *YunosOsupdateOsfotaAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosOsupdateOsfotaAddAPIResponseModel).Reset()
}

// YunosOsupdateOsfotaAddAPIResponseModel is 添加系统升级任务 成功返回结果
type YunosOsupdateOsfotaAddAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_osupdate_osfota_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *YunosOsupdateOsfotaAddResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *YunosOsupdateOsfotaAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolYunosOsupdateOsfotaAddAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosOsupdateOsfotaAddAPIResponse)
	},
}

// GetYunosOsupdateOsfotaAddAPIResponse 从 sync.Pool 获取 YunosOsupdateOsfotaAddAPIResponse
func GetYunosOsupdateOsfotaAddAPIResponse() *YunosOsupdateOsfotaAddAPIResponse {
	return poolYunosOsupdateOsfotaAddAPIResponse.Get().(*YunosOsupdateOsfotaAddAPIResponse)
}

// ReleaseYunosOsupdateOsfotaAddAPIResponse 将 YunosOsupdateOsfotaAddAPIResponse 保存到 sync.Pool
func ReleaseYunosOsupdateOsfotaAddAPIResponse(v *YunosOsupdateOsfotaAddAPIResponse) {
	v.Reset()
	poolYunosOsupdateOsfotaAddAPIResponse.Put(v)
}
