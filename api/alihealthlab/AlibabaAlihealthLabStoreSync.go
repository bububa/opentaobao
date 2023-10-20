package alihealthlab

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthlab"
)

// AlibabaAlihealthLabStoreSync 阿里健康检验检测业务，isv门店同步到健康
// alibaba.alihealth.lab.store.sync
//
// 阿里健康检验检测业务，isv门店同步到健康。支持门店的上线、下线操作
func AlibabaAlihealthLabStoreSync(clt *core.SDKClient, req *alihealthlab.AlibabaAlihealthLabStoreSyncAPIRequest, resp *alihealthlab.AlibabaAlihealthLabStoreSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
