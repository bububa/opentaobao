package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

/* AlibabaDamaiMevOpenLockticket
大麦换验平台-第三方对外开放-票单接口lockTicket
alibaba.damai.mev.open.lockticket

开放接口 冻结票单 */
func AlibabaDamaiMevOpenLockticket(clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenLockticketAPIRequest, session string) (*damai.AlibabaDamaiMevOpenLockticketAPIResponse, error) {
	var resp damai.AlibabaDamaiMevOpenLockticketAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
