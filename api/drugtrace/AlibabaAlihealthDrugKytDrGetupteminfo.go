package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

/* AlibabaAlihealthDrugKytDrGetupteminfo
获取上游温度信息（疫苗）
alibaba.alihealth.drug.kyt.dr.getupteminfo

根据追溯码及企业ID获取上游运输及存储温度信息（疫苗） */
func AlibabaAlihealthDrugKytDrGetupteminfo(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytDrGetupteminfoAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytDrGetupteminfoAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytDrGetupteminfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
