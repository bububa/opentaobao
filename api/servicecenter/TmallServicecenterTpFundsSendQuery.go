package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// Tmallservicecentertpfundssendquery 服务商资金权益发放的查询接口
// tmall.servicecenter.tp.funds.send.query
//
// 服务商资金权益发放结果的查询接口
func Tmallservicecentertpfundssendquery(clt *core.SDKClient, req *servicecenter.TmallservicecentertpfundssendqueryAPIRequest, session string) (*servicecenter.TmallservicecentertpfundssendqueryAPIResponse, error) {
	var resp servicecenter.TmallservicecentertpfundssendqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
