package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugtraceTopLsydUploadinoutbill 出入库单据上传
// alibaba.alihealth.drugtrace.top.lsyd.uploadinoutbill
//
// 零售企业上传出入库信息，包括102, "采购入库"；103, "退货入库"；104, "调拨入库"；107, "供应入库"；108, "召回入库"；110,"赠品入库"；111,"盘盈入库"；112,"报废入库"；113,"其他入库"
// 201, "销售出库"；202, "退货出库"；203, "调拨出库"；204, "返工出库"；205, "销毁出库"；206, "抽检出库"；207, "直调出库"；209, "供应出库"；211, "召回出库"；212,"赠品出库"；214,"盘亏出库"；215,"损坏出库"；216,"报废出库"；217,"其他出库"；237, "直调退货"。
// 不包括对个人的零售出库，疫苗接种，领药出库。
func AlibabaAlihealthDrugtraceTopLsydUploadinoutbill(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugtraceTopLsydUploadinoutbillAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugtraceTopLsydUploadinoutbillAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugtraceTopLsydUploadinoutbillAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
