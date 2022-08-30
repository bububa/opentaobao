package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeQueryHouseBaseInfo 查询房源基本信息
// alibaba.alihouse.existinghome.query.house.base.info
//
// 查询房源基本信息
func AlibabaAlihouseExistinghomeQueryHouseBaseInfo(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeQueryHouseBaseInfoAPIRequest, session string) (*alihouse.AlibabaAlihouseExistinghomeQueryHouseBaseInfoAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseExistinghomeQueryHouseBaseInfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
