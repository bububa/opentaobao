package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytGetcodebaseinfo 码的药品信息查询
// alibaba.alihealth.drug.kyt.getcodebaseinfo
//
// 提供根据码查询码基本信息接口
func AlibabaAlihealthDrugKytGetcodebaseinfo(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytGetcodebaseinfoAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytGetcodebaseinfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
