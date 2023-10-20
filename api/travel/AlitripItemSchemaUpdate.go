package travel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/travel"
)

// Alitripitemschemaupdate 使用schema进行商品编辑
// alitrip.item.schema.update
//
// 飞猪度假商品使用schema进行商品编辑。目前支持类目：出境自由行(50278002)、境内自由行(50272002)、出境跟团游(50258005)、境内跟团游(50258004)、境外一日游/多日游(50276003)
func Alitripitemschemaupdate(clt *core.SDKClient, req *travel.AlitripitemschemaupdateAPIRequest, session string) (*travel.AlitripitemschemaupdateAPIResponse, error) {
	var resp travel.AlitripitemschemaupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
