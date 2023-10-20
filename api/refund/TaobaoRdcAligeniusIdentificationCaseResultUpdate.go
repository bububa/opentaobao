package refund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/refund"
)

// Taobaordcaligeniusidentificationcaseresultupdate 鉴定工单结果同步
// taobao.rdc.aligenius.identification.case.result.update
//
// 同步鉴定工单结果信息
func Taobaordcaligeniusidentificationcaseresultupdate(clt *core.SDKClient, req *refund.TaobaordcaligeniusidentificationcaseresultupdateAPIRequest, session string) (*refund.TaobaordcaligeniusidentificationcaseresultupdateAPIResponse, error) {
	var resp refund.TaobaordcaligeniusidentificationcaseresultupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
