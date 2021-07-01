package servicecenter

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIResponse
回收商同步前置补贴红包的代扣成功事件 API返回值
taobao.recycle.ofnpreredpacket.tpdeductsuccess

回收商->天猫后端，同步前置补贴红包的代扣成功事件 */
type TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIResponse struct {
	model.CommonResponse
	TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIResponseModel
}

// TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIResponseModel is 回收商同步前置补贴红包的代扣成功事件 成功返回结果
type TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIResponseModel struct {
	XMLName xml.Name `xml:"recycle_ofnpreredpacket_tpdeductsuccess_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作
	Data *OfnPreRedPacketActionDto `json:"data,omitempty" xml:"data,omitempty"`
}
