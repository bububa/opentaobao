package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// AlibabaAilabsIotBusinessRecipestepInsertorupdate 插入或更新食谱步骤
// alibaba.ailabs.iot.business.recipestep.insertorupdate
//
// 插入或更新食谱步骤
func AlibabaAilabsIotBusinessRecipestepInsertorupdate(clt *core.SDKClient, req *iot.AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIRequest, resp *iot.AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
