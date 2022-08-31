package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeQueryHouseTagInfo 活动标查询接口
// alibaba.alihouse.existinghome.query.house.tag.info
//
// 活动标查询接口
func AlibabaAlihouseExistinghomeQueryHouseTagInfo(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeQueryHouseTagInfoAPIRequest, session string) (*alihouse.AlibabaAlihouseExistinghomeQueryHouseTagInfoAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseExistinghomeQueryHouseTagInfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
