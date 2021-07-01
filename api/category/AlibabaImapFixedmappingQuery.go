package category

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/category"
)

/* AlibabaImapFixedmappingQuery
查询两个渠道之间的固定映射关系，不通过算法兜底
alibaba.imap.fixedmapping.query

查询两个渠道之间的固定映射关系，不通过算法兜底 */
func AlibabaImapFixedmappingQuery(clt *core.SDKClient, req *category.AlibabaImapFixedmappingQueryAPIRequest, session string) (*category.AlibabaImapFixedmappingQueryAPIResponse, error) {
	var resp category.AlibabaImapFixedmappingQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
