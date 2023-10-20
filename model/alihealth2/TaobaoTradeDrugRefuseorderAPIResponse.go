package alihealth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTradeDrugRefuseorderAPIResponse 阿里健康020拒单 API返回值
// taobao.trade.drug.refuseorder
//
// 阿里健康020拒单
type TaobaoTradeDrugRefuseorderAPIResponse struct {
	model.CommonResponse
	TaobaoTradeDrugRefuseorderAPIResponseModel
}

// TaobaoTradeDrugRefuseorderAPIResponseModel is 阿里健康020拒单 成功返回结果
type TaobaoTradeDrugRefuseorderAPIResponseModel struct {
	XMLName xml.Name `xml:"trade_drug_refuseorder_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// life的返回值
	Result *LifeResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
