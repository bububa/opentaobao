package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusguarddatasync 卡巴数据同步
// alibaba.campus.guard.data.sync
//
// 数据同步门禁系统
func Alibabacampusguarddatasync(clt *core.SDKClient, req *campus.AlibabacampusguarddatasyncAPIRequest, session string) (*campus.AlibabacampusguarddatasyncAPIResponse, error) {
	var resp campus.AlibabacampusguarddatasyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
