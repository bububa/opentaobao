package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// Alibabadchainmiaoshifucustomercomplaintsput 服务商工人客诉数据上传
// alibaba.dchain.miaoshifu.customer.complaints.put
//
// 数字服务供应链平台提供给服务商上传工人客诉数据
func Alibabadchainmiaoshifucustomercomplaintsput(clt *core.SDKClient, req *tmallsc.AlibabadchainmiaoshifucustomercomplaintsputAPIRequest, session string) (*tmallsc.AlibabadchainmiaoshifucustomercomplaintsputAPIResponse, error) {
	var resp tmallsc.AlibabadchainmiaoshifucustomercomplaintsputAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
