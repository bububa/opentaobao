package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWholesaleShippinglineTemplateListAPIRequest
获取运费模板 API请求
alibaba.wholesale.shippingline.template.list

查询运费模板信息 */
type AlibabaWholesaleShippinglineTemplateListAPIRequest struct {
	model.Params
	// 第几页从1开始
	_pageNum int64
	// 每页返回的数据个数
	_count int64
}

// New
