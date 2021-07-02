package alihealthoutflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// AlibabaAlihealthOutflowDrugSaveorupdate 处方外流-药品同步接口
// alibaba.alihealth.outflow.drug.saveorupdate
//
// 处方外流-药品同步接口
func AlibabaAlihealthOutflowDrugSaveorupdate(clt *core.SDKClient, req *alihealthoutflow.AlibabaAlihealthOutflowDrugSaveorupdateAPIRequest, session string) (*alihealthoutflow.AlibabaAlihealthOutflowDrugSaveorupdateAPIResponse, error) {
	var resp alihealthoutflow.AlibabaAlihealthOutflowDrugSaveorupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
