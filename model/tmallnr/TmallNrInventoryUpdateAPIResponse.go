package tmallnr

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallnrinventoryupdateAPIResponse 门店业务同步库存 API返回值
// tmall.nr.inventory.update
//
// 用于商家每日同步更新门店库存
type TmallnrinventoryupdateAPIResponse struct {
	model.CommonResponse
	TmallnrinventoryupdateAPIResponseModel
}

// TmallnrinventoryupdateAPIResponseModel is 门店业务同步库存 成功返回结果
type TmallnrinventoryupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nr_inventory_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *NrResult `json:"result,omitempty" xml:"result,omitempty"`
}
