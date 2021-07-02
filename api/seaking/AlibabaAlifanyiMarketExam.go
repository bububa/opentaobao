package seaking

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/seaking"
)

// AlibabaAlifanyiMarketExam 通过考试用户
// alibaba.alifanyi.market.exam
//
// 企业或组织购买软件服务后可参与阿里翻译在线系统的考试认证，接口返回该企业或组织认证通过的用户
func AlibabaAlifanyiMarketExam(clt *core.SDKClient, req *seaking.AlibabaAlifanyiMarketExamAPIRequest, session string) (*seaking.AlibabaAlifanyiMarketExamAPIResponse, error) {
	var resp seaking.AlibabaAlifanyiMarketExamAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
