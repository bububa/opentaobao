package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaouniversalbpstdcategoryfindcategorycondition 获取类目过滤条件
// taobao.universalbp.stdcategory.findcategorycondition
//
// 查询全量类目信息列表，用于进行类目兴趣人群相关定向
func Taobaouniversalbpstdcategoryfindcategorycondition(clt *core.SDKClient, req *simba.TaobaouniversalbpstdcategoryfindcategoryconditionAPIRequest, session string) (*simba.TaobaouniversalbpstdcategoryfindcategoryconditionAPIResponse, error) {
	var resp simba.TaobaouniversalbpstdcategoryfindcategoryconditionAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
