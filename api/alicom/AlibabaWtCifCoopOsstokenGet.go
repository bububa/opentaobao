package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// Alibabawtcifcooposstokenget 获取oss签名接口
// alibaba.wt.cif.coop.osstoken.get
//
// 商家合作上传oss图片获取token接口
func Alibabawtcifcooposstokenget(clt *core.SDKClient, req *alicom.AlibabawtcifcooposstokengetAPIRequest, session string) (*alicom.AlibabawtcifcooposstokengetAPIResponse, error) {
	var resp alicom.AlibabawtcifcooposstokengetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
