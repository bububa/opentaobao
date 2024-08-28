package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugcodeDrugfactoryBlindfileGetbatchinfo 获取盲底文件中的批次信息
// alibaba.alihealth.drugcode.drugfactory.blindfile.getbatchinfo
//
// 获取盲底文件中的批次信息
func AlibabaAlihealthDrugcodeDrugfactoryBlindfileGetbatchinfo(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugcodeDrugfactoryBlindfileGetbatchinfoAPIRequest, resp *drugtrace.AlibabaAlihealthDrugcodeDrugfactoryBlindfileGetbatchinfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
