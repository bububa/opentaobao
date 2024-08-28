package ott

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ott"
)

// YoukuOttAlicbFacadeserviceGetdata 影视SDK获取设备能力值
// youku.ott.alicb.facadeservice.getdata
//
// 影视SDK获取设备能力值
func YoukuOttAlicbFacadeserviceGetdata(ctx context.Context, clt *core.SDKClient, req *ott.YoukuOttAlicbFacadeserviceGetdataAPIRequest, resp *ott.YoukuOttAlicbFacadeserviceGetdataAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
