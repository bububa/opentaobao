package normalvisa

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/normalvisa"
)

// Alitriptravelvisaapplicantimport 签证申请人导入
// alitrip.travel.visa.applicant.import
//
// 签证线下申请人导入接口。供商家将线下的签证申请人信息导入，进行签证线上化办理
func Alitriptravelvisaapplicantimport(clt *core.SDKClient, req *normalvisa.AlitriptravelvisaapplicantimportAPIRequest, session string) (*normalvisa.AlitriptravelvisaapplicantimportAPIResponse, error) {
	var resp normalvisa.AlitriptravelvisaapplicantimportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
