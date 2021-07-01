package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMediaFileAddAPIRequest
多媒体平台文件添加 API请求
taobao.media.file.add

用户通过top上传文件到多媒体平台 */
type TaobaoMediaFileAddAPIRequest struct {
	model.Params
	// 文件属于的那个目录的目录编号
	_dirId int64
	// 上传文件的名称
	_name string
	// 额外信息
	_ext int64
	// 文件
	_fileData *model.File
	// 接入多媒体平台的业务code每个应用有一个特有的业务code
	_bizCode string
}

// New
