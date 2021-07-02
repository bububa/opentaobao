package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// AlibabaLsyCrmUpdate 跟进客资状态接口
// alibaba.lsy.crm.update
//
// 同步客资状态接口
func AlibabaLsyCrmUpdate(clt *core.SDKClient, req *user.AlibabaLsyCrmUpdateAPIRequest, session string) (*user.AlibabaLsyCrmUpdateAPIResponse, error) {
	var resp user.AlibabaLsyCrmUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
