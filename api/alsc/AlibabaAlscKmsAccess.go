package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsckmsaccess 本地生活风控数据接入
// alibaba.alsc.kms.access
//
// 第三方使用本地生活数据对外提供服务，上报访问日志信息接口
func Alibabaalsckmsaccess(clt *core.SDKClient, req *alsc.AlibabaalsckmsaccessAPIRequest, session string) (*alsc.AlibabaalsckmsaccessAPIResponse, error) {
	var resp alsc.AlibabaalsckmsaccessAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
