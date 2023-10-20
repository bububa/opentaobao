package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// CainiaoCbossWorkplatformBiztypeQuerybyid 菜鸟工单平台根据业务类型id查询业务类型详细信息
// cainiao.cboss.workplatform.biztype.querybyid
//
// 菜鸟工单平台根据业务类型id查询业务类型详细信息。 目前调用者ISV
func CainiaoCbossWorkplatformBiztypeQuerybyid(clt *core.SDKClient, req *logistic.CainiaoCbossWorkplatformBiztypeQuerybyidAPIRequest, session string) (*logistic.CainiaoCbossWorkplatformBiztypeQuerybyidAPIResponse, error) {
	var resp logistic.CainiaoCbossWorkplatformBiztypeQuerybyidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
