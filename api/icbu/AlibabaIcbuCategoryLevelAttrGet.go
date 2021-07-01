package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

/* AlibabaIcbuCategoryLevelAttrGet
层级属性的子属性获取
alibaba.icbu.category.level.attr.get

用于获取层级属性（车型库）的子属性和属性值 */
func AlibabaIcbuCategoryLevelAttrGet(clt *core.SDKClient, req *icbu.AlibabaIcbuCategoryLevelAttrGetAPIRequest, session string) (*icbu.AlibabaIcbuCategoryLevelAttrGetAPIResponse, error) {
	var resp icbu.AlibabaIcbuCategoryLevelAttrGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
