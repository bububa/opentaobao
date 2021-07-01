package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

/* AlibabaWtCifCoopOsstokenGet
获取oss签名接口
alibaba.wt.cif.coop.osstoken.get

商家合作上传oss图片获取token接口 */
func AlibabaWtCifCoopOsstokenGet(clt *core.SDKClient, req *alicom.AlibabaWtCifCoopOsstokenGetAPIRequest, session string) (*alicom.AlibabaWtCifCoopOsstokenGetAPIResponse, error) {
	var resp alicom.AlibabaWtCifCoopOsstokenGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
