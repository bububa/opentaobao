package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeBrokerMigrate 融合店经纪人迁移
// alibaba.alihouse.existinghome.broker.migrate
//
// 融合店经纪人迁移
func AlibabaAlihouseExistinghomeBrokerMigrate(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeBrokerMigrateAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeBrokerMigrateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
