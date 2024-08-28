package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytYyUploadinoutbill 医院出入库信息上传
// alibaba.alihealth.drug.kyt.yy.uploadinoutbill
//
// 医院上传出入库信息，包括采购入库（102），退货入库（103），供应入库（107）,退货出库（202），销毁出库（205），抽检出库（206）， 供应出库（209）,
// 不包括对个人的零售出库，疫苗接种，领药出库。
func AlibabaAlihealthDrugKytYyUploadinoutbill(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytYyUploadinoutbillAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytYyUploadinoutbillAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
