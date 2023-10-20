package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytScqyPutpackagebind 码拼箱建立父子关系接口
// alibaba.alihealth.drug.kyt.scqy.putpackagebind
//
// 码拼箱建立父子关系接口
func AlibabaAlihealthDrugKytScqyPutpackagebind(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytScqyPutpackagebindAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytScqyPutpackagebindAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytScqyPutpackagebindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
