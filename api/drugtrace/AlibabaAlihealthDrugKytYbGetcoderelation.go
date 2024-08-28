package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytYbGetcoderelation 医保-查询码的所有子码
// alibaba.alihealth.drug.kyt.yb.getcoderelation
//
// 应用于药店或医院入库环节，通过扫码获取下级码进行入库；
// 通过码查询所有子码以及包装比例
func AlibabaAlihealthDrugKytYbGetcoderelation(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytYbGetcoderelationAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytYbGetcoderelationAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
