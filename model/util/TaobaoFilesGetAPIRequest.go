package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFilesGetAPIRequest
业务文件获取 API请求
taobao.files.get

获取业务方暂存给ISV的文件列表 */
type TaobaoFilesGetAPIRequest struct {
	model.Params
	// 搜索开始时间
	_startDate string
	// 搜索结束时间
	_endDate string
	// 下载链接状态。1:未下载。2:已下载
	_status int64
}

// New
