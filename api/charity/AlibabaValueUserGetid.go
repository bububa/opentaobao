package charity

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// Alibabavalueusergetid 获取用户userId
// alibaba.value.user.getid
//
// 获取用户userId
func Alibabavalueusergetid(clt *core.SDKClient, req *charity.AlibabavalueusergetidAPIRequest, session string) (*charity.AlibabavalueusergetidAPIResponse, error) {
	var resp charity.AlibabavalueusergetidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
