package miniappcloud

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMiniappCloudStoreListfileAPIRequest
云存储根据文件名反查地址 API请求
taobao.miniapp.cloud.store.listfile

云存储中，根据文件名反查地址 */
type TaobaoMiniappCloudStoreListfileAPIRequest struct {
	model.Params
	// 环境(online/test)
	_env string
	// 文件全路径名
	_filePath string
}

// New
