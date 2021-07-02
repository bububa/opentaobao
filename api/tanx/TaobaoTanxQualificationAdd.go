package tanx

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tanx"
)

// TaobaoTanxQualificationAdd 提交资质接口
// taobao.tanx.qualification.add
//
// dsp客户提交客户资质和行业资质
func TaobaoTanxQualificationAdd(clt *core.SDKClient, req *tanx.TaobaoTanxQualificationAddAPIRequest, session string) (*tanx.TaobaoTanxQualificationAddAPIResponse, error) {
	var resp tanx.TaobaoTanxQualificationAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
