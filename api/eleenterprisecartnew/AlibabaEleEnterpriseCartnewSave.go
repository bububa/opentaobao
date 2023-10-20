package eleenterprisecartnew

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterprisecartnew"
)

// Alibabaeleenterprisecartnewsave 新版创建购物车
// alibaba.ele.enterprise.cartnew.save
//
// 新版创建购物车
func Alibabaeleenterprisecartnewsave(clt *core.SDKClient, req *eleenterprisecartnew.AlibabaeleenterprisecartnewsaveAPIRequest, session string) (*eleenterprisecartnew.AlibabaeleenterprisecartnewsaveAPIResponse, error) {
	var resp eleenterprisecartnew.AlibabaeleenterprisecartnewsaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
