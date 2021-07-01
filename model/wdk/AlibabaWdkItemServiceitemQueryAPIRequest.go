package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkItemServiceitemQueryAPIRequest
查询服务商品 API请求
alibaba.wdk.item.serviceitem.query

查询服务商品 */
type AlibabaWdkItemServiceitemQueryAPIRequest struct {
	model.Params
	// 后台类目
	_hemaCategoryId string
	// 机构编码
	_orgNo string
}

// New
