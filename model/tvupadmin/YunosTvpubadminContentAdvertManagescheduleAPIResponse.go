package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentAdvertManagescheduleAPIResponse 广告牌照管控修改 API返回值
// yunos.tvpubadmin.content.advert.manageschedule
//
// 广告牌照管控修改
type YunosTvpubadminContentAdvertManagescheduleAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminContentAdvertManagescheduleAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminContentAdvertManagescheduleAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminContentAdvertManagescheduleAPIResponseModel).Reset()
}

// YunosTvpubadminContentAdvertManagescheduleAPIResponseModel is 广告牌照管控修改 成功返回结果
type YunosTvpubadminContentAdvertManagescheduleAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_content_advert_manageschedule_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 管理广告排期
	Object int64 `json:"object,omitempty" xml:"object,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminContentAdvertManagescheduleAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Object = 0
}

var poolYunosTvpubadminContentAdvertManagescheduleAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminContentAdvertManagescheduleAPIResponse)
	},
}

// GetYunosTvpubadminContentAdvertManagescheduleAPIResponse 从 sync.Pool 获取 YunosTvpubadminContentAdvertManagescheduleAPIResponse
func GetYunosTvpubadminContentAdvertManagescheduleAPIResponse() *YunosTvpubadminContentAdvertManagescheduleAPIResponse {
	return poolYunosTvpubadminContentAdvertManagescheduleAPIResponse.Get().(*YunosTvpubadminContentAdvertManagescheduleAPIResponse)
}

// ReleaseYunosTvpubadminContentAdvertManagescheduleAPIResponse 将 YunosTvpubadminContentAdvertManagescheduleAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminContentAdvertManagescheduleAPIResponse(v *YunosTvpubadminContentAdvertManagescheduleAPIResponse) {
	v.Reset()
	poolYunosTvpubadminContentAdvertManagescheduleAPIResponse.Put(v)
}
