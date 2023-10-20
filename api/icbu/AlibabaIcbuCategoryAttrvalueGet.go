package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// Alibabaicbucategoryattrvalueget 属性值获取
// alibaba.icbu.category.attrvalue.get
//
// 属性值获取
func Alibabaicbucategoryattrvalueget(clt *core.SDKClient, req *icbu.AlibabaicbucategoryattrvaluegetAPIRequest, session string) (*icbu.AlibabaicbucategoryattrvaluegetAPIResponse, error) {
	var resp icbu.AlibabaicbucategoryattrvaluegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
