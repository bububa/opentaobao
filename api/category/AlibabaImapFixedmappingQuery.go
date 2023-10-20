package category

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/category"
)

// Alibabaimapfixedmappingquery 查询两个渠道之间的固定映射关系，不通过算法兜底
// alibaba.imap.fixedmapping.query
//
// 查询两个渠道之间的固定映射关系，不通过算法兜底
func Alibabaimapfixedmappingquery(clt *core.SDKClient, req *category.AlibabaimapfixedmappingqueryAPIRequest, session string) (*category.AlibabaimapfixedmappingqueryAPIResponse, error) {
	var resp category.AlibabaimapfixedmappingqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
