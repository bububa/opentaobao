package westcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/westcrm"
)

// AlibabaWestcrmGradeGet 获取等级列表
// alibaba.westcrm.grade.get
//
// 获取会员卡等级列表
func AlibabaWestcrmGradeGet(clt *core.SDKClient, req *westcrm.AlibabaWestcrmGradeGetAPIRequest, resp *westcrm.AlibabaWestcrmGradeGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
