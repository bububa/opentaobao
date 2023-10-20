package normalvisa

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/normalvisa"
)

// Taobaoalitriptravelnormalvisaupdatepersonstauts 更新签证办理进度
// taobao.alitrip.travel.normalvisa.updatepersonstauts
//
// 更新签证办理进度
func Taobaoalitriptravelnormalvisaupdatepersonstauts(clt *core.SDKClient, req *normalvisa.TaobaoalitriptravelnormalvisaupdatepersonstautsAPIRequest, session string) (*normalvisa.TaobaoalitriptravelnormalvisaupdatepersonstautsAPIResponse, error) {
	var resp normalvisa.TaobaoalitriptravelnormalvisaupdatepersonstautsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
