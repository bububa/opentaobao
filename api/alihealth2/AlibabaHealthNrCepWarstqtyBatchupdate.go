package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaHealthNrCepWarstqtyBatchupdate 批量更新ISV库存
// alibaba.health.nr.cep.warstqty.batchupdate
//
// 青岛医保服务-ISV批量更新孔雀翎中库存数据
func AlibabaHealthNrCepWarstqtyBatchupdate(clt *core.SDKClient, req *alihealth2.AlibabaHealthNrCepWarstqtyBatchupdateAPIRequest, resp *alihealth2.AlibabaHealthNrCepWarstqtyBatchupdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
