package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIcbuPhotobankGroupListAPIRequest
图片银行分组信息获取 API请求
alibaba.icbu.photobank.group.list

图片银行分组信息获取 */
type AlibabaIcbuPhotobankGroupListAPIRequest struct {
	model.Params
	// 补充信息
	_extraContext string
	// 查询图片分组信息，如果传入id，则获取当前分组和所有子分组信息，否则获取所有一级分组信息
	_id int64
}

// New
