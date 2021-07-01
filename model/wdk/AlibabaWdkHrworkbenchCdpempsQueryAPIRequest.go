package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkHrworkbenchCdpempsQueryAPIRequest
homs员工信息核对查询服务 API请求
alibaba.wdk.hrworkbench.cdpemps.query

给盒马可靠软件服务商Cdp系统，做非阿里编员工数据一致性核对检查 */
type AlibabaWdkHrworkbenchCdpempsQueryAPIRequest struct {
	model.Params
	// 页面大小
	_pageSize int64
	// 业务授权key
	_bizKey string
	// 业务授权code
	_bizCode string
	// 起始页
	_currentPage int64
}

// New
