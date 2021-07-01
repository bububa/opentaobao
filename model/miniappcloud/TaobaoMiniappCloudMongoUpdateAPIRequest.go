package miniappcloud

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMiniappCloudMongoUpdateAPIRequest
更新MongoDB中的数据 API请求
taobao.miniapp.cloud.mongo.update

更新MongoDB中的数据 */
type TaobaoMiniappCloudMongoUpdateAPIRequest struct {
	model.Params
	// MongoDB表名
	_collection string
	// 更新条件
	_filter string
	// 待写入的数据
	_record string
	// 要操作的环境，默认是测试环境
	_env string
}

// New
