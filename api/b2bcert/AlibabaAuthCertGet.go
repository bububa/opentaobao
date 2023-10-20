package b2bcert

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/b2bcert"
)

// Alibabaauthcertget 获取证书数据
// alibaba.auth.cert.get
//
// 获取证书数据
func Alibabaauthcertget(clt *core.SDKClient, req *b2bcert.AlibabaauthcertgetAPIRequest, session string) (*b2bcert.AlibabaauthcertgetAPIResponse, error) {
	var resp b2bcert.AlibabaauthcertgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
