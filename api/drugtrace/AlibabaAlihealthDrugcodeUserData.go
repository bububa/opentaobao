package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugcodeUserData 西安杨森同步用户行为接口
// alibaba.alihealth.drugcode.user.data
//
// 西安杨森同步用户行为接口
func AlibabaAlihealthDrugcodeUserData(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugcodeUserDataAPIRequest, resp *drugtrace.AlibabaAlihealthDrugcodeUserDataAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
