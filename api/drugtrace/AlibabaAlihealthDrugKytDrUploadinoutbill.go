package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytDrUploadinoutbill 疫苗企业出入库上传
// alibaba.alihealth.drug.kyt.dr.uploadinoutbill
//
// 零售企业上传出入库信息，包括采购入库（102），退货入库（103），供应入库（107）,退货出库（202），销毁出库（205），抽检出库（206）， 供应出库（209）,
// 不包括对个人的零售出库，疫苗接种，领药出库。
func AlibabaAlihealthDrugKytDrUploadinoutbill(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytDrUploadinoutbillAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytDrUploadinoutbillAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
