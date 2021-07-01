package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIcbuCategoryGetAPIRequest
商品发布类目获取 API请求
alibaba.icbu.category.get

获取商品发布类目 */
type AlibabaIcbuCategoryGetAPIRequest struct {
	model.Params
	// 发布类目id,必须大于等于0， 如果为0，则查询所有一级类目
	_catId int64
}

// New
