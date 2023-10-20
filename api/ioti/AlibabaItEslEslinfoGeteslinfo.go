package ioti

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ioti"
)

// AlibabaItEslEslinfoGeteslinfo 厂测查询价签当前信息
// alibaba.it.esl.eslinfo.geteslinfo
//
// 工厂对生产出的电子价签进行全流程功能测试，查询价签当前上报的信息
func AlibabaItEslEslinfoGeteslinfo(clt *core.SDKClient, req *ioti.AlibabaItEslEslinfoGeteslinfoAPIRequest, resp *ioti.AlibabaItEslEslinfoGeteslinfoAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
