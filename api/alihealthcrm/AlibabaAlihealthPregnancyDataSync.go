package alihealthcrm

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthcrm"
)

// AlibabaAlihealthPregnancyDataSync 四类数据同步
// alibaba.alihealth.pregnancy.data.sync
//
// 经期调整；基础体温；排卵试纸；B超测排数据同步
func AlibabaAlihealthPregnancyDataSync(ctx context.Context, clt *core.SDKClient, req *alihealthcrm.AlibabaAlihealthPregnancyDataSyncAPIRequest, resp *alihealthcrm.AlibabaAlihealthPregnancyDataSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
