package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugcodeDrugfactoryTransferblind 传输盲底文件
// alibaba.alihealth.drugcode.drugfactory.transferblind
//
// 临床用药试验-传输盲底文件
func AlibabaAlihealthDrugcodeDrugfactoryTransferblind(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest, resp *drugtrace.AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
