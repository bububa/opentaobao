package brandhub

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/brandhub"
)

// Taobaobrandstarshoprpttargetget 明星店铺定向维度报表
// taobao.brand.starshop.rpt.target.get
//
// 获取明星店铺定向维度分日报表数据，只能查询近90天内的数据，包括展现量，点击量等
func Taobaobrandstarshoprpttargetget(clt *core.SDKClient, req *brandhub.TaobaobrandstarshoprpttargetgetAPIRequest, session string) (*brandhub.TaobaobrandstarshoprpttargetgetAPIResponse, error) {
	var resp brandhub.TaobaobrandstarshoprpttargetgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
