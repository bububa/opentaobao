package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallItemVipAddSchemaGetAPIResponse vip商家发布商品的获取规则接口 API返回值
// tmall.item.vip.add.schema.get
//
// 获取vip商家发布商品的规则
type TmallItemVipAddSchemaGetAPIResponse struct {
	model.CommonResponse
	TmallItemVipAddSchemaGetAPIResponseModel
}

// TmallItemVipAddSchemaGetAPIResponseModel is vip商家发布商品的获取规则接口 成功返回结果
type TmallItemVipAddSchemaGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_vip_add_schema_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值是发布商品时需要的字段及基本类型
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}
