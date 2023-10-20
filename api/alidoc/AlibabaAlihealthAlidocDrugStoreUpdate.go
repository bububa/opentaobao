package alidoc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alidoc"
)

// Alibabaalihealthalidocdrugstoreupdate 更新药店
// alibaba.alihealth.alidoc.drug.store.update
//
// 药店信息更新接口
func Alibabaalihealthalidocdrugstoreupdate(clt *core.SDKClient, req *alidoc.AlibabaalihealthalidocdrugstoreupdateAPIRequest, session string) (*alidoc.AlibabaalihealthalidocdrugstoreupdateAPIResponse, error) {
	var resp alidoc.AlibabaalihealthalidocdrugstoreupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
