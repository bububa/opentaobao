package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemVipSchemaUpdateAPIResponse 大商家商品编辑接口 API返回值
// tmall.item.vip.schema.update
//
// 大商家编辑商品
type TmallItemVipSchemaUpdateAPIResponse struct {
	model.CommonResponse
	TmallItemVipSchemaUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TmallItemVipSchemaUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallItemVipSchemaUpdateAPIResponseModel).Reset()
}

// TmallItemVipSchemaUpdateAPIResponseModel is 大商家商品编辑接口 成功返回结果
type TmallItemVipSchemaUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_vip_schema_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 编辑商品的id
	UpdateItemResult string `json:"update_item_result,omitempty" xml:"update_item_result,omitempty"`
	// 编辑商品操作成功时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// sku与outerId映射信息
	SkuMapJson string `json:"sku_map_json,omitempty" xml:"sku_map_json,omitempty"`
}

// Reset 清空结构体
func (m *TmallItemVipSchemaUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.UpdateItemResult = ""
	m.GmtCreate = ""
	m.SkuMapJson = ""
}

var poolTmallItemVipSchemaUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallItemVipSchemaUpdateAPIResponse)
	},
}

// GetTmallItemVipSchemaUpdateAPIResponse 从 sync.Pool 获取 TmallItemVipSchemaUpdateAPIResponse
func GetTmallItemVipSchemaUpdateAPIResponse() *TmallItemVipSchemaUpdateAPIResponse {
	return poolTmallItemVipSchemaUpdateAPIResponse.Get().(*TmallItemVipSchemaUpdateAPIResponse)
}

// ReleaseTmallItemVipSchemaUpdateAPIResponse 将 TmallItemVipSchemaUpdateAPIResponse 保存到 sync.Pool
func ReleaseTmallItemVipSchemaUpdateAPIResponse(v *TmallItemVipSchemaUpdateAPIResponse) {
	v.Reset()
	poolTmallItemVipSchemaUpdateAPIResponse.Put(v)
}
