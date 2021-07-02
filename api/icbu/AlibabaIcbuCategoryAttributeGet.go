package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// AlibabaIcbuCategoryAttributeGet 类目属性获取
// alibaba.icbu.category.attribute.get
//
// 根据类目ID获取系统定义的属性
func AlibabaIcbuCategoryAttributeGet(clt *core.SDKClient, req *icbu.AlibabaIcbuCategoryAttributeGetAPIRequest, session string) (*icbu.AlibabaIcbuCategoryAttributeGetAPIResponse, error) {
	var resp icbu.AlibabaIcbuCategoryAttributeGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
