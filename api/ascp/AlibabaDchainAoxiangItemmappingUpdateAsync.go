package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangItemmappingUpdateAsync 创建/更新商货品关联关系
// alibaba.dchain.aoxiang.itemmapping.update.async
//
// 创建/更新商货品关联关系
func AlibabaDchainAoxiangItemmappingUpdateAsync(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangItemmappingUpdateAsyncAPIRequest, session string) (*ascp.AlibabaDchainAoxiangItemmappingUpdateAsyncAPIResponse, error) {
	var resp ascp.AlibabaDchainAoxiangItemmappingUpdateAsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
