package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// Tmallworkcardidentify 工单核销
// tmall.workcard.identify
//
// 工单核销，当工单完成以后，通过调用此接口核销
// 可以按照多维度核销工单，
// 电器预约安装按照工单维度核销，必选字段workcard_id,buyer_id,identify_code，可选字段attrs，通过扩展字段attrs回传机器码，格式{sn:&#39;机器码&#39;}
func Tmallworkcardidentify(clt *core.SDKClient, req *tmallsc.TmallworkcardidentifyAPIRequest, session string) (*tmallsc.TmallworkcardidentifyAPIResponse, error) {
	var resp tmallsc.TmallworkcardidentifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
