package miniapp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMiniappCloudStoreRelationAddAPIRequest
云存储添加 API请求
taobao.miniapp.cloud.store.relation.add

用于用户上传文件之后回写云存储的关联关系 */
type TaobaoMiniappCloudStoreRelationAddAPIRequest struct {
	model.Params
	// 环境 test/online
	_env string
	// 文件类型 image/audio/video/font/other
	_fileType string
	// 上传平台返回的文件唯一ID
	_specialId string
	// 上传平台返回的文件url，部分文件类型无固定url，非必填
	_url string
	// 文件路径
	_cloudPath string
}

// New
