package iot

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsAligenieAlbumsGetAPIResponse 专辑详情 API返回值
// alibaba.ailabs.aligenie.albums.get
//
// 给予厂商查询专辑下的音频详情
type AlibabaAilabsAligenieAlbumsGetAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsAligenieAlbumsGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAilabsAligenieAlbumsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAilabsAligenieAlbumsGetAPIResponseModel).Reset()
}

// AlibabaAilabsAligenieAlbumsGetAPIResponseModel is 专辑详情 成功返回结果
type AlibabaAilabsAligenieAlbumsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_aligenie_albums_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAilabsAligenieAlbumsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAilabsAligenieAlbumsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsAligenieAlbumsGetAPIResponse)
	},
}

// GetAlibabaAilabsAligenieAlbumsGetAPIResponse 从 sync.Pool 获取 AlibabaAilabsAligenieAlbumsGetAPIResponse
func GetAlibabaAilabsAligenieAlbumsGetAPIResponse() *AlibabaAilabsAligenieAlbumsGetAPIResponse {
	return poolAlibabaAilabsAligenieAlbumsGetAPIResponse.Get().(*AlibabaAilabsAligenieAlbumsGetAPIResponse)
}

// ReleaseAlibabaAilabsAligenieAlbumsGetAPIResponse 将 AlibabaAilabsAligenieAlbumsGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAilabsAligenieAlbumsGetAPIResponse(v *AlibabaAilabsAligenieAlbumsGetAPIResponse) {
	v.Reset()
	poolAlibabaAilabsAligenieAlbumsGetAPIResponse.Put(v)
}
