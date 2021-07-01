package fenxiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoInventoryInitialAPIResponse
库存初始化 API返回值
taobao.inventory.initial

建议使用新接口：taobao.inventory.merchant.adjust ，该接口会逐步停用。
商家仓库存初始化接口，直接按照商家指定的商品库存数进行填充，没有单据核对，不参与库存对账。 */
type TaobaoInventoryInitialAPIResponse struct {
	model.CommonResponse
	TaobaoInventoryInitialAPIResponseModel
}

// TaobaoInventoryInitialAPIResponseModel is 库存初始化 成功返回结果
type TaobaoInventoryInitialAPIResponseModel struct {
	XMLName xml.Name `xml:"inventory_initial_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 提示信息
	TipInfos []TipInfo `json:"tip_infos,omitempty" xml:"tip_infos>tip_info,omitempty"`
}
