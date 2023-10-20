package alihealthlab

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthlab"
)

// Alibabaalihealthlabstoresync 阿里健康检验检测业务，isv门店同步到健康
// alibaba.alihealth.lab.store.sync
//
// 阿里健康检验检测业务，isv门店同步到健康。支持门店的上线、下线操作
func Alibabaalihealthlabstoresync(clt *core.SDKClient, req *alihealthlab.AlibabaalihealthlabstoresyncAPIRequest, session string) (*alihealthlab.AlibabaalihealthlabstoresyncAPIResponse, error) {
	var resp alihealthlab.AlibabaalihealthlabstoresyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
