package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// Alibabadatabankopenoneservicegetdata 瓴羊DaaS消费者运营CGP取数接口
// alibaba.databank.open.oneservice.getdata
//
// 瓴羊DaaS消费者运营CGP取数接口
func Alibabadatabankopenoneservicegetdata(clt *core.SDKClient, req *user.AlibabadatabankopenoneservicegetdataAPIRequest, session string) (*user.AlibabadatabankopenoneservicegetdataAPIResponse, error) {
	var resp user.AlibabadatabankopenoneservicegetdataAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
