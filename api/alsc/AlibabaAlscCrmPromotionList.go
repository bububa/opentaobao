package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmpromotionlist 获取促销规则列表
// alibaba.alsc.crm.promotion.list
//
// 获取品牌的促销规则列表
func Alibabaalsccrmpromotionlist(clt *core.SDKClient, req *alsc.AlibabaalsccrmpromotionlistAPIRequest, session string) (*alsc.AlibabaalsccrmpromotionlistAPIResponse, error) {
	var resp alsc.AlibabaalsccrmpromotionlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
