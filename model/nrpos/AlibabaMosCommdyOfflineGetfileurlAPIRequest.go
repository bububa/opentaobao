package nrpos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMosCommdyOfflineGetfileurlAPIRequest
去前置机pos商品离线文件下载地址查询接口 API请求
alibaba.mos.commdy.offline.getfileurl

去前置机-pos查询离线文件下载地址接口 */
type AlibabaMosCommdyOfflineGetfileurlAPIRequest struct {
	model.Params
	// 离线文件名称
	_fileKeys []string
}

// New
