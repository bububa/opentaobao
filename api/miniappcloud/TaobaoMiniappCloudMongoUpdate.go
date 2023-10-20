package miniappcloud

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappcloud"
)

// Taobaominiappcloudmongoupdate 更新MongoDB中的数据
// taobao.miniapp.cloud.mongo.update
//
// 更新MongoDB中的数据
func Taobaominiappcloudmongoupdate(clt *core.SDKClient, req *miniappcloud.TaobaominiappcloudmongoupdateAPIRequest, session string) (*miniappcloud.TaobaominiappcloudmongoupdateAPIResponse, error) {
	var resp miniappcloud.TaobaominiappcloudmongoupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
