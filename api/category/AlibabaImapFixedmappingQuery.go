package category

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/category"
)

// AlibabaImapFixedmappingQuery 查询两个渠道之间的固定映射关系，不通过算法兜底
// alibaba.imap.fixedmapping.query
//
// 查询两个渠道之间的固定映射关系，不通过算法兜底
func AlibabaImapFixedmappingQuery(clt *core.SDKClient, req *category.AlibabaImapFixedmappingQueryAPIRequest, resp *category.AlibabaImapFixedmappingQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
