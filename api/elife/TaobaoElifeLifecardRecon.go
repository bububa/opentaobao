package elife

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/elife"
)

/* TaobaoElifeLifecardRecon
查询对账文件地址接口
taobao.elife.lifecard.recon

查询对账文件地址接口 */
func TaobaoElifeLifecardRecon(clt *core.SDKClient, req *elife.TaobaoElifeLifecardReconAPIRequest, session string) (*elife.TaobaoElifeLifecardReconAPIResponse, error) {
	var resp elife.TaobaoElifeLifecardReconAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
