package tanx

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tanx"
)

// TaobaoTanxQualificationFind 资质查询接口
// taobao.tanx.qualification.find
//
// 资质查询接口
func TaobaoTanxQualificationFind(clt *core.SDKClient, req *tanx.TaobaoTanxQualificationFindAPIRequest, resp *tanx.TaobaoTanxQualificationFindAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
