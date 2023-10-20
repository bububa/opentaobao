package alihealthlab

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthlab"
)

// Alibabaalihealthlabitemtbitemskurelationsync 阿里健康检验检测业务，检验检测项目淘宝商品SKU关系同步
// alibaba.alihealth.lab.item.tbitemsku.relation.sync
//
// 阿里健康检验检测业务，检验检测项目淘宝商品SKU关系同步
func Alibabaalihealthlabitemtbitemskurelationsync(clt *core.SDKClient, req *alihealthlab.AlibabaalihealthlabitemtbitemskurelationsyncAPIRequest, session string) (*alihealthlab.AlibabaalihealthlabitemtbitemskurelationsyncAPIResponse, error) {
	var resp alihealthlab.AlibabaalihealthlabitemtbitemskurelationsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
