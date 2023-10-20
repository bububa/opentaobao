package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeTalentBindStore 达人号门店关系绑定
// alibaba.alihouse.newhome.talent.bind.store
//
// 达人号门店关系绑定
func AlibabaAlihouseNewhomeTalentBindStore(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeTalentBindStoreAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeTalentBindStoreAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeTalentBindStoreAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
