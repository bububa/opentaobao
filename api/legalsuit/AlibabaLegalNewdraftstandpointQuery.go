package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// Alibabalegalnewdraftstandpointquery 未采纳口径查询(新)
// alibaba.legal.newdraftstandpoint.query
//
// 未采纳口径查询(新)
func Alibabalegalnewdraftstandpointquery(clt *core.SDKClient, req *legalsuit.AlibabalegalnewdraftstandpointqueryAPIRequest, session string) (*legalsuit.AlibabalegalnewdraftstandpointqueryAPIResponse, error) {
	var resp legalsuit.AlibabalegalnewdraftstandpointqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
