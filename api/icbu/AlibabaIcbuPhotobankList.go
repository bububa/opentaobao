package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// AlibabaIcbuPhotobankList 国际站图片银行查询接口
// alibaba.icbu.photobank.list
//
// 国际站图片银行查询接口
func AlibabaIcbuPhotobankList(clt *core.SDKClient, req *icbu.AlibabaIcbuPhotobankListAPIRequest, resp *icbu.AlibabaIcbuPhotobankListAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
