package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentShowGetbyshowlongidAPIResponse 迎客松根据节目longid获取节目元数据 API返回值
// yunos.tvpubadmin.content.show.getbyshowlongid
//
// 迎客松根据节目longid获取节目元数据
type YunosTvpubadminContentShowGetbyshowlongidAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminContentShowGetbyshowlongidAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminContentShowGetbyshowlongidAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminContentShowGetbyshowlongidAPIResponseModel).Reset()
}

// YunosTvpubadminContentShowGetbyshowlongidAPIResponseModel is 迎客松根据节目longid获取节目元数据 成功返回结果
type YunosTvpubadminContentShowGetbyshowlongidAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_content_show_getbyshowlongid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 节目元数据信息
	Object string `json:"object,omitempty" xml:"object,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminContentShowGetbyshowlongidAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Object = ""
}

var poolYunosTvpubadminContentShowGetbyshowlongidAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminContentShowGetbyshowlongidAPIResponse)
	},
}

// GetYunosTvpubadminContentShowGetbyshowlongidAPIResponse 从 sync.Pool 获取 YunosTvpubadminContentShowGetbyshowlongidAPIResponse
func GetYunosTvpubadminContentShowGetbyshowlongidAPIResponse() *YunosTvpubadminContentShowGetbyshowlongidAPIResponse {
	return poolYunosTvpubadminContentShowGetbyshowlongidAPIResponse.Get().(*YunosTvpubadminContentShowGetbyshowlongidAPIResponse)
}

// ReleaseYunosTvpubadminContentShowGetbyshowlongidAPIResponse 将 YunosTvpubadminContentShowGetbyshowlongidAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminContentShowGetbyshowlongidAPIResponse(v *YunosTvpubadminContentShowGetbyshowlongidAPIResponse) {
	v.Reset()
	poolYunosTvpubadminContentShowGetbyshowlongidAPIResponse.Put(v)
}
