package miniappcloud

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMiniappCloudMongoInsertAPIRequest
MongoDB插入单条数据 API请求
taobao.miniapp.cloud.mongo.insert

向商家应用云中插入一条记录，用于外部数据同步到应用中 */
type TaobaoMiniappCloudMongoInsertAPIRequest struct {
	model.Params
	// 待插入的数据，JSON格式
	_record string
	// MongoDB表名
	_collection string
	// 要操作的环境，默认是测试环境
	_env string
}

// New
