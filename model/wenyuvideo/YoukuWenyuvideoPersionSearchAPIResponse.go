package wenyuvideo

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YoukuWenyuvideoPersionSearchAPIResponse 根据人物名称查询人物列表 API返回值
// youku.wenyuvideo.persion.search
//
// 根据人物名称查询人物列表
type YoukuWenyuvideoPersionSearchAPIResponse struct {
	model.CommonResponse
	YoukuWenyuvideoPersionSearchAPIResponseModel
}

// Reset 清空结构体
func (m *YoukuWenyuvideoPersionSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YoukuWenyuvideoPersionSearchAPIResponseModel).Reset()
}

// YoukuWenyuvideoPersionSearchAPIResponseModel is 根据人物名称查询人物列表 成功返回结果
type YoukuWenyuvideoPersionSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"youku_wenyuvideo_persion_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *YoukuWenyuvideoPersionSearchResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *YoukuWenyuvideoPersionSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolYoukuWenyuvideoPersionSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(YoukuWenyuvideoPersionSearchAPIResponse)
	},
}

// GetYoukuWenyuvideoPersionSearchAPIResponse 从 sync.Pool 获取 YoukuWenyuvideoPersionSearchAPIResponse
func GetYoukuWenyuvideoPersionSearchAPIResponse() *YoukuWenyuvideoPersionSearchAPIResponse {
	return poolYoukuWenyuvideoPersionSearchAPIResponse.Get().(*YoukuWenyuvideoPersionSearchAPIResponse)
}

// ReleaseYoukuWenyuvideoPersionSearchAPIResponse 将 YoukuWenyuvideoPersionSearchAPIResponse 保存到 sync.Pool
func ReleaseYoukuWenyuvideoPersionSearchAPIResponse(v *YoukuWenyuvideoPersionSearchAPIResponse) {
	v.Reset()
	poolYoukuWenyuvideoPersionSearchAPIResponse.Put(v)
}
