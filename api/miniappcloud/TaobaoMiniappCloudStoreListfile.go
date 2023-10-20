package miniappcloud

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappcloud"
)

// Taobaominiappcloudstorelistfile 云存储根据文件名反查地址
// taobao.miniapp.cloud.store.listfile
//
// 云存储中，根据文件名反查地址
func Taobaominiappcloudstorelistfile(clt *core.SDKClient, req *miniappcloud.TaobaominiappcloudstorelistfileAPIRequest, session string) (*miniappcloud.TaobaominiappcloudstorelistfileAPIResponse, error) {
	var resp miniappcloud.TaobaominiappcloudstorelistfileAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
