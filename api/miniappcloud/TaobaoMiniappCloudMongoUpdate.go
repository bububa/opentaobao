package miniappcloud

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappcloud"
)

// TaobaoMiniappCloudMongoUpdate 更新MongoDB中的数据
// taobao.miniapp.cloud.mongo.update
//
// 更新MongoDB中的数据
func TaobaoMiniappCloudMongoUpdate(clt *core.SDKClient, req *miniappcloud.TaobaoMiniappCloudMongoUpdateAPIRequest, resp *miniappcloud.TaobaoMiniappCloudMongoUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
