package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemVipUpdateSchemaGetAPIResponse vip商家编辑商品的规则获取接口 API返回值
// tmall.item.vip.update.schema.get
//
// 获取vip商家编辑商品的规则
type TmallItemVipUpdateSchemaGetAPIResponse struct {
	model.CommonResponse
	TmallItemVipUpdateSchemaGetAPIResponseModel
}

// Reset 清空结构体
func (m *TmallItemVipUpdateSchemaGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallItemVipUpdateSchemaGetAPIResponseModel).Reset()
}

// TmallItemVipUpdateSchemaGetAPIResponseModel is vip商家编辑商品的规则获取接口 成功返回结果
type TmallItemVipUpdateSchemaGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_vip_update_schema_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 获取的编辑商品的规则
	UpdateGetResult string `json:"update_get_result,omitempty" xml:"update_get_result,omitempty"`
}

// Reset 清空结构体
func (m *TmallItemVipUpdateSchemaGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.UpdateGetResult = ""
}

var poolTmallItemVipUpdateSchemaGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallItemVipUpdateSchemaGetAPIResponse)
	},
}

// GetTmallItemVipUpdateSchemaGetAPIResponse 从 sync.Pool 获取 TmallItemVipUpdateSchemaGetAPIResponse
func GetTmallItemVipUpdateSchemaGetAPIResponse() *TmallItemVipUpdateSchemaGetAPIResponse {
	return poolTmallItemVipUpdateSchemaGetAPIResponse.Get().(*TmallItemVipUpdateSchemaGetAPIResponse)
}

// ReleaseTmallItemVipUpdateSchemaGetAPIResponse 将 TmallItemVipUpdateSchemaGetAPIResponse 保存到 sync.Pool
func ReleaseTmallItemVipUpdateSchemaGetAPIResponse(v *TmallItemVipUpdateSchemaGetAPIResponse) {
	v.Reset()
	poolTmallItemVipUpdateSchemaGetAPIResponse.Put(v)
}
