package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// Alibabadatabankopenoneservicedataready 瓴羊DaaS消费者增长CGP查询DataReady
// alibaba.databank.open.oneservice.dataready
//
// 瓴羊DaaS消费者增长CGP取数接口
func Alibabadatabankopenoneservicedataready(clt *core.SDKClient, req *user.AlibabadatabankopenoneservicedatareadyAPIRequest, session string) (*user.AlibabadatabankopenoneservicedatareadyAPIResponse, error) {
	var resp user.AlibabadatabankopenoneservicedatareadyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
