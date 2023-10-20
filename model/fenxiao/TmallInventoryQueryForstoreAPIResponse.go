package fenxiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallinventoryqueryforstoreAPIResponse 查询后端商品仓库库存 API返回值
// tmall.inventory.query.forstore
//
// 商家查询后端商品仓库库存
type TmallinventoryqueryforstoreAPIResponse struct {
	model.CommonResponse
	TmallinventoryqueryforstoreAPIResponseModel
}

// TmallinventoryqueryforstoreAPIResponseModel is 查询后端商品仓库库存 成功返回结果
type TmallinventoryqueryforstoreAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_inventory_query_forstore_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误code
	Errorcode string `json:"errorcode,omitempty" xml:"errorcode,omitempty"`
	// 错误信息
	Errormessage string `json:"errormessage,omitempty" xml:"errormessage,omitempty"`
	// 查询结果
	Result *InventoryQueryResult `json:"result,omitempty" xml:"result,omitempty"`
	// 整体成功或失败
	Issuccess bool `json:"issuccess,omitempty" xml:"issuccess,omitempty"`
}
