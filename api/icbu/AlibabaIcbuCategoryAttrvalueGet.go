package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// AlibabaIcbuCategoryAttrvalueGet 属性值获取
// alibaba.icbu.category.attrvalue.get
//
// 属性值获取
func AlibabaIcbuCategoryAttrvalueGet(clt *core.SDKClient, req *icbu.AlibabaIcbuCategoryAttrvalueGetAPIRequest, session string) (*icbu.AlibabaIcbuCategoryAttrvalueGetAPIResponse, error) {
	var resp icbu.AlibabaIcbuCategoryAttrvalueGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
