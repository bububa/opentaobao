package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// Alibabaicbucategoryattributeget 类目属性获取
// alibaba.icbu.category.attribute.get
//
// 根据类目ID获取系统定义的属性
func Alibabaicbucategoryattributeget(clt *core.SDKClient, req *icbu.AlibabaicbucategoryattributegetAPIRequest, session string) (*icbu.AlibabaicbucategoryattributegetAPIResponse, error) {
	var resp icbu.AlibabaicbucategoryattributegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
