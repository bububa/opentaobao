package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabadchainaoxiangchannelinventorybatchupload ERP全量同步销售库存数量
// alibaba.dchain.aoxiang.channel.inventory.batch.upload
//
// ERP全量同步销售库存数量
func Alibabadchainaoxiangchannelinventorybatchupload(clt *core.SDKClient, req *ascp.AlibabadchainaoxiangchannelinventorybatchuploadAPIRequest, session string) (*ascp.AlibabadchainaoxiangchannelinventorybatchuploadAPIResponse, error) {
	var resp ascp.AlibabadchainaoxiangchannelinventorybatchuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
