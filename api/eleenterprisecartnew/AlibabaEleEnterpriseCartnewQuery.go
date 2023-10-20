package eleenterprisecartnew

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterprisecartnew"
)

// Alibabaeleenterprisecartnewquery 新版购物车查询
// alibaba.ele.enterprise.cartnew.query
//
// 新版购物车查询
func Alibabaeleenterprisecartnewquery(clt *core.SDKClient, req *eleenterprisecartnew.AlibabaeleenterprisecartnewqueryAPIRequest, session string) (*eleenterprisecartnew.AlibabaeleenterprisecartnewqueryAPIResponse, error) {
	var resp eleenterprisecartnew.AlibabaeleenterprisecartnewqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
