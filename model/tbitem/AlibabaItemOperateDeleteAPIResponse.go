package tbitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaItemOperateDeleteAPIResponse 商品删除 API返回值
// alibaba.item.operate.delete
//
// 商品删除
type AlibabaItemOperateDeleteAPIResponse struct {
	model.CommonResponse
	AlibabaItemOperateDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaItemOperateDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaItemOperateDeleteAPIResponseModel).Reset()
}

// AlibabaItemOperateDeleteAPIResponseModel is 商品删除 成功返回结果
type AlibabaItemOperateDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_item_operate_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品删除是否成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaItemOperateDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolAlibabaItemOperateDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaItemOperateDeleteAPIResponse)
	},
}

// GetAlibabaItemOperateDeleteAPIResponse 从 sync.Pool 获取 AlibabaItemOperateDeleteAPIResponse
func GetAlibabaItemOperateDeleteAPIResponse() *AlibabaItemOperateDeleteAPIResponse {
	return poolAlibabaItemOperateDeleteAPIResponse.Get().(*AlibabaItemOperateDeleteAPIResponse)
}

// ReleaseAlibabaItemOperateDeleteAPIResponse 将 AlibabaItemOperateDeleteAPIResponse 保存到 sync.Pool
func ReleaseAlibabaItemOperateDeleteAPIResponse(v *AlibabaItemOperateDeleteAPIResponse) {
	v.Reset()
	poolAlibabaItemOperateDeleteAPIResponse.Put(v)
}
