package lstlogistics

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstShiporderCreateAPIResponse
零售通发货单创建 API返回值
alibaba.lst.shiporder.create

通过该接口可以创建零售通运保保发货单，并处理相关业务流程。 */
type AlibabaLstShiporderCreateAPIResponse struct {
	model.CommonResponse
	AlibabaLstShiporderCreateAPIResponseModel
}

// AlibabaLstShiporderCreateAPIResponseModel is 零售通发货单创建 成功返回结果
type AlibabaLstShiporderCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_shiporder_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
