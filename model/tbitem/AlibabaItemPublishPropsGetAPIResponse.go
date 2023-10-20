package tbitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaItemPublishPropsGetAPIResponse 商品级联属性信息获取 API返回值
// alibaba.item.publish.props.get
//
// 新商品发布，商品级联属性信息获取
type AlibabaItemPublishPropsGetAPIResponse struct {
	model.CommonResponse
	AlibabaItemPublishPropsGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaItemPublishPropsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaItemPublishPropsGetAPIResponseModel).Reset()
}

// AlibabaItemPublishPropsGetAPIResponseModel is 商品级联属性信息获取 成功返回结果
type AlibabaItemPublishPropsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_item_publish_props_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品发布规则信息，XML格式.
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaItemPublishPropsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolAlibabaItemPublishPropsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaItemPublishPropsGetAPIResponse)
	},
}

// GetAlibabaItemPublishPropsGetAPIResponse 从 sync.Pool 获取 AlibabaItemPublishPropsGetAPIResponse
func GetAlibabaItemPublishPropsGetAPIResponse() *AlibabaItemPublishPropsGetAPIResponse {
	return poolAlibabaItemPublishPropsGetAPIResponse.Get().(*AlibabaItemPublishPropsGetAPIResponse)
}

// ReleaseAlibabaItemPublishPropsGetAPIResponse 将 AlibabaItemPublishPropsGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaItemPublishPropsGetAPIResponse(v *AlibabaItemPublishPropsGetAPIResponse) {
	v.Reset()
	poolAlibabaItemPublishPropsGetAPIResponse.Put(v)
}
