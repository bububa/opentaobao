package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// AlibabaIcbuCategoryAttrvalueGet 属性值获取
// alibaba.icbu.category.attrvalue.get
//
// 属性值获取
func AlibabaIcbuCategoryAttrvalueGet(clt *core.SDKClient, req *icbu.AlibabaIcbuCategoryAttrvalueGetAPIRequest, resp *icbu.AlibabaIcbuCategoryAttrvalueGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
