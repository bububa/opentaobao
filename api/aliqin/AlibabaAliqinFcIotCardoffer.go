package aliqin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliqin"
)

/* AlibabaAliqinFcIotCardoffer
查询物联网卡上订购的offer
alibaba.aliqin.fc.iot.cardoffer

查询物联网卡上订购的offer */
func AlibabaAliqinFcIotCardoffer(clt *core.SDKClient, req *aliqin.AlibabaAliqinFcIotCardofferAPIRequest, session string) (*aliqin.AlibabaAliqinFcIotCardofferAPIResponse, error) {
	var resp aliqin.AlibabaAliqinFcIotCardofferAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
