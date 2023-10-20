package ioti

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ioti"
)

// Alibabaitesleslinfogeteslinfo 厂测查询价签当前信息
// alibaba.it.esl.eslinfo.geteslinfo
//
// 工厂对生产出的电子价签进行全流程功能测试，查询价签当前上报的信息
func Alibabaitesleslinfogeteslinfo(clt *core.SDKClient, req *ioti.AlibabaitesleslinfogeteslinfoAPIRequest, session string) (*ioti.AlibabaitesleslinfogeteslinfoAPIResponse, error) {
	var resp ioti.AlibabaitesleslinfogeteslinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
