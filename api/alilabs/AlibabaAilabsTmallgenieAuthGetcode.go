package alilabs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

/* AlibabaAilabsTmallgenieAuthGetcode
获取token
alibaba.ailabs.tmallgenie.auth.getcode

获取天猫精灵authCode */
func AlibabaAilabsTmallgenieAuthGetcode(clt *core.SDKClient, req *alilabs.AlibabaAilabsTmallgenieAuthGetcodeAPIRequest, session string) (*alilabs.AlibabaAilabsTmallgenieAuthGetcodeAPIResponse, error) {
	var resp alilabs.AlibabaAilabsTmallgenieAuthGetcodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
