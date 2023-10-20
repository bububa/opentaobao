package alihealth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaodrugquantitybatchupdateAPIResponse 批量同步库存接口 API返回值
// taobao.drug.quantity.batch.update
//
// 商家通过top接口可以批量修改商品库存
type TaobaodrugquantitybatchupdateAPIResponse struct {
	model.CommonResponse
	TaobaodrugquantitybatchupdateAPIResponseModel
}

// TaobaodrugquantitybatchupdateAPIResponseModel is 批量同步库存接口 成功返回结果
type TaobaodrugquantitybatchupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"drug_quantity_batch_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaodrugquantitybatchupdateResult `json:"result,omitempty" xml:"result,omitempty"`
}
