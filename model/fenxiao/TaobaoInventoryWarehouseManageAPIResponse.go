package fenxiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoInventoryWarehouseManageAPIResponse
创建商家仓或者更新商家仓信息 API返回值
taobao.inventory.warehouse.manage

创建商家仓或者更新商家仓信息 */
type TaobaoInventoryWarehouseManageAPIResponse struct {
	model.CommonResponse
	TaobaoInventoryWarehouseManageAPIResponseModel
}

// TaobaoInventoryWarehouseManageAPIResponseModel is 创建商家仓或者更新商家仓信息 成功返回结果
type TaobaoInventoryWarehouseManageAPIResponseModel struct {
	XMLName xml.Name `xml:"inventory_warehouse_manage_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoInventoryWarehouseManageResult `json:"result,omitempty" xml:"result,omitempty"`
}
