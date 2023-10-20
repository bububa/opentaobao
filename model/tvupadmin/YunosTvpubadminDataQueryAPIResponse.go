package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDataQueryAPIResponse 魔盒统计数据查询接口 API返回值
// yunos.tvpubadmin.data.query
//
// 用于华数查询魔盒上的一些用户统计数据
type YunosTvpubadminDataQueryAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminDataQueryAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminDataQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminDataQueryAPIResponseModel).Reset()
}

// YunosTvpubadminDataQueryAPIResponseModel is 魔盒统计数据查询接口 成功返回结果
type YunosTvpubadminDataQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_data_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *DfPageResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminDataQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolYunosTvpubadminDataQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminDataQueryAPIResponse)
	},
}

// GetYunosTvpubadminDataQueryAPIResponse 从 sync.Pool 获取 YunosTvpubadminDataQueryAPIResponse
func GetYunosTvpubadminDataQueryAPIResponse() *YunosTvpubadminDataQueryAPIResponse {
	return poolYunosTvpubadminDataQueryAPIResponse.Get().(*YunosTvpubadminDataQueryAPIResponse)
}

// ReleaseYunosTvpubadminDataQueryAPIResponse 将 YunosTvpubadminDataQueryAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminDataQueryAPIResponse(v *YunosTvpubadminDataQueryAPIResponse) {
	v.Reset()
	poolYunosTvpubadminDataQueryAPIResponse.Put(v)
}
