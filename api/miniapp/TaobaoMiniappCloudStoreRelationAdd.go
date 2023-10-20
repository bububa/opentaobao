package miniapp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniapp"
)

// TaobaoMiniappCloudStoreRelationAdd 云存储添加
// taobao.miniapp.cloud.store.relation.add
//
// 用于用户上传文件之后回写云存储的关联关系
func TaobaoMiniappCloudStoreRelationAdd(clt *core.SDKClient, req *miniapp.TaobaoMiniappCloudStoreRelationAddAPIRequest, resp *miniapp.TaobaoMiniappCloudStoreRelationAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
