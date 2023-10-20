package alihealthmedical

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthmedical"
)

// Alibabaalihealthmedicalitempublish 三方入驻-开通服务
// alibaba.alihealth.medical.item.publish
//
// 三方入驻-开通服务
func Alibabaalihealthmedicalitempublish(clt *core.SDKClient, req *alihealthmedical.AlibabaalihealthmedicalitempublishAPIRequest, session string) (*alihealthmedical.AlibabaalihealthmedicalitempublishAPIResponse, error) {
	var resp alihealthmedical.AlibabaalihealthmedicalitempublishAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
