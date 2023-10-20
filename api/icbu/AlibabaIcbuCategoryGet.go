package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// AlibabaIcbuCategoryGet 商品发布类目获取
// alibaba.icbu.category.get
//
// 获取商品发布类目
func AlibabaIcbuCategoryGet(clt *core.SDKClient, req *icbu.AlibabaIcbuCategoryGetAPIRequest, resp *icbu.AlibabaIcbuCategoryGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
