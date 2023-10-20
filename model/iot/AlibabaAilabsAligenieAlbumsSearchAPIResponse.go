package iot

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsAligenieAlbumsSearchAPIResponse 查询专辑 API返回值
// alibaba.ailabs.aligenie.albums.search
//
// 搜索类目下的专辑信息
type AlibabaAilabsAligenieAlbumsSearchAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsAligenieAlbumsSearchAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAilabsAligenieAlbumsSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAilabsAligenieAlbumsSearchAPIResponseModel).Reset()
}

// AlibabaAilabsAligenieAlbumsSearchAPIResponseModel is 查询专辑 成功返回结果
type AlibabaAilabsAligenieAlbumsSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_aligenie_albums_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAilabsAligenieAlbumsSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAilabsAligenieAlbumsSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsAligenieAlbumsSearchAPIResponse)
	},
}

// GetAlibabaAilabsAligenieAlbumsSearchAPIResponse 从 sync.Pool 获取 AlibabaAilabsAligenieAlbumsSearchAPIResponse
func GetAlibabaAilabsAligenieAlbumsSearchAPIResponse() *AlibabaAilabsAligenieAlbumsSearchAPIResponse {
	return poolAlibabaAilabsAligenieAlbumsSearchAPIResponse.Get().(*AlibabaAilabsAligenieAlbumsSearchAPIResponse)
}

// ReleaseAlibabaAilabsAligenieAlbumsSearchAPIResponse 将 AlibabaAilabsAligenieAlbumsSearchAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAilabsAligenieAlbumsSearchAPIResponse(v *AlibabaAilabsAligenieAlbumsSearchAPIResponse) {
	v.Reset()
	poolAlibabaAilabsAligenieAlbumsSearchAPIResponse.Put(v)
}
