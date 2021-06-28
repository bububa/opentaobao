package cmns

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/cmns"
)

/* 
CMNS消息发送到达查询 
yunos.service.cmns.coa.messageresult.get

CMNS消息发送到达查询,根据消息ID查询，仅能查询该appKey所发送的消息
*/
func YunosServiceCmnsCoaMessageresultGet(clt *core.SDKClient, req *cmns.YunosServiceCmnsCoaMessageresultGetRequest, session string) (*cmns.YunosServiceCmnsCoaMessageresultGetAPIResponse, error) {
    var resp cmns.YunosServiceCmnsCoaMessageresultGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
