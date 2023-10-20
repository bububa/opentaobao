package tbitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaItemOperateUpshelfAPIResponse 商品上架 API返回值
// alibaba.item.operate.upshelf
//
// 商品上架
type AlibabaItemOperateUpshelfAPIResponse struct {
	model.CommonResponse
	AlibabaItemOperateUpshelfAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaItemOperateUpshelfAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaItemOperateUpshelfAPIResponseModel).Reset()
}

// AlibabaItemOperateUpshelfAPIResponseModel is 商品上架 成功返回结果
type AlibabaItemOperateUpshelfAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_item_operate_upshelf_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品上架是否成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaItemOperateUpshelfAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolAlibabaItemOperateUpshelfAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaItemOperateUpshelfAPIResponse)
	},
}

// GetAlibabaItemOperateUpshelfAPIResponse 从 sync.Pool 获取 AlibabaItemOperateUpshelfAPIResponse
func GetAlibabaItemOperateUpshelfAPIResponse() *AlibabaItemOperateUpshelfAPIResponse {
	return poolAlibabaItemOperateUpshelfAPIResponse.Get().(*AlibabaItemOperateUpshelfAPIResponse)
}

// ReleaseAlibabaItemOperateUpshelfAPIResponse 将 AlibabaItemOperateUpshelfAPIResponse 保存到 sync.Pool
func ReleaseAlibabaItemOperateUpshelfAPIResponse(v *AlibabaItemOperateUpshelfAPIResponse) {
	v.Reset()
	poolAlibabaItemOperateUpshelfAPIResponse.Put(v)
}
