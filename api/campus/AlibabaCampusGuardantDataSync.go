package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusguardantdatasync 刷卡数据同步
// alibaba.campus.guardant.data.sync
//
// 数据同步门禁系统
func Alibabacampusguardantdatasync(clt *core.SDKClient, req *campus.AlibabacampusguardantdatasyncAPIRequest, session string) (*campus.AlibabacampusguardantdatasyncAPIResponse, error) {
	var resp campus.AlibabacampusguardantdatasyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
