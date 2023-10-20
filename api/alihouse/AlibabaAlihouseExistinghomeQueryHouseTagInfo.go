package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeQueryHouseTagInfo 活动标查询接口
// alibaba.alihouse.existinghome.query.house.tag.info
//
// 活动标查询接口
func AlibabaAlihouseExistinghomeQueryHouseTagInfo(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeQueryHouseTagInfoAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeQueryHouseTagInfoAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
