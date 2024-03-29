package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentShowGetshowexemptauditmapAPIResponse 迎客松批量查询节目某个牌照的免审状态 API返回值
// yunos.tvpubadmin.content.show.getshowexemptauditmap
//
// 迎客松批量查询节目某个牌照的免审状态
type YunosTvpubadminContentShowGetshowexemptauditmapAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminContentShowGetshowexemptauditmapAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminContentShowGetshowexemptauditmapAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminContentShowGetshowexemptauditmapAPIResponseModel).Reset()
}

// YunosTvpubadminContentShowGetshowexemptauditmapAPIResponseModel is 迎客松批量查询节目某个牌照的免审状态 成功返回结果
type YunosTvpubadminContentShowGetshowexemptauditmapAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_content_show_getshowexemptauditmap_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// object
	Object string `json:"object,omitempty" xml:"object,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminContentShowGetshowexemptauditmapAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Object = ""
}

var poolYunosTvpubadminContentShowGetshowexemptauditmapAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminContentShowGetshowexemptauditmapAPIResponse)
	},
}

// GetYunosTvpubadminContentShowGetshowexemptauditmapAPIResponse 从 sync.Pool 获取 YunosTvpubadminContentShowGetshowexemptauditmapAPIResponse
func GetYunosTvpubadminContentShowGetshowexemptauditmapAPIResponse() *YunosTvpubadminContentShowGetshowexemptauditmapAPIResponse {
	return poolYunosTvpubadminContentShowGetshowexemptauditmapAPIResponse.Get().(*YunosTvpubadminContentShowGetshowexemptauditmapAPIResponse)
}

// ReleaseYunosTvpubadminContentShowGetshowexemptauditmapAPIResponse 将 YunosTvpubadminContentShowGetshowexemptauditmapAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminContentShowGetshowexemptauditmapAPIResponse(v *YunosTvpubadminContentShowGetshowexemptauditmapAPIResponse) {
	v.Reset()
	poolYunosTvpubadminContentShowGetshowexemptauditmapAPIResponse.Put(v)
}
