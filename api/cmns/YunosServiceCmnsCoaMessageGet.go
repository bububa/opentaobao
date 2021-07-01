package cmns

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cmns"
)

/* YunosServiceCmnsCoaMessageGet
消息详情查询
yunos.service.cmns.coa.message.get

第三方应用开发者调用此接口查询消息详情，只能查询此appKey发的消息 */
func YunosServiceCmnsCoaMessageGet(clt *core.SDKClient, req *cmns.YunosServiceCmnsCoaMessageGetAPIRequest, session string) (*cmns.YunosServiceCmnsCoaMessageGetAPIResponse, error) {
	var resp cmns.YunosServiceCmnsCoaMessageGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
