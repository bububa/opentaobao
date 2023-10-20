package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// AlibabaIcbuCategoryIdMapping 新旧属性的映射
// alibaba.icbu.category.id.mapping
//
// 商品发布接口升级，需要传入新的类目。这个接口 根据旧的类目id，获取新的类目id
func AlibabaIcbuCategoryIdMapping(clt *core.SDKClient, req *icbu.AlibabaIcbuCategoryIdMappingAPIRequest, resp *icbu.AlibabaIcbuCategoryIdMappingAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
