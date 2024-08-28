package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytUploadretail 门店销售扫码回传接口
// alibaba.alihealth.drug.kyt.uploadretail
//
// 门店在销售给顾客时，扫描追溯码的数据按照单据回传；
func AlibabaAlihealthDrugKytUploadretail(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytUploadretailAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytUploadretailAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
