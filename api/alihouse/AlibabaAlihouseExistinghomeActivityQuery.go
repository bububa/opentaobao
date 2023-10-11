package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeActivityQuery 五福活动经纪人获取
// alibaba.alihouse.existinghome.activity.query
//
// 五福活动经纪人获取
func AlibabaAlihouseExistinghomeActivityQuery(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeActivityQueryAPIRequest, session string) (*alihouse.AlibabaAlihouseExistinghomeActivityQueryAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseExistinghomeActivityQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
