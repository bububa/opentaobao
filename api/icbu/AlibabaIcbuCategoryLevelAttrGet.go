package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// Alibabaicbucategorylevelattrget 层级属性的子属性获取
// alibaba.icbu.category.level.attr.get
//
// 用于获取层级属性（车型库）的子属性和属性值
func Alibabaicbucategorylevelattrget(clt *core.SDKClient, req *icbu.AlibabaicbucategorylevelattrgetAPIRequest, session string) (*icbu.AlibabaicbucategorylevelattrgetAPIResponse, error) {
	var resp icbu.AlibabaicbucategorylevelattrgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
