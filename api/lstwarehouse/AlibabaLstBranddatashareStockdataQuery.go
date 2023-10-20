package lstwarehouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstwarehouse"
)

// Alibabalstbranddatasharestockdataquery 查询品牌商品实仓库存/周转效能
// alibaba.lst.branddatashare.stockdata.query
//
// 品牌商查询授权供应商实仓库存数据
func Alibabalstbranddatasharestockdataquery(clt *core.SDKClient, req *lstwarehouse.AlibabalstbranddatasharestockdataqueryAPIRequest, session string) (*lstwarehouse.AlibabalstbranddatasharestockdataqueryAPIResponse, error) {
	var resp lstwarehouse.AlibabalstbranddatasharestockdataqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
