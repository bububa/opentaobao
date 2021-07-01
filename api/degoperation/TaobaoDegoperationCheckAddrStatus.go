package degoperation

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/degoperation"
)

/* TaobaoDegoperationCheckAddrStatus
地址
taobao.degoperation.check.addr.status

激励 */
func TaobaoDegoperationCheckAddrStatus(clt *core.SDKClient, req *degoperation.TaobaoDegoperationCheckAddrStatusAPIRequest, session string) (*degoperation.TaobaoDegoperationCheckAddrStatusAPIResponse, error) {
	var resp degoperation.TaobaoDegoperationCheckAddrStatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
