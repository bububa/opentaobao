package seaking

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/seaking"
)

// Alibabaseakingservicepack 获取海王用户权限包
// alibaba.seaking.servicepack
//
// 获取海王用户权限包
func Alibabaseakingservicepack(clt *core.SDKClient, req *seaking.AlibabaseakingservicepackAPIRequest, session string) (*seaking.AlibabaseakingservicepackAPIResponse, error) {
	var resp seaking.AlibabaseakingservicepackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
