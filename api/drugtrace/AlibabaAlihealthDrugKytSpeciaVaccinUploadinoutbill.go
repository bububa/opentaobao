package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytspeciavaccinuploadinoutbill 疫苗企业出入库上传
// alibaba.alihealth.drug.kyt.specia.vaccin.uploadinoutbill
//
// 零售企业上传出入库信息，包括采购入库（102），退货入库（103），供应入库（107）,退货出库（202），销毁出库（205），抽检出库（206）， 供应出库（209）,
// 不包括对个人的零售出库，疫苗接种，领药出库。
func Alibabaalihealthdrugkytspeciavaccinuploadinoutbill(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytspeciavaccinuploadinoutbillAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytspeciavaccinuploadinoutbillAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytspeciavaccinuploadinoutbillAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
