package tanx

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tanx"
)

// TaobaoTanxQualificationFind 资质查询接口
// taobao.tanx.qualification.find
//
// 资质查询接口
func TaobaoTanxQualificationFind(clt *core.SDKClient, req *tanx.TaobaoTanxQualificationFindAPIRequest, session string) (*tanx.TaobaoTanxQualificationFindAPIResponse, error) {
	var resp tanx.TaobaoTanxQualificationFindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
