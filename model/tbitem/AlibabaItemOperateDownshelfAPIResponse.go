package tbitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaItemOperateDownshelfAPIResponse 商品下架 API返回值
// alibaba.item.operate.downshelf
//
// 商品下架
type AlibabaItemOperateDownshelfAPIResponse struct {
	model.CommonResponse
	AlibabaItemOperateDownshelfAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaItemOperateDownshelfAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaItemOperateDownshelfAPIResponseModel).Reset()
}

// AlibabaItemOperateDownshelfAPIResponseModel is 商品下架 成功返回结果
type AlibabaItemOperateDownshelfAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_item_operate_downshelf_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品下架是否成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaItemOperateDownshelfAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolAlibabaItemOperateDownshelfAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaItemOperateDownshelfAPIResponse)
	},
}

// GetAlibabaItemOperateDownshelfAPIResponse 从 sync.Pool 获取 AlibabaItemOperateDownshelfAPIResponse
func GetAlibabaItemOperateDownshelfAPIResponse() *AlibabaItemOperateDownshelfAPIResponse {
	return poolAlibabaItemOperateDownshelfAPIResponse.Get().(*AlibabaItemOperateDownshelfAPIResponse)
}

// ReleaseAlibabaItemOperateDownshelfAPIResponse 将 AlibabaItemOperateDownshelfAPIResponse 保存到 sync.Pool
func ReleaseAlibabaItemOperateDownshelfAPIResponse(v *AlibabaItemOperateDownshelfAPIResponse) {
	v.Reset()
	poolAlibabaItemOperateDownshelfAPIResponse.Put(v)
}
