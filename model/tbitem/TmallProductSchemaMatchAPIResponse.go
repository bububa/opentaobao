package tbitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallProductSchemaMatchAPIResponse product匹配接口 API返回值
// tmall.product.schema.match
//
// 根据tmall.product.match.schema.get获取到的规则，填充相应地的字段值以及类目，匹配符合条件的产品，返回匹配product结果，注意，有可能返回多个产品ID，以逗号分隔（尤其是图书类目）；
type TmallProductSchemaMatchAPIResponse struct {
	model.CommonResponse
	TmallProductSchemaMatchAPIResponseModel
}

// Reset 清空结构体
func (m *TmallProductSchemaMatchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallProductSchemaMatchAPIResponseModel).Reset()
}

// TmallProductSchemaMatchAPIResponseModel is product匹配接口 成功返回结果
type TmallProductSchemaMatchAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_product_schema_match_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回匹配产品ID，部分类目可能返回多个产品ID，以逗号分隔。
	MatchResult string `json:"match_result,omitempty" xml:"match_result,omitempty"`
}

// Reset 清空结构体
func (m *TmallProductSchemaMatchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MatchResult = ""
}

var poolTmallProductSchemaMatchAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallProductSchemaMatchAPIResponse)
	},
}

// GetTmallProductSchemaMatchAPIResponse 从 sync.Pool 获取 TmallProductSchemaMatchAPIResponse
func GetTmallProductSchemaMatchAPIResponse() *TmallProductSchemaMatchAPIResponse {
	return poolTmallProductSchemaMatchAPIResponse.Get().(*TmallProductSchemaMatchAPIResponse)
}

// ReleaseTmallProductSchemaMatchAPIResponse 将 TmallProductSchemaMatchAPIResponse 保存到 sync.Pool
func ReleaseTmallProductSchemaMatchAPIResponse(v *TmallProductSchemaMatchAPIResponse) {
	v.Reset()
	poolTmallProductSchemaMatchAPIResponse.Put(v)
}
