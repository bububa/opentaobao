package fenxiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoinventoryadjustexternalAPIResponse 非交易库存调整单 API返回值
// taobao.inventory.adjust.external
//
// 商家非交易调整库存，调拨出库、盘点等时调用
type TaobaoinventoryadjustexternalAPIResponse struct {
	model.CommonResponse
	TaobaoinventoryadjustexternalAPIResponseModel
}

// TaobaoinventoryadjustexternalAPIResponseModel is 非交易库存调整单 成功返回结果
type TaobaoinventoryadjustexternalAPIResponseModel struct {
	XMLName xml.Name `xml:"inventory_adjust_external_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 提示信息
	TipInfos []TipInfo `json:"tip_infos,omitempty" xml:"tip_infos>tip_info,omitempty"`
	// 操作返回码
	OperateCode string `json:"operate_code,omitempty" xml:"operate_code,omitempty"`
}
