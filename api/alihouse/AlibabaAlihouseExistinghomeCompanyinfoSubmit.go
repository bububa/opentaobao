package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeCompanyinfoSubmit 进件接口
// alibaba.alihouse.existinghome.companyinfo.submit
//
// 进件接口
func AlibabaAlihouseExistinghomeCompanyinfoSubmit(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeCompanyinfoSubmitAPIRequest, session string) (*alihouse.AlibabaAlihouseExistinghomeCompanyinfoSubmitAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseExistinghomeCompanyinfoSubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
