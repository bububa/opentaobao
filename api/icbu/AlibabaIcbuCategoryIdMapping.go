package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// Alibabaicbucategoryidmapping 新旧属性的映射
// alibaba.icbu.category.id.mapping
//
// 商品发布接口升级，需要传入新的类目。这个接口 根据旧的类目id，获取新的类目id
func Alibabaicbucategoryidmapping(clt *core.SDKClient, req *icbu.AlibabaicbucategoryidmappingAPIRequest, session string) (*icbu.AlibabaicbucategoryidmappingAPIResponse, error) {
	var resp icbu.AlibabaicbucategoryidmappingAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
