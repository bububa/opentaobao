package ott

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ott"
)

// YoukuOttAlicbFacadeserviceGetdata 影视SDK获取设备能力值
// youku.ott.alicb.facadeservice.getdata
//
// 影视SDK获取设备能力值
func YoukuOttAlicbFacadeserviceGetdata(clt *core.SDKClient, req *ott.YoukuOttAlicbFacadeserviceGetdataAPIRequest, resp *ott.YoukuOttAlicbFacadeserviceGetdataAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
