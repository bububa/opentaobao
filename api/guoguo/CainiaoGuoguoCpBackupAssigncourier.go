package guoguo

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/guoguo"
)

// CainiaoGuoguoCpBackupAssigncourier CP兜底后指定接单的小件员
// cainiao.guoguo.cp.backup.assigncourier
//
// CP兜底后指定接单的小件员；CP改派小件员
func CainiaoGuoguoCpBackupAssigncourier(clt *core.SDKClient, req *guoguo.CainiaoGuoguoCpBackupAssigncourierAPIRequest, resp *guoguo.CainiaoGuoguoCpBackupAssigncourierAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
