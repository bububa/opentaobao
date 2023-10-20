package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmopenpointoperate 积分操作接口
// alibaba.alsc.crm.open.point.operate
//
// 同步积分接口
func Alibabaalsccrmopenpointoperate(clt *core.SDKClient, req *alsc.AlibabaalsccrmopenpointoperateAPIRequest, session string) (*alsc.AlibabaalsccrmopenpointoperateAPIResponse, error) {
	var resp alsc.AlibabaalsccrmopenpointoperateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
