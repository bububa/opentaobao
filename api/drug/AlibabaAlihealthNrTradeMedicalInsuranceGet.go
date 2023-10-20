package drug

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drug"
)

// AlibabaAlihealthNrTradeMedicalInsuranceGet 阿里健康医保支付信息获取
// alibaba.alihealth.nr.trade.medical.insurance.get
//
// 阿里健康医保支付信息获取
func AlibabaAlihealthNrTradeMedicalInsuranceGet(clt *core.SDKClient, req *drug.AlibabaAlihealthNrTradeMedicalInsuranceGetAPIRequest, resp *drug.AlibabaAlihealthNrTradeMedicalInsuranceGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
