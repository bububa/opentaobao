package icbu

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuPhotobankGroupOperateAPIResponse 图片银行分组操作接口 API返回值
// alibaba.icbu.photobank.group.operate
//
// 修改用户图片银行的分组信息，包括 新增分组，删除分组，分组重命名
type AlibabaIcbuPhotobankGroupOperateAPIResponse struct {
	model.CommonResponse
	AlibabaIcbuPhotobankGroupOperateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIcbuPhotobankGroupOperateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIcbuPhotobankGroupOperateAPIResponseModel).Reset()
}

// AlibabaIcbuPhotobankGroupOperateAPIResponseModel is 图片银行分组操作接口 成功返回结果
type AlibabaIcbuPhotobankGroupOperateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_photobank_group_operate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回的数据结果
	PhotoGroupResult *PhotoGroupResult `json:"photo_group_result,omitempty" xml:"photo_group_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIcbuPhotobankGroupOperateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.PhotoGroupResult = nil
}

var poolAlibabaIcbuPhotobankGroupOperateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIcbuPhotobankGroupOperateAPIResponse)
	},
}

// GetAlibabaIcbuPhotobankGroupOperateAPIResponse 从 sync.Pool 获取 AlibabaIcbuPhotobankGroupOperateAPIResponse
func GetAlibabaIcbuPhotobankGroupOperateAPIResponse() *AlibabaIcbuPhotobankGroupOperateAPIResponse {
	return poolAlibabaIcbuPhotobankGroupOperateAPIResponse.Get().(*AlibabaIcbuPhotobankGroupOperateAPIResponse)
}

// ReleaseAlibabaIcbuPhotobankGroupOperateAPIResponse 将 AlibabaIcbuPhotobankGroupOperateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIcbuPhotobankGroupOperateAPIResponse(v *AlibabaIcbuPhotobankGroupOperateAPIResponse) {
	v.Reset()
	poolAlibabaIcbuPhotobankGroupOperateAPIResponse.Put(v)
}
