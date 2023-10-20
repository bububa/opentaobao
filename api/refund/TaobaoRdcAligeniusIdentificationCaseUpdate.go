package refund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/refund"
)

// Taobaordcaligeniusidentificationcaseupdate 鉴定工单信息同步
// taobao.rdc.aligenius.identification.case.update
//
// 同步商家鉴定工单信息
func Taobaordcaligeniusidentificationcaseupdate(clt *core.SDKClient, req *refund.TaobaordcaligeniusidentificationcaseupdateAPIRequest, session string) (*refund.TaobaordcaligeniusidentificationcaseupdateAPIResponse, error) {
	var resp refund.TaobaordcaligeniusidentificationcaseupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
