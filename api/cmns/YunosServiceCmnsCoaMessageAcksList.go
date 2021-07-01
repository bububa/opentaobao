package cmns

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cmns"
)

/* YunosServiceCmnsCoaMessageAcksList
消息ack记录查询
yunos.service.cmns.coa.message.acks.list

第三方应用开发者调用此接口查询消息ack记录 */
func YunosServiceCmnsCoaMessageAcksList(clt *core.SDKClient, req *cmns.YunosServiceCmnsCoaMessageAcksListAPIRequest, session string) (*cmns.YunosServiceCmnsCoaMessageAcksListAPIResponse, error) {
	var resp cmns.YunosServiceCmnsCoaMessageAcksListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
