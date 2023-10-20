package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugwxinfoupload 小程序数据回传
// alibaba.alihealth.drug.wxinfo.upload
//
// 小程序数据回传
func Alibabaalihealthdrugwxinfoupload(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugwxinfouploadAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugwxinfouploadAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugwxinfouploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
