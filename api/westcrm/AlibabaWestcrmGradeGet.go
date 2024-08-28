package westcrm

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/westcrm"
)

// AlibabaWestcrmGradeGet 获取等级列表
// alibaba.westcrm.grade.get
//
// 获取会员卡等级列表
func AlibabaWestcrmGradeGet(ctx context.Context, clt *core.SDKClient, req *westcrm.AlibabaWestcrmGradeGetAPIRequest, resp *westcrm.AlibabaWestcrmGradeGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
