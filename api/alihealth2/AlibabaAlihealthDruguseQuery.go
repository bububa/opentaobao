package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthDruguseQuery 合理用药规则查询
// alibaba.alihealth.druguse.query
//
// 查询用户购买的药品命中的风险规则
func AlibabaAlihealthDruguseQuery(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthDruguseQueryAPIRequest, session string) (*alihealth2.AlibabaAlihealthDruguseQueryAPIResponse, error) {
	var resp alihealth2.AlibabaAlihealthDruguseQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
