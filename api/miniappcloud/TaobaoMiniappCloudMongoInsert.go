package miniappcloud

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappcloud"
)

/* TaobaoMiniappCloudMongoInsert
MongoDB插入单条数据
taobao.miniapp.cloud.mongo.insert

向商家应用云中插入一条记录，用于外部数据同步到应用中 */
func TaobaoMiniappCloudMongoInsert(clt *core.SDKClient, req *miniappcloud.TaobaoMiniappCloudMongoInsertAPIRequest, session string) (*miniappcloud.TaobaoMiniappCloudMongoInsertAPIResponse, error) {
	var resp miniappcloud.TaobaoMiniappCloudMongoInsertAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
