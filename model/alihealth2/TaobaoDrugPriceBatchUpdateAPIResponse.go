package alihealth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaodrugpricebatchupdateAPIResponse 商家批量更新宝贝价格 API返回值
// taobao.drug.price.batch.update
//
// 商家批量更新宝贝价格
type TaobaodrugpricebatchupdateAPIResponse struct {
	model.CommonResponse
	TaobaodrugpricebatchupdateAPIResponseModel
}

// TaobaodrugpricebatchupdateAPIResponseModel is 商家批量更新宝贝价格 成功返回结果
type TaobaodrugpricebatchupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"drug_price_batch_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaodrugpricebatchupdateResult `json:"result,omitempty" xml:"result,omitempty"`
}
