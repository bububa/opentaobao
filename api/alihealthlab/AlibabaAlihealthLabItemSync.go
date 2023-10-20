package alihealthlab

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthlab"
)

// Alibabaalihealthlabitemsync 阿里健康检验检测商品发布
// alibaba.alihealth.lab.item.sync
//
// iSV发布检验检测商品基本信息给健康，内部关联一个淘宝商品或SKU
func Alibabaalihealthlabitemsync(clt *core.SDKClient, req *alihealthlab.AlibabaalihealthlabitemsyncAPIRequest, session string) (*alihealthlab.AlibabaalihealthlabitemsyncAPIResponse, error) {
	var resp alihealthlab.AlibabaalihealthlabitemsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
