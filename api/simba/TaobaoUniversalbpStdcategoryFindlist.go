package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpStdcategoryFindlist 人群相关类目查询
// taobao.universalbp.stdcategory.findlist
//
// 入参商品信息，出参商品所属类别
func TaobaoUniversalbpStdcategoryFindlist(clt *core.SDKClient, req *simba.TaobaoUniversalbpStdcategoryFindlistAPIRequest, session string) (*simba.TaobaoUniversalbpStdcategoryFindlistAPIResponse, error) {
	var resp simba.TaobaoUniversalbpStdcategoryFindlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
