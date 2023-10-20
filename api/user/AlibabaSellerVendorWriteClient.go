package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// Alibabasellervendorwriteclient 客户动态回写
// alibaba.seller.vendor.write.client
//
// 客户动态开放API接口，外部服务商回写数据
func Alibabasellervendorwriteclient(clt *core.SDKClient, req *user.AlibabasellervendorwriteclientAPIRequest, session string) (*user.AlibabasellervendorwriteclientAPIResponse, error) {
	var resp user.AlibabasellervendorwriteclientAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
