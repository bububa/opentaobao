package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugdownloadgetentauthent 获取授权企业列表
// alibaba.alihealth.drug.download.getentauthent
//
// D2D数据落地获取授权企业列表
func Alibabaalihealthdrugdownloadgetentauthent(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugdownloadgetentauthentAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugdownloadgetentauthentAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugdownloadgetentauthentAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
