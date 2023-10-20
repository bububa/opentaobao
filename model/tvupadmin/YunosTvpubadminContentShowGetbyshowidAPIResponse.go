package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentShowGetbyshowidAPIResponse 迎客松根据节目id获取节目元数据 API返回值
// yunos.tvpubadmin.content.show.getbyshowid
//
// 迎客松根据节目id获取节目元数据
type YunosTvpubadminContentShowGetbyshowidAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminContentShowGetbyshowidAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminContentShowGetbyshowidAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminContentShowGetbyshowidAPIResponseModel).Reset()
}

// YunosTvpubadminContentShowGetbyshowidAPIResponseModel is 迎客松根据节目id获取节目元数据 成功返回结果
type YunosTvpubadminContentShowGetbyshowidAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_content_show_getbyshowid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 节目元数据
	Object string `json:"object,omitempty" xml:"object,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminContentShowGetbyshowidAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Object = ""
}

var poolYunosTvpubadminContentShowGetbyshowidAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminContentShowGetbyshowidAPIResponse)
	},
}

// GetYunosTvpubadminContentShowGetbyshowidAPIResponse 从 sync.Pool 获取 YunosTvpubadminContentShowGetbyshowidAPIResponse
func GetYunosTvpubadminContentShowGetbyshowidAPIResponse() *YunosTvpubadminContentShowGetbyshowidAPIResponse {
	return poolYunosTvpubadminContentShowGetbyshowidAPIResponse.Get().(*YunosTvpubadminContentShowGetbyshowidAPIResponse)
}

// ReleaseYunosTvpubadminContentShowGetbyshowidAPIResponse 将 YunosTvpubadminContentShowGetbyshowidAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminContentShowGetbyshowidAPIResponse(v *YunosTvpubadminContentShowGetbyshowidAPIResponse) {
	v.Reset()
	poolYunosTvpubadminContentShowGetbyshowidAPIResponse.Put(v)
}
