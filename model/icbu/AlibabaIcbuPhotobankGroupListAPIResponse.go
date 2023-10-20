package icbu

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuPhotobankGroupListAPIResponse 图片银行分组信息获取 API返回值
// alibaba.icbu.photobank.group.list
//
// 图片银行分组信息获取
type AlibabaIcbuPhotobankGroupListAPIResponse struct {
	model.CommonResponse
	AlibabaIcbuPhotobankGroupListAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIcbuPhotobankGroupListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIcbuPhotobankGroupListAPIResponseModel).Reset()
}

// AlibabaIcbuPhotobankGroupListAPIResponseModel is 图片银行分组信息获取 成功返回结果
type AlibabaIcbuPhotobankGroupListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_photobank_group_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// groups
	Groups []PhotoAlbumGroup `json:"groups,omitempty" xml:"groups>photo_album_group,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIcbuPhotobankGroupListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Groups = m.Groups[:0]
}

var poolAlibabaIcbuPhotobankGroupListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIcbuPhotobankGroupListAPIResponse)
	},
}

// GetAlibabaIcbuPhotobankGroupListAPIResponse 从 sync.Pool 获取 AlibabaIcbuPhotobankGroupListAPIResponse
func GetAlibabaIcbuPhotobankGroupListAPIResponse() *AlibabaIcbuPhotobankGroupListAPIResponse {
	return poolAlibabaIcbuPhotobankGroupListAPIResponse.Get().(*AlibabaIcbuPhotobankGroupListAPIResponse)
}

// ReleaseAlibabaIcbuPhotobankGroupListAPIResponse 将 AlibabaIcbuPhotobankGroupListAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIcbuPhotobankGroupListAPIResponse(v *AlibabaIcbuPhotobankGroupListAPIResponse) {
	v.Reset()
	poolAlibabaIcbuPhotobankGroupListAPIResponse.Put(v)
}
