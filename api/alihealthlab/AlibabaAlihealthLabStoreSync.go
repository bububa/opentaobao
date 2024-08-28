package alihealthlab

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthlab"
)

// AlibabaAlihealthLabStoreSync 阿里健康检验检测业务，isv门店同步到健康
// alibaba.alihealth.lab.store.sync
//
// 阿里健康检验检测业务，isv门店同步到健康。支持门店的上线、下线操作
func AlibabaAlihealthLabStoreSync(ctx context.Context, clt *core.SDKClient, req *alihealthlab.AlibabaAlihealthLabStoreSyncAPIRequest, resp *alihealthlab.AlibabaAlihealthLabStoreSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
