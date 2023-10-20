package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghouseHouseBase （租房）同步房屋信息
// alibaba.alihouse.existinghouse.house.base
//
// 房屋信息上翻
func AlibabaAlihouseExistinghouseHouseBase(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghouseHouseBaseAPIRequest, session string) (*alihouse.AlibabaAlihouseExistinghouseHouseBaseAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseExistinghouseHouseBaseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
