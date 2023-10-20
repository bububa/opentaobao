package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmcardsearchcard 搜索卡实例列表(支持号段查询)
// alibaba.alsc.crm.card.searchcard
//
// 搜索卡实例列表(支持号段查询)
func Alibabaalsccrmcardsearchcard(clt *core.SDKClient, req *alsc.AlibabaalsccrmcardsearchcardAPIRequest, session string) (*alsc.AlibabaalsccrmcardsearchcardAPIResponse, error) {
	var resp alsc.AlibabaalsccrmcardsearchcardAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
