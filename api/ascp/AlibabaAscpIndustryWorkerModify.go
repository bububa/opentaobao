package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabaascpindustryworkermodify 送货入户并安装修改师傅信息
// alibaba.ascp.industry.worker.modify
//
// 送货入户并安装修改师傅信息
func Alibabaascpindustryworkermodify(clt *core.SDKClient, req *ascp.AlibabaascpindustryworkermodifyAPIRequest, session string) (*ascp.AlibabaascpindustryworkermodifyAPIResponse, error) {
	var resp ascp.AlibabaascpindustryworkermodifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
