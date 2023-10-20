package yunos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/yunos"
)

// Aliyuncosmogatewayinvoke alios cosmo服务调用
// aliyun.cosmo.gateway.invoke
//
// AliOS cosmo服务分发平台对外调用接口
func Aliyuncosmogatewayinvoke(clt *core.SDKClient, req *yunos.AliyuncosmogatewayinvokeAPIRequest, session string) (*yunos.AliyuncosmogatewayinvokeAPIResponse, error) {
	var resp yunos.AliyuncosmogatewayinvokeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
