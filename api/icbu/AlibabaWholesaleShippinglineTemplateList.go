package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// Alibabawholesaleshippinglinetemplatelist 获取运费模板
// alibaba.wholesale.shippingline.template.list
//
// 查询运费模板信息
func Alibabawholesaleshippinglinetemplatelist(clt *core.SDKClient, req *icbu.AlibabawholesaleshippinglinetemplatelistAPIRequest, session string) (*icbu.AlibabawholesaleshippinglinetemplatelistAPIResponse, error) {
	var resp icbu.AlibabawholesaleshippinglinetemplatelistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
