package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabahealthnrcepwarstqtybatchupdate 批量更新ISV库存
// alibaba.health.nr.cep.warstqty.batchupdate
//
// 青岛医保服务-ISV批量更新孔雀翎中库存数据
func Alibabahealthnrcepwarstqtybatchupdate(clt *core.SDKClient, req *alihealth2.AlibabahealthnrcepwarstqtybatchupdateAPIRequest, session string) (*alihealth2.AlibabahealthnrcepwarstqtybatchupdateAPIResponse, error) {
	var resp alihealth2.AlibabahealthnrcepwarstqtybatchupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
